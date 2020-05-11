package router

import (
	"net/http"
	"swblog/controllers"

	"github.com/gin-gonic/gin"
)

var backstage *gin.RouterGroup

//RegisterBackStageRoute 注册后台的路由
func RegisterBackStageRoute(eng *gin.Engine) {
	backstage = eng.Group("backstage")
	backstage.GET("/index", backstageIndex)
}

func backstageIndex(ctx *gin.Context) {
	data := controllers.GetBgIndex("")
	ctx.HTML(http.StatusOK, "backstageIndex", data)
}
