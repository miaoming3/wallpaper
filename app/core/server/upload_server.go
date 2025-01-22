package server

import (
	"fmt"
	"github.com/miaoming3/wallpaper/app/global"
	"github.com/miaoming3/wallpaper/app/message"
	"github.com/miaoming3/wallpaper/app/message/dro"
	"github.com/miaoming3/wallpaper/app/message/dto"
	response2 "github.com/miaoming3/wallpaper/app/response"
	"github.com/miaoming3/wallpaper/app/utils"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"github.com/miaoming3/wallpaper/app/response"
)

type UploadServer struct {
}

func NewUploadServer() UploadServer {
	return UploadServer{}
}
func (us *UploadServer) UploadOneFile(c *gin.Context, file *dto.UploadFile) *response.APi {
	ext := filepath.Ext(file.File.Filename)
	if !us.IsExtensionAllowed(ext) {
		return response2.ApiError(message.UploadNotSupportExt, nil)
	}
	if file.File.Size > global.SysConfig.UploadFileSize {
		return response2.ApiError(message.UploadMaxSize, nil)
	}
	uploadDir := filepath.Join(global.SysConfig.Dir, (time.Now()).Format("2006-01-02"))
	if err := us.IsMakeDirectoryCreate(uploadDir); err != nil {
		return response2.ApiError(message.UploadCreateDirErr, err)
	}
	saveFileName := filepath.Join(uploadDir, uuid.New().String()+ext)
	if err := c.SaveUploadedFile(&file.File, saveFileName); err != nil {
		return response2.ApiError(message.UploadSaveError, err)
	}

	return response2.ApiSuccess(dro.UploadResponseData{
		Path: saveFileName,
		Salt: utils.Md5(file.File.Filename),
		Url:  fmt.Sprintf("%v/%v", c.Request.Host, saveFileName),
		Host: c.Request.Host,
	})

}

func (us *UploadServer) UploadsFileMust(c *gin.Context, file *dto.UploadFileMust) *response.APi {
	var updateResponse []dro.UploadResponseData
	for k, v := range file.File {
		ext := filepath.Ext(v.Filename)
		if !us.IsExtensionAllowed(ext) {
			updateResponse = append(updateResponse, dro.UploadResponseData{
				Code:  message.UploadNotSupportExt,
				Error: nil,
				Msg:   message.GetMessage(message.UploadNotSupportExt),
				Key:   k + 1,
			})
			continue
		}
		if v.Size > global.SysConfig.UploadFileSize {
			updateResponse = append(updateResponse, dro.UploadResponseData{
				Code:  message.UploadMaxSize,
				Error: nil,
				Msg:   message.GetMessage(message.UploadMaxSize),
				Key:   k + 1,
			})
			continue
		}

		uploadDir := filepath.Join(global.SysConfig.Dir, (time.Now()).Format("2006-01-02"))
		if err := us.IsMakeDirectoryCreate(uploadDir); err != nil {
			updateResponse = append(updateResponse, dro.UploadResponseData{
				Code:  message.UploadCreateDirErr,
				Error: nil,
				Msg:   message.GetMessage(message.UploadCreateDirErr),
				Key:   k + 1,
			})
			continue
		}
		saveFileName := filepath.Join(uploadDir, uuid.New().String()+ext)
		if err := c.SaveUploadedFile(&v, saveFileName); err != nil {
			updateResponse = append(updateResponse, dro.UploadResponseData{
				Code:  message.UploadSaveError,
				Error: err,
				Msg:   message.GetMessage(message.UploadSaveError),
				Key:   k + 1,
			})
			continue
		}
		updateResponse = append(updateResponse, dro.UploadResponseData{
			Code: http.StatusOK,
			Msg:  "Success",
			Path: saveFileName,
			Salt: utils.Md5(v.Filename),
			Host: c.Request.Host,
			Key:  k + 1,
			Url:  fmt.Sprintf("%v/%v", c.Request.Host, saveFileName),
		})
	}
	return response2.ApiSuccess(updateResponse)

}

func (us *UploadServer) IsMakeDirectoryCreate(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		if err = os.Mkdir(path, 0755); err != nil {
			return err
		}
	}
	return nil
}

func (us *UploadServer) IsExtensionAllowed(ext string) bool {
	if !strings.HasPrefix(ext, ".") {
		ext = "." + ext
	}
	for _, allowedExt := range global.SysConfig.AllowedExtensions {
		if strings.EqualFold(ext, allowedExt) {
			return true // 找到匹配的扩展名，返回 true
		}
	}
	return false
}
