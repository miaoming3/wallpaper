package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/miaoming3/wallpaper/controller/dto"
	"github.com/miaoming3/wallpaper/response"
	"github.com/miaoming3/wallpaper/server"
)

type ImageController struct {
	server.BaseServiceInterface
}

func NewImageController() *ImageController {
	return &ImageController{server.NewImageServer()}
}

// Index
// @Summary 图片列表
// @Tags 图片管理
// @Description 获取图片列表
// @Param ImageSearch query dto.ImageSearch true "图片搜素参数"
// @Produce application/json
// @Router /image/index [get]
func (ic *ImageController) Index(c *gin.Context) {
	var imageSearch dto.ImageSearch
	if err := c.ShouldBind(&imageSearch); err != nil {
		response.Response(c, response.ApiError(response.CLIENTERROR, err))
		return
	}
	response.Response(c, ic.IndexServer(c, &imageSearch))
}

// Save
// @Summary 图片保存
// @Tags 图片管理
// @Description 图片保存
// @Param ImageSave body dto.ImageSave true "保存图片参数"
// @Produce application/json
// @Router /image/save [post]
func (ic *ImageController) Save(c *gin.Context) {
	var imageSave dto.ImageSave
	if err := c.ShouldBind(&imageSave); err != nil {
		response.Response(c, response.ApiError(response.CLIENTERROR, err))
		return
	}
	response.Response(c, ic.CreateServer(&imageSave))

}

// Update
// @Summary 图片修改
// @Tags 图片管理
// @Description 图片修改
// @Param ImageUpdate body dto.ImageUpdate true "修改图片参数"
// @Produce application/json
// @Router /image/update [put]
func (ic *ImageController) Update(c *gin.Context) {
	var imageUpdate dto.ImageUpdate
	if err := c.ShouldBind(&imageUpdate); err != nil {
		response.Response(c, response.ApiError(response.CLIENTERROR, err))
		return
	}
	response.Response(c, ic.UpdateServer(&imageUpdate))
}

func (ic *ImageController) Delete(c *gin.Context) {
	var imageDelete dto.ImageDelete
	if err := c.ShouldBind(&imageDelete); err != nil {
		response.Response(c, response.ApiError(response.CLIENTERROR, err))
		return
	}
	response.Response(c, ic.DeleteServer(&imageDelete))
}
