package v1

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/miaoming3/wallpaper/dao"
	"github.com/miaoming3/wallpaper/response"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type loginRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func PostLogin(c *gin.Context) {
	var login loginRequest
	if err := c.ShouldBind(&login); err != nil {
		response.Error(c, response.CLIENTERROR, err)
		return
	}
	admin, err := dao.NewAdminDao().FindByUsername(login.Username)
	if err != nil {
		response.Error(c, response.CLIENTERROR, err)
		return
	}
	if err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(login.Password)); err != nil {
		response.Error(c, response.CLIENTERROR, err)
		return
	}
	session := sessions.Default(c)
	session.Set("username", admin.Username)
	_ = session.Save()
	response.Success(c, admin)
}
