package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/miaoming3/wallpaper/server"
)

type UserController struct {
	server.BaseServiceInterface
}

func NewUserController() *UserController {
	return &UserController{server.NewMenuServe()}
}

func (ac *UserController) Index(c *gin.Context) {

}
func (ac *UserController) Save(c *gin.Context) {

}
func (ac *UserController) Update(c *gin.Context) {

}
func (ac *UserController) Delete(c *gin.Context) {

}
