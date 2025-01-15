package v1

import (
	"github.com/gin-gonic/gin"
	server2 "github.com/miaoming3/wallpaper/core/server"
)

type UserController struct {
	server2.BaseServiceInterface
}

func NewUserController() *UserController {
	return &UserController{server2.NewMenuServe()}
}

func (ac *UserController) Index(c *gin.Context) {

}
func (ac *UserController) Save(c *gin.Context) {

}
func (ac *UserController) Update(c *gin.Context) {

}
func (ac *UserController) Delete(c *gin.Context) {

}
