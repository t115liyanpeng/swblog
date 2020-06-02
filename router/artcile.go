package router

import (
	"fmt"
	"net/http"
	"strconv"
	"swblog/controllers"
	"swblog/models/artciles"
	"swblog/models/artclassify"
	"swblog/models/page"
	"swblog/tools"

	"github.com/gin-gonic/gin"
)

var artilcegroup *gin.RouterGroup

//RegisterArtilcesGroup 注册文章管理的路由
func RegisterArtilcesGroup(eng *gin.Engine) {
	artilcegroup = eng.Group("/article")

	//注册查看文章详情路由
	artilcegroup.GET("/detail", countNum, artDetail)
	//文章归档
	artilcegroup.GET("/articlefile", artFile)
	//首页文章汇总分页
	artilcegroup.GET("/page", artPage)
	//文章分类列表
	artilcegroup.GET("/artclassifylst", artList)
	//文章分类列表分页
	artilcegroup.GET("/artclassifylstpage", artListPage)
	//获取所有分类
	artilcegroup.GET("/artclassify", artClassify)
	//获取所有分类分页
	artilcegroup.GET("/artclassifypage", artClassifyPage)
	//文章点赞
	artilcegroup.GET("like", setLike)
	//添加文章
	artilcegroup.POST("/addart", addart)
	//删除文章
	artilcegroup.GET("/delart", delart)
}

//统计文章访问数量
func countNum(ctx *gin.Context) {
	artid := ctx.Query("artid")
	if artid != "" {
		mycookie := "art-detail"
		cookie, err := ctx.Cookie(mycookie)
		if err != nil || cookie != artid {
			go controllers.SetArtSeeCnt(artid)
			tools.SetMyCookies(mycookie, artid, 1800, ctx)
		}

	}
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
		//fmt.Printf("test loading page\n")
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

func artList(ctx *gin.Context) {
	pageindex := 1
	index := ctx.Query("index")
	classify := ctx.Query("classify")
	tag := ctx.Query("tag")
	title := ctx.Query("title")
	name := ctx.Query("name")
	if index != "" {
		pageindex, _ = strconv.Atoi(index)
	}
	data := controllers.GetArtList(pageindex, tools.SvrCfg.Server.IndexPageSize, classify, tag, tools.SvrCfg.Server.UserID, name)
	data.Title = title
	data.Param = &artciles.ArtListParam{
		Classify: classify,
		Tag:      tag,
		Name:     name,
	}
	ctx.HTML(http.StatusOK, "artclassifylst", data)
}

func artListPage(ctx *gin.Context) {
	pageindex := 1
	index := ctx.Query("index")
	classify := ctx.Query("classify")
	tag := ctx.Query("tag")
	//title := ctx.Query("title")
	name := ctx.Query("name")
	if index != "" {
		pageindex, _ = strconv.Atoi(index)
	}
	data := controllers.GetArtList(pageindex, tools.SvrCfg.Server.IndexPageSize, classify, tag, tools.SvrCfg.Server.UserID, name)
	ctx.HTML(http.StatusOK, "artpage", data.List)
}

func artClassify(ctx *gin.Context) {
	pageindex := 1
	index := ctx.Query("index")
	param := ctx.Query("param")
	pageindex, _ = strconv.Atoi(index)
	paramInt, _ := strconv.Atoi(param)
	items, cnt := controllers.GetClassOrTags(pageindex, tools.SvrCfg.Server.FilePageSize, paramInt, tools.SvrCfg.Server.UserID)
	title := "全部分类"
	if paramInt == 1 {
		title = "全部标签"
	}
	data := artclassify.ArtClassTag{
		Title:      title,
		List:       items,
		PageSize:   tools.SvrCfg.Server.FilePageSize,
		ArtCount:   cnt,
		ClassOrTag: paramInt,
	}
	ctx.HTML(http.StatusOK, "artclassify", data)
}

func artClassifyPage(ctx *gin.Context) {
	pageindex := 1
	index := ctx.Query("index")
	param := ctx.Query("param")
	pageindex, _ = strconv.Atoi(index)
	paramInt, _ := strconv.Atoi(param)
	items, _ := controllers.GetClassOrTags(pageindex, tools.SvrCfg.Server.FilePageSize, paramInt, tools.SvrCfg.Server.UserID)
	ctx.HTML(http.StatusOK, "artclassifypage", items)
}

func setLike(ctx *gin.Context) {
	artid := ctx.Query("id")
	if artid != "" {
		key := fmt.Sprintf("likeartid-%s", artid)
		ck, err := ctx.Cookie(key)
		if err == nil {
			if ck == artid {
				ctx.JSON(http.StatusOK, gin.H{
					"code": "0",
					"msg":  "您已经点过喜欢了哟！^_^",
				})
				return
			}
		}
		tools.SetMyCookies(key, artid, 1800, ctx)
		go controllers.SetArtLike(tools.SvrCfg.Server.UserID, artid)
		ctx.JSON(http.StatusOK, gin.H{
			"code": "1",
			"msg":  "您的点赞支持是对我的鼓励！(>‿◠)",
		})
	}
}

func addart(ctx *gin.Context) {
	art := artciles.Article{}
	code := -1
	msg := ""
	err := ctx.BindJSON(&art)
	if err == nil {
		err = controllers.AddArticle(&art)
		if err == nil {
			code = 1
			msg = "添加成功"
		} else {
			code = 1
			msg = "写入数据库失败!"
		}
	} else {
		code = 0
		msg = "参数不正确！"
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

func delart(ctx *gin.Context) {
	artid := ctx.Query("artid")
	if artid == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "参数错误",
		})
		return
	}
	err := controllers.DelArticle(artid)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "删除失败！",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "删除成功",
	})
}
