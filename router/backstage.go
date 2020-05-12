package router

import (
	"net/http"
	"swblog/controllers"
	"swblog/models/conf"
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
}

func backstageIndex(ctx *gin.Context) {
	data := controllers.GetBgIndex("")
	ctx.HTML(http.StatusOK, "backstageIndex", data)
}

func backstageBaseConfig(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "wbconfig", tools.SvrCfg)
}

func setconfig(ctx *gin.Context) {
	cfg := &conf.Config{}
	err := ctx.BindJSON(cfg)
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": "1",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":  "0",
		"error": err.Error,
	})

}
