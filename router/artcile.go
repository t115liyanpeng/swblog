package router

import (
	"net/http"
	"swblog/controllers"
	"swblog/tools"

	"github.com/gin-gonic/gin"
)

var artilcegroup *gin.RouterGroup

//RegisterArtilcesGroup 注册文章管理的路由
func RegisterArtilcesGroup(eng *gin.Engine) {
	artilcegroup = eng.Group("/article")

	//注册查看文章详情路由
	artilcegroup.GET("/detail", artDetail)
}

func artDetail(ctx *gin.Context) {

	artid := ctx.Query("artid")
	artdetail := controllers.GetArticleDetail(artid, tools.SvrCfg.Server.UserID)
	if artid == "" {
		ctx.HTML(http.StatusOK, "404", nil)
		return
	}
	ctx.HTML(http.StatusOK, "detail", artdetail)
}
