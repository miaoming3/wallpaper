package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/miaoming3/wallpaper/server"
)

type MenuController struct {
	server.MenuService
}

func NewMenuController() *MenuController {
	return &MenuController{server.NewMenuServe()}
}

func (mc *MenuController) Index(c *gin.Context) {

}
