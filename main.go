package main

import (
	"fmt"
	"net/http"
	"swblog/models/conf"
	"swblog/models/page"
	"swblog/router"
	"swblog/swsqlx"
	"swblog/tools"

	"github.com/gin-gonic/gin"
)

//配置信息
var svrCfg *conf.Config

//gin引擎
var engine *gin.Engine

func main() {

	//读取配置文件

	var err error
	svrCfg, err = tools.ReadConfig()
	if err != nil {
		fmt.Printf("err %v\n", err)
		return
	}
	fmt.Printf("servename:%s\n", svrCfg.Server.WebName)
	fmt.Printf("serverport:%s\n", svrCfg.Server.Port)
	fmt.Printf("database ip address %s\n", svrCfg.Database.IPAddress)
	fmt.Printf("database port %d\n", svrCfg.Database.Port)
	fmt.Printf("database user %s\n", svrCfg.Database.User)
	fmt.Printf("database password %s\n", svrCfg.Database.Password)

	//初始化数据库连接
	swsqlx.CreateDbcInstance(svrCfg.Database)

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
	router.RegisterUserGroup(engine)

	engine.Run(fmt.Sprintf(":%s", svrCfg.Server.Port))

}

//Index 默认页
func indexPage(ctx *gin.Context) {
	list := page.GetLeftNavData()

	um := page.GetWebSietUserInfo()
	index := page.FirstPage{
		Title:    svrCfg.Server.WebName,
		UserInfo: um,
		Left:     list,
	}

	ctx.HTML(http.StatusOK, "index", index)
}
