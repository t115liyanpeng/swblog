package controllers

import (
	"encoding/json"
	"net/http"
	"swblog/models/artclassify"
	"swblog/models/user"
	"swblog/swsqlx"

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

//GetUserDropList 获取用户下拉列表
func GetUserDropList() []*artclassify.UserSimple {
	userlist := make([]*artclassify.UserSimple, 0)
	swsqlx.Dbc.SQLDb.Select(&userlist, "select id,name from t_userb")
	return userlist
}
