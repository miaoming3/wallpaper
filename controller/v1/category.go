package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/miaoming3/wallpaper/controller/dto"
	"github.com/miaoming3/wallpaper/response"
	"github.com/miaoming3/wallpaper/server"
)

type CategoryController struct {
	Server *server.CategoryServer
}

func NewCategoryController() *CategoryController {
	return &CategoryController{
		Server: server.NewCategoryServer(),
	}
}

func (cg *CategoryController) Index(c *gin.Context) {
	var searchData dto.CategoryIndex
	if err := c.ShouldBind(&searchData); err != nil {
		response.Response(c, response.ApiError(response.CLIENTERROR, err))
		return
	}

	response.Response(c, cg.Server.IndexServer(c, &searchData))
}

func (cg *CategoryController) Save(c *gin.Context) {
	var saveCategory dto.SaveCategory
	if err := c.ShouldBind(&saveCategory); err != nil {
		response.Response(c, response.ApiError(response.CLIENTERROR, err))
		return
	}
	response.Response(c, cg.Server.CreateServer(c, &saveCategory))
}
