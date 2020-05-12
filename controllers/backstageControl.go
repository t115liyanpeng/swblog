package controllers

import (
	"encoding/json"
	"swblog/models/conf"
	"swblog/models/page"
	"swblog/tools"
)

//GetBgIndex 获取后台默认页的数据
func GetBgIndex(uid string) *page.BackageStage {
	return &page.BackageStage{
		ArtCount:     10,
		CommentCount: 9,
		AllSee:       1000,
		TodaySee:     2,
	}
}

//SetConfig 设置配置信息
func SetConfig(cfg *conf.Config) bool {
	/*
		tools.SvrCfg.Server.WebName = cfg.Server.WebName
		tools.SvrCfg.Server.Port = cfg.Server.Port
		tools.SvrCfg.Server.IndexPageSize = cfg.Server.IndexPageSize
		tools.SvrCfg.Server.FilePageSize = cfg.Server.FilePageSize
		tools.SvrCfg.Database.IPAddress = cfg.Database.IPAddress
		tools.SvrCfg.Database.DbName = cfg.Database.DbName
		tools.SvrCfg.Database.Port = cfg.Database.Port
		tools.SvrCfg.Database.User = cfg.Database.User
		tools.SvrCfg.Database.Password = cfg.Database.Password
	*/
	cfg.Server.UserID = tools.SvrCfg.Server.UserID
	js, err := json.Marshal(cfg)
	if err == nil {
		return tools.SetConfigJSONFile(js)
	}
	return false
}
