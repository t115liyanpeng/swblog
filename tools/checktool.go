package tools

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"swblog/models/conf"
	"swblog/swsqlx"
)

//SvrCfg 配置信息 全局变量存储服务的基础信息
var SvrCfg *conf.Config = &conf.Config{}

//TodayCnt 日浏览量
var TodayCnt int

//ReadConfig 读取配置文件
func ReadConfig() (cf *conf.Config, err error) {
	cfg, err := CheckConfig()
	if err != nil {
		return nil, err
	}
	cfgbytes, err := ioutil.ReadFile(cfg)
	if err != nil {
		fmt.Println("read file err")
		return nil, err
	}
	cf = &conf.Config{}
	err = json.Unmarshal(cfgbytes, cf)
	if err != nil {
		return nil, err
	}
	return cf, nil
}

//CheckConfig 检查配置文件
func CheckConfig() (string, error) {
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

//SetConfigJSONFile 更新配置文件
func SetConfigJSONFile(json []byte) bool {
	cfgpath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return false
	}

	cfgpath = fmt.Sprintf("%s/conf/conf.json", cfgpath)
	err = ioutil.WriteFile(cfgpath, json, 0777)
	if err != nil {
		return false
	}
	return true
}

//UpdateSeeCnt 更新访问量
func UpdateSeeCnt() {
	swsqlx.Dbc.SQLDb.Exec("update t_swsystemb set allsee=allsee+1 where id = 1")
	TodayCnt++
}

//GetMd5Str 获取md5加密字符串
func GetMd5Str(src string) string {
	data := []byte(src)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}
