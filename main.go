package main

import (
	"fmt"
	"net/http"
	"swblog/models/page"
	"swblog/router"
	"swblog/swsqlx"
	"swblog/tools"

	"github.com/gin-gonic/gin"
)

//gin引擎
var engine *gin.Engine

func main() {

	//读取配置文件

	var err error

	tools.SvrCfg, err = tools.ReadConfig()
	if err != nil {
		fmt.Printf("err %v\n", err)
		return
	}
	/*
		fmt.Printf("servename:%s\n", svrCfg.Server.WebName)
		fmt.Printf("serverport:%s\n", svrCfg.Server.Port)
		fmt.Printf("database ip address %s\n", svrCfg.Database.IPAddress)
		fmt.Printf("database port %d\n", svrCfg.Database.Port)
		fmt.Printf("database user %s\n", svrCfg.Database.User)
		fmt.Printf("database password %s\n", svrCfg.Database.Password)
	*/
	//初始化数据库连接
	swsqlx.CreateDbcInstance(tools.SvrCfg.Database)

	//初始化gin
	gin.SetMode(gin.ReleaseMode)
	engine = gin.New()
	//使用日志
	engine.Use(gin.Logger())
	//使用Panic处理方案
	engine.Use(gin.Recovery())
	//engine = gin.Default()

	//加载静态文件
	engine.StaticFS("/static", http.Dir("./static"))

	//加载模板文件
	engine.LoadHTMLGlob("views/**/*")
	//engine.LoadHTMLFiles("views/index.html")

	//注册默认页
	engine.GET("/", indexPage)

	//注册路由
	//注册用户模块
	router.RegisterUserGroup(engine)
	//注册文章模块
	router.RegisterArtilcesGroup(engine)
	//注册后台模块
	router.RegisterBackStageRoute(engine)

	//定义404
	engine.NoRoute(go404)

	engine.Run(fmt.Sprintf(":%s", tools.SvrCfg.Server.Port))

}

//Index 默认页
func indexPage(ctx *gin.Context) {
	//userid 默认展示用户的信息的用户id
	//userid := "c999a2f041c84dc1b5970bb973c1da74"
	list := page.GetLeftNavData(tools.SvrCfg.Server.UserID)
	um := page.GetWebSietUserInfo(tools.SvrCfg.Server.UserID)
	articles, count := page.GetContent(tools.SvrCfg.Server.UserID, tools.SvrCfg.Server.IndexPageSize, 1)
	index := page.FirstPage{
		Title:       tools.SvrCfg.Server.WebName,
		UserInfo:    um,
		Left:        list,
		Articles:    articles,
		ArtCount:    count,
		ArtPageSize: tools.SvrCfg.Server.IndexPageSize,
		News:        page.ArtBiref(articles),
		Hots:        page.GetHots(tools.SvrCfg.Server.UserID),
	}
	ctx.HTML(http.StatusOK, "index", index)
}

//404
func go404(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "404", nil)
}
