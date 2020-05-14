package router

import (
	"fmt"
	"net/http"
	"swblog/controllers"
	"swblog/models/conf"
	"swblog/models/user"
	"swblog/tools"

	"github.com/gin-gonic/gin"
)

var backstage *gin.RouterGroup

//RegisterBackStageRoute 注册后台的路由
func RegisterBackStageRoute(eng *gin.Engine) {
	backstage = eng.Group("backstage")
	backstage.GET("/index", backstageIndex)
	backstage.GET("/wbconfig", backstageBaseConfig)
	backstage.POST("/setconfig", setconfig)
	backstage.GET("/wuconfig", wuconfig)
	backstage.POST("/setuser", setuserconfig)
}

func backstageIndex(ctx *gin.Context) {
	data := controllers.GetBgIndex(tools.SvrCfg.Server.UserID)
	ctx.HTML(http.StatusOK, "backstageIndex", data)
}

func backstageBaseConfig(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "wbconfig", tools.SvrCfg)
}

func setconfig(ctx *gin.Context) {
	cfg := conf.Config{
		Server:   &conf.Servercfg{},
		Database: &conf.Databasecfg{},
	}
	err := ctx.BindJSON(&cfg)
	if err == nil {
		if controllers.SetConfig(&cfg) {
			ctx.JSON(http.StatusOK, gin.H{
				"code": "1",
			})
			return
		}
		err = fmt.Errorf("set json faild")
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":  "0",
		"error": err.Error,
	})

}

func wuconfig(ctx *gin.Context) {
	data := controllers.GetDbUserInfo(tools.SvrCfg.Server.UserID)
	ctx.HTML(http.StatusOK, "wuserconfig", data)
}

func setuserconfig(ctx *gin.Context) {
	dbu := &user.DbUser{}
	err := ctx.BindJSON(dbu)
	if err == nil {
		ret, msg := dbu.UpdateDbUser()
		if ret {
			ctx.JSON(http.StatusOK, gin.H{
				"code": "1",
				"msg":  msg,
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"code": "1",
				"msg":  msg,
			})
		}
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": "-1",
			"msg":  "read json faild",
		})
	}
}
