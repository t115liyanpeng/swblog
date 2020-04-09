package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"swblog/models/conf"
	"swblog/swsqlx"

	"github.com/gin-gonic/gin"
)

//配置信息
var svrCfg *conf.Config

//gin引擎
var engine *gin.Engine

func main() {

	//读取配置文件

	/*

	 */
	svrCfg, err := readconfig()
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

	//main 函数结束的时候关闭mysql 连接
	defer swsqlx.Dbc.Close()
}

func readconfig() (cf *conf.Config, err error) {
	cfg, err := checkconfig()
	if err != nil {
		return nil, err
	}
	cfgbytes, err := ioutil.ReadFile(cfg)
	if err != nil {
		fmt.Println("read file err")
		return nil, err
	} else {
		cf = &conf.Config{}
		err = json.Unmarshal(cfgbytes, cf)
		if err != nil {
			return nil, err
		}
		return cf, nil
	}
}

func checkconfig() (string, error) {
	//判断根路径下配置文件是否存在
	cfgpath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}

	cfgpath = fmt.Sprintf("%s/conf/conf.json", cfgpath)
	_, err = os.Stat(cfgpath)
	if err != nil {
		if os.IsExist(err) {
			return cfgpath, nil
		}
		return "", err
	}
	return cfgpath, nil
}

//初始化gin引擎
func initedGinEngine(cfg *conf.Config) *gin.Engine {
	r := gin.Default()
	return r
}

//初始化mysql连接
func initedMysql() error {
	err := swsqlx.InitedMysql(svrCfg.Database)
	return err
}
