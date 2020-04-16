package router

import (
	"net/http"
	"swblog/controllers"

	"github.com/gin-gonic/gin"
)

var usergroup *gin.RouterGroup

//RegisterUserGroup 注册用户组路由
func RegisterUserGroup(eng *gin.Engine) {
	if eng != nil {
		usergroup = eng.Group("/user")

		//注册登录路由
		usergroup.POST("/login", userLogin)

		usergroup.GET("/login", loginPage)
	}
}

func userLogin(context *gin.Context) {
	controllers.UserLoginFunc(context)
}

func loginPage(context *gin.Context) {
	context.HTML(http.StatusOK, "login", gin.H{})
}
