package controllers

import (
	"net/http"
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
			//jsonstr, _ := json.Marshal(u)
			//fmt.Printf("resault:%s\n", string(jsonstr))
			//ctx.JSON(http.StatusOK, string(jsonstr))
			ctx.JSON(http.StatusOK, gin.H{
				"loginname": u.LoginName,
				"code":      u.State.Code,
				"name":      u.Name,
				"msg":       u.State.Msg,
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": u.State.Msg,
			})
		}
	} else {
		ctx.JSON(http.StatusOK, gin.H{})
	}
}
