package server

import (
	"fmt"
	"github.com/miaoming3/wallpaper/app/global"
	response2 "github.com/miaoming3/wallpaper/app/response"
	"github.com/miaoming3/wallpaper/app/response/dro"
	"github.com/miaoming3/wallpaper/app/utils"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"github.com/miaoming3/wallpaper/app/dto"
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
		return response2.ApiError(response2.UploadNotSupportExt, nil)
	}
	if file.File.Size > global.SysConfig.UploadFileSize {
		return response2.ApiError(response2.UploadMaxSize, nil)
	}
	uploadDir := filepath.Join(global.SysConfig.Dir, (time.Now()).Format("2006-01-02"))
	if err := us.IsMakeDirectoryCreate(uploadDir); err != nil {
		return response2.ApiError(response2.UploadCreateDirErr, err)
	}
	saveFileName := filepath.Join(uploadDir, uuid.New().String()+ext)
	if err := c.SaveUploadedFile(&file.File, saveFileName); err != nil {
		return response2.ApiError(response2.UploadSaveError, err)
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
				Code:  response2.UploadNotSupportExt,
				Error: nil,
				Msg:   response2.GetMessage(response2.UploadNotSupportExt),
				Key:   k + 1,
			})
			continue
		}
		if v.Size > global.SysConfig.UploadFileSize {
			updateResponse = append(updateResponse, dro.UploadResponseData{
				Code:  response2.UploadMaxSize,
				Error: nil,
				Msg:   response2.GetMessage(response2.UploadMaxSize),
				Key:   k + 1,
			})
			continue
		}

		uploadDir := filepath.Join(global.SysConfig.Dir, (time.Now()).Format("2006-01-02"))
		if err := us.IsMakeDirectoryCreate(uploadDir); err != nil {
			updateResponse = append(updateResponse, dro.UploadResponseData{
				Code:  response2.UploadCreateDirErr,
				Error: nil,
				Msg:   response2.GetMessage(response2.UploadCreateDirErr),
				Key:   k + 1,
			})
			continue
		}
		saveFileName := filepath.Join(uploadDir, uuid.New().String()+ext)
		if err := c.SaveUploadedFile(&v, saveFileName); err != nil {
			updateResponse = append(updateResponse, dro.UploadResponseData{
				Code:  response2.UploadSaveError,
				Error: err,
				Msg:   response2.GetMessage(response2.UploadSaveError),
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
