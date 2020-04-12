package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var usergroup *gin.RouterGroup

//RegisterUserGroup 注册用户组路由
func RegisterUserGroup(eng *gin.Engine) {
	if eng != nil {
		usergroup = eng.Group("/user")

		//用户登录页
		eng.LoadHTMLGlob("views/static/**/*")

		//注册登录路由
		usergroup.POST("/login", userLogin)

		usergroup.GET("/login", loginPage)
	}
}

func userLogin(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "success",
	})
}

func loginPage(context *gin.Context) {
	context.HTML(http.StatusOK, "login.html", gin.H{})
}
