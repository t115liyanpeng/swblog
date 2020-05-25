package controllers

import (
	"encoding/json"
	"swblog/models/artclassify"
	"swblog/models/conf"
	"swblog/models/page"
	"swblog/models/user"
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

//GetDbUserInfo 获取数据库用户信息
func GetDbUserInfo(uid string) *user.DbUser {
	dbu := &user.DbUser{}
	if dbu.GetUserByID(uid) == nil {
		return dbu
	}
	return nil
}

//GetClassJosn 获取用户所有分类
func GetClassJosn(uid string) *page.TableJSONData {

	classify := make([]*artclassify.Classify, 0)
	err := swsqlx.Dbc.SQLDb.Select(&classify, "SELECT id,name,icon,brief FROM t_classifyb WHERE pid=0 AND userid=?", uid)
	if err == nil {
		data := page.TableJSONData{
			Code:  0,
			Msg:   "",
			Count: 1,
			Data:  classify,
		}

		return &data
	}
	return nil
}

//AddClassify 添加新分类
func AddClassify(uid string, classinfo *artclassify.Classify) bool {
	_, err := swsqlx.Dbc.SQLDb.Exec("insert into t_classifyb (pid,userid,name,icon,brief)values(?,?,?,?,?)",
		0, uid, classinfo.Name, classinfo.Icon, classinfo.Brief)
	if err == nil {
		return true
	}
	return false
}

//DelClassify 删除分类by classid
func DelClassify(classid string) (ret bool, err error) {
	ret = false
	_, err = swsqlx.Dbc.SQLDb.Exec("DELETE FROM t_classifyb WHERE id=?", classid)
	if err == nil {
		ret = true
		return
	}
	return
}
