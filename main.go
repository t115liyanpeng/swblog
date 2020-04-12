package main

import (
	"fmt"
	"swblog/models/conf"
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

	svrCfg, err := tools.ReadConfig()
	if err != nil {
		fmt.Printf("err %v\n", err)
		return
	}
	fmt.Printf("serverport:%s\n", svrCfg.Server.Port)
	fmt.Printf("database ip address %s\n", svrCfg.Database.IPAddress)
	fmt.Printf("database port %d\n", svrCfg.Database.Port)
	fmt.Printf("database user %s\n", svrCfg.Database.User)
	fmt.Printf("database password %s\n", svrCfg.Database.Password)

	//初始化gin
	gin.SetMode(gin.ReleaseMode)
	engine = gin.New()
	//使用日志
	engine.Use(gin.Logger())
	//使用Panic处理方案
	engine.Use(gin.Recovery())

	//注册默认页
	engine.GET("/", Index)

	//注册路由
	router.RegisterUserGroup(engine)

	engine.Run(fmt.Sprintf(":%s", svrCfg.Server.Port))

}

//初始化mysql对象
func initedMysql() {
	swsqlx.CreateDbcInstance(svrCfg.Database)
}

//Index 默认页
func Index(ctx *gin.Context) {

}
