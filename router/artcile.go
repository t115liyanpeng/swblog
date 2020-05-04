package router

import (
	"fmt"
	"net/http"
	"strconv"
	"swblog/controllers"
	"swblog/models/page"
	"swblog/tools"

	"github.com/gin-gonic/gin"
)

var artilcegroup *gin.RouterGroup

//RegisterArtilcesGroup 注册文章管理的路由
func RegisterArtilcesGroup(eng *gin.Engine) {
	artilcegroup = eng.Group("/article")

	//注册查看文章详情路由
	artilcegroup.GET("/detail", artDetail)
	//文章归档
	artilcegroup.GET("/articlefile", artFile)
	//首页文章汇总分页
	artilcegroup.GET("page", artPage)
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

func artFile(ctx *gin.Context) {
	var pageindex int = 1
	index := ctx.Query("index")
	if index == "" {
		//跳转
		line := controllers.GetTimeLineArticle(tools.SvrCfg.Server.UserID, tools.SvrCfg.Server.FilePageSize, pageindex)
		ctx.HTML(http.StatusOK, "articlefile", line)
	} else {
		//分页
		fmt.Printf("test loading page\n")
		pageindex, _ = strconv.Atoi(index)
		line := controllers.GetTimeLineArticle(tools.SvrCfg.Server.UserID, tools.SvrCfg.Server.FilePageSize, pageindex)
		ctx.HTML(http.StatusOK, "artfilepage", line)
	}
}

func artPage(ctx *gin.Context) {
	index := ctx.Query("index")
	pageindex, err := strconv.Atoi(index)
	if err == nil {
		articles, _ := page.GetContent(tools.SvrCfg.Server.UserID, tools.SvrCfg.Server.IndexPageSize, pageindex)
		//fmt.Printf("arts : %v count %d", articles, count)
		ctx.HTML(http.StatusOK, "artpage", articles)
	}
	//ctx.JSON(http.StatusOK, gin.H{
	//	"msg": "test",
	//})
}
