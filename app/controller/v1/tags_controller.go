package v1

import (
	"github.com/gin-gonic/gin"
	server2 "github.com/miaoming3/wallpaper/core/server"
)

type TagsController struct {
	server2.BaseServiceInterface
}

func NewTagsController() *TagsController {
	return &TagsController{server2.NewMenuServe()}
}

func (ac *TagsController) Index(c *gin.Context) {

}
func (ac *TagsController) Save(c *gin.Context) {

}
func (ac *TagsController) Update(c *gin.Context) {

}
func (ac *TagsController) Delete(c *gin.Context) {

}
