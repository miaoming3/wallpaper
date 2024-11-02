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

func (ic *ImageController) Index(c *gin.Context) {
	var imageSearch dto.ImageSearch
	if err := c.ShouldBind(&imageSearch); err != nil {
		response.Response(c, response.ApiError(response.CLIENTERROR, err))
		return
	}

	response.Response(c, ic.IndexServer(&imageSearch))
}

func (ic *ImageController) Save(c *gin.Context) {
	var imageSave dto.ImageSave
	if err := c.ShouldBind(&imageSave); err != nil {
		response.Response(c, response.ApiError(response.CLIENTERROR, err))
		return
	}

	response.Response(c, ic.CreateServer(&imageSave))

}

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
