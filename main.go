package main

import (
	"fmt"
	"net/http"
	"swblog/models/artciles"
	"swblog/models/page"
	"swblog/router"
	"swblog/swsqlx"
	"swblog/tools"
	"sync"
	"time"

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

	fmt.Printf("servename:%s\n", tools.SvrCfg.Server.WebName)
	fmt.Printf("serverport:%s\n", tools.SvrCfg.Server.Port)
	fmt.Printf("database ip address %s\n", tools.SvrCfg.Database.IPAddress)
	fmt.Printf("database port %d\n", tools.SvrCfg.Database.Port)
	fmt.Printf("database user %s\n", tools.SvrCfg.Database.User)
	fmt.Printf("database password %s\n", tools.SvrCfg.Database.Password)

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
	fmt.Printf("user module ready...\n")
	//注册文章模块
	router.RegisterArtilcesGroup(engine)
	fmt.Printf("artiles module ready...\n")
	//注册后台模块
	router.RegisterBackStageRoute(engine)
	fmt.Printf("backstage module ready...\n")
	//定义404
	engine.NoRoute(go404)

	//启动维护日访问量的定时器
	todayCntTimer()

	engine.Run(fmt.Sprintf(":%s", tools.SvrCfg.Server.Port))

}

//Index 默认页
func indexPage(ctx *gin.Context) {

	//获取cookies查看有没有浏览过
	ck, err := ctx.Cookie(tools.SvrCfg.Server.WebName)
	if err != nil {
		setCountSeeCookies(tools.SvrCfg.Server.WebName, tools.SvrCfg.Server.UserID, ctx)
	}
	if ck != tools.SvrCfg.Server.UserID {
		setCountSeeCookies(tools.SvrCfg.Server.WebName, tools.SvrCfg.Server.UserID, ctx)
	}
	//userid 默认展示用户的信息的用户id
	//userid := "c999a2f041c84dc1b5970bb973c1da74"
	var wait sync.WaitGroup

	list := make([]*page.LeftTags, 0)
	um := &page.UserModule{}
	articles := make([]*artciles.Article, 0)
	count := 0
	pics := make([]string, 0)
	wait.Add(4)
	go func() {
		list = page.GetLeftNavData(tools.SvrCfg.Server.UserID)
		wait.Done()
	}()
	go func() {
		um = page.GetWebSietUserInfo(tools.SvrCfg.Server.UserID)
		wait.Done()
	}()
	go func() {
		articles, count = page.GetContent(tools.SvrCfg.Server.UserID, tools.SvrCfg.Server.IndexPageSize, 1)
		wait.Done()
	}()
	go func() {
		pics = page.GetLunBoPicPath(tools.SvrCfg.Server.UserID)
		wait.Done()
	}()
	wait.Wait()
	index := page.FirstPage{
		Title:       tools.SvrCfg.Server.WebName,
		UserInfo:    um,
		Left:        list,
		Articles:    articles,
		ArtCount:    count,
		ArtPageSize: tools.SvrCfg.Server.IndexPageSize,
		News:        page.ArtBiref(articles),
		Hots:        page.GetHots(tools.SvrCfg.Server.UserID),
		Pics:        pics,
	}
	ctx.HTML(http.StatusOK, "index", index)

}

//404
func go404(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "404", nil)
}

//set see cookies
func setCountSeeCookies(key, value string, ctx *gin.Context) {
	//设置cookies
	ctx.SetCookie(key, value, 1800, "/", "localhost", false, true)
	//记录访问
	go tools.UpdateSeeCnt()
}

//维护日访问量的计时器
func todayCntTimer() {
	today := time.NewTimer(10 * time.Second)
	select {
	case <-today.C:
		updatetoday()
	}

}

func updatetoday() {
	go func() {
		dt := time.Now()
		if dt.Hour() == 1 {
			tools.TodayCnt = 0
		}
	}()
}
