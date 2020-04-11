package main

import (
	"fmt"
	"swblog/models/conf"
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

	engine.Run(svrCfg.Server.Port)

}

//初始化gin引擎
func initedGinEngine(cfg *conf.Config) *gin.Engine {
	r := gin.Default()
	return r
}

//初始化mysql对象
func initedMysql() {
	swsqlx.CreateDbcInstance(svrCfg.Database)
}
