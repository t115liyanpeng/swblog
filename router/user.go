package router

import "github.com/gin-gonic/gin"

var usergroup *gin.RouterGroup

//RegisterUserGroup 注册用户组路由
func RegisterUserGroup(eng *gin.Engine) {
	if eng != nil {
		usergroup = eng.Group("/user")

		//注册登录路由
		usergroup.POST("/login", userLogin)
	}
}

func userLogin(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "success",
	})
}
