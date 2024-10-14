package v1

import "github.com/gin-gonic/gin"

type loginRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func Login(c *gin.Context) {
	var login loginRequest
	if err := c.ShouldBind(&login); err != nil {
		c.Redirect(200, "/login")
	}
	c.Redirect(200, "/index")
}
