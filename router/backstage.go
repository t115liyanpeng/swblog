package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var backstage *gin.RouterGroup

//RegisterBackStageRoute 注册后台的路由
func RegisterBackStageRoute(eng *gin.Engine) {
	backstage = eng.Group("backstage")
	backstage.GET("/index", backstageIndex)
}

func backstageIndex(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "backstageIndex", nil)
}
