package router

import (
	"net/http"
	"swblog/models/artciles"

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

	artdetail := &artciles.Article{
		ID:         1,
		Name:       "章名称",
		Content:    "章摘要",
		Top:        false,
		Like:       10,
		Click:      1,
		Classify:   "类",
		Tag:        "签",
		Author:     "者",
		CreateTime: "建时间",
	}
	ctx.HTML(http.StatusOK, "detail", artdetail)
}
