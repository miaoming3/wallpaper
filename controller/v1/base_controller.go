package v1

import (
	"github.com/gin-gonic/gin"
)

type BaseControllerInterface interface {
	Index(c *gin.Context)
	Save(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}
type BaseController struct {
	CategoryController *CategoryController
	MenuController     *MenuController
	AdminController    *AdminController
	TagsController     *TagsController
	GradeController    *GradeController
	UserController     *UserController
	ImageController    *ImageController
}

func NewBaseController() *BaseController {
	return &BaseController{
		CategoryController: NewCategoryController(),
		MenuController:     NewMenuController(),
		AdminController:    NewAdminController(),
		TagsController:     NewTagsController(),
		GradeController:    NewGradeController(),
		UserController:     NewUserController(),
		ImageController:    NewImageController(),
	}
}
