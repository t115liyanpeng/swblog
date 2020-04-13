package controllers

import (
	"swblog/models/user"

	"github.com/gin-gonic/gin"
)

//UserLoginFunc 用户登录处理函数
func UserLoginFunc(ctx *gin.Context) {
	var u user.User = user.User{
		State: &user.LoginState{},
	}
	err := ctx.Bind(&u)
	if err == nil {
		err = u.UserLogin()
		if err == nil {
			if u.State.State {
				//登录成功
			} else {
				ctx.JSON(200, gin.H{
					"msg": u.State.Msg,
				})
			}

		} else {
			ctx.JSON(200, gin.H{
				"msg": u.State.Msg,
			})
		}
	} else {

	}
}
