package controllers

import (
	"encoding/json"
	"net/http"
	"swblog/models/user"

	"github.com/gin-gonic/gin"
)

//UserLoginFunc 用户登录处理函数
func UserLoginFunc(ctx *gin.Context) {

	var u user.User = user.User{
		State: &user.LoginState{},
	}
	err := ctx.BindJSON(&u)
	if err == nil {
		err = u.UserLogin()
		jsbyte, err := json.Marshal(u)
		if err == nil {
			ctx.JSON(http.StatusOK, string(jsbyte))
		}
	}
}
