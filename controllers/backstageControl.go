package controllers

import (
	"encoding/json"
	"swblog/models/conf"
	"swblog/models/page"
	"swblog/swsqlx"
	"swblog/tools"
	"sync"
)

//GetBgIndex 获取后台默认页的数据
func GetBgIndex(uid string) *page.BackageStage {
	wg := sync.WaitGroup{}
	all := 0
	artc := 0
	wg.Add(2)
	go func() {
		swsqlx.Dbc.SQLDb.Get(&all, "select allsee from t_swsystemb where id=1")
		wg.Done()
	}()
	go func() {
		swsqlx.Dbc.SQLDb.Get(&artc, "SELECT COUNT(id) as allsee FROM t_articleb WHERE userid=?", uid)
		wg.Done()
	}()

	wg.Wait()
	return &page.BackageStage{
		ArtCount:     artc,
		CommentCount: 0,
		AllSee:       all,
		TodaySee:     tools.TodayCnt,
	}
}

//SetConfig 设置配置信息
func SetConfig(cfg *conf.Config) bool {
	cfg.Server.UserID = tools.SvrCfg.Server.UserID
	js, err := json.MarshalIndent(cfg, "", "\t")
	if err == nil {
		return tools.SetConfigJSONFile(js)
	}
	return false
}
