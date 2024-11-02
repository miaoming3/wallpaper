package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/miaoming3/wallpaper/server"
)

type TagsController struct {
	server.BaseServiceInterface
}

func NewTagsController() *TagsController {
	return &TagsController{server.NewMenuServe()}
}

func (ac *TagsController) Index(c *gin.Context) {

}
func (ac *TagsController) Save(c *gin.Context) {

}
func (ac *TagsController) Update(c *gin.Context) {

}
func (ac *TagsController) Delete(c *gin.Context) {

}
