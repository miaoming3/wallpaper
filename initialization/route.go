package initialization

import (
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "github.com/miaoming3/wallpaper/controller/v1"
	docs "github.com/miaoming3/wallpaper/docs"
	"github.com/miaoming3/wallpaper/global"
	"github.com/miaoming3/wallpaper/response"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func InitRoutes() *gin.Engine {
	gin.SetMode(global.SysConfig.Model)
	r := gin.Default()
	r.NoRoute(func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusNotFound, response.ApiError(http.StatusNotFound, nil))
	})
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Static("/uploads", "./uploads")
	baseController := v1.NewBaseController()
	uploadController := v1.NewUploadController()
	r.POST("uploads/file", uploadController.UploadFile)
	r.POST("uploads/must", uploadController.UploadFileMust)
	setupRoutes(r, baseController)

	return r
}

func setupRoutes(r *gin.Engine, controller *v1.BaseController) {
	api := r.Group("/api/v1")
	{

		registerRoutes(api, "admin", controller.AdminController)
		registerRoutes(api, "category", controller.CategoryController)
		registerRoutes(api, "menu", controller.MenuController)
		registerRoutes(api, "tags", controller.TagsController)
		registerRoutes(api, "grade", controller.GradeController)
		registerRoutes(api, "user", controller.UserController)
		registerRoutes(api, "image", controller.ImageController) // 新增ImageController路由注册
	}
}

func registerRoutes(r *gin.RouterGroup, prefix string, controller interface{}) {
	if crudController, ok := controller.(v1.BaseControllerInterface); ok {
		if adminController, ok := controller.(*v1.AdminController); ok {
			r.POST(fmt.Sprintf("/%s/login", prefix), adminController.Login)
		}
		r.GET(fmt.Sprintf("/%s/index", prefix), crudController.Index)
		r.POST(fmt.Sprintf("/%s/save", prefix), crudController.Save)
		r.PUT(fmt.Sprintf("/%s/update", prefix), crudController.Update)
		r.DELETE(fmt.Sprintf("/%s/del", prefix), crudController.Delete)
	}

}
