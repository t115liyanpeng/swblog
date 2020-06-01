package controllers

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"swblog/models/artciles"
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
	tx := ""
	author := ""
	wg.Add(4)
	go func() {
		swsqlx.Dbc.SQLDb.Get(&all, "select allsee from t_swsystemb where id=1")
		wg.Done()
	}()
	go func() {
		swsqlx.Dbc.SQLDb.Get(&artc, "SELECT COUNT(id) as allsee FROM t_articleb WHERE userid=?", uid)
		wg.Done()
	}()
	go func() {
		swsqlx.Dbc.SQLDb.Get(&tx, "SELECT txurl FROM t_userb WHERE id=?", uid)
		wg.Done()
	}()
	go func() {
		swsqlx.Dbc.SQLDb.Get(&author, "SELECT name FROM t_userb WHERE id=?", uid)
		wg.Done()
	}()

	wg.Wait()
	return &page.BackageStage{
		Author:       author,
		ArtCount:     artc,
		CommentCount: 0,
		AllSee:       all,
		TodaySee:     tools.TodayCnt,
		UserTx:       tx,
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

//GetClassDropList 获取类别下拉列表
func GetClassDropList(uid string) []*artclassify.ClassSimple {
	droplist := make([]*artclassify.ClassSimple, 0)
	swsqlx.Dbc.SQLDb.Select(&droplist, "SELECT id,name FROM t_classifyb WHERE pid=0 AND userid=?", uid)
	return droplist
}

//GetTagDropListByClassID 根据分类id获取标签下拉列表
func GetTagDropListByClassID(id int) []*artclassify.ClassSimple {
	droplist := make([]*artclassify.ClassSimple, 0)
	swsqlx.Dbc.SQLDb.Select(&droplist, "SELECT id,name FROM t_classifyb WHERE  pid=?", id)
	return droplist
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

//UpdateClass 更新分类信息
func UpdateClass(data *artclassify.Classify) (ret bool, err error) {
	ret = false
	_, err = swsqlx.Dbc.SQLDb.Exec("UPDATE t_classifyb SET name=?,icon=?,brief=? WHERE id=?", data.Name, data.Icon, data.Brief, data.ID)
	if err == nil {
		ret = true
		return
	}
	return
}

//GetTagsJSON 获取标签
func GetTagsJSON(uid string) *page.TableJSONData {
	tags := make([]*artclassify.Tags, 0)
	err := swsqlx.Dbc.SQLDb.Select(&tags, "SELECT t1.id,t1.pid,t1.name,t2.name as classname,t1.brief FROM t_classifyb t1 LEFT JOIN t_classifyb t2 ON t1.pid=t2.id WHERE t1.pid<>0 AND t1.userid=?", uid)
	if err == nil {
		data := page.TableJSONData{
			Code:  0,
			Msg:   "",
			Count: 1,
			Data:  tags,
		}

		return &data
	}
	return nil
}

//AddTag 添加新标签
func AddTag(uid string, tag *artclassify.Tags) bool {
	_, err := swsqlx.Dbc.SQLDb.Exec("INSERT INTO t_classifyb (userid,pid,name,brief) VALUES(?,?,?,?)",
		uid, tag.PID, tag.Name, tag.Brief)
	if err == nil {
		return true
	}
	return false
}

//UpdateTag 更新标签内容
func UpdateTag(data *artclassify.Tags) (ret bool) {
	ret = false
	_, err := swsqlx.Dbc.SQLDb.Exec("UPDATE t_classifyb SET name=?,pid=?,brief=? WHERE id=?", data.Name, data.PID, data.Brief, data.ID)
	if err == nil {
		ret = true
		return
	}
	return
}

//SaveLunBoPic 保存录播图片
func SaveLunBoPic(path, name string) (err error) {
	md5, err := tools.GetMd5File(path)
	if err != nil {
		err = fmt.Errorf("获取文件MD5值失败！")
		return
	}
	//查重
	id := ""
	err = swsqlx.Dbc.SQLDb.Get(&id, "select id from t_lubopicb where md5=?", md5)
	if err == nil || id != "" {
		err = fmt.Errorf("文件已存在！请不要重复上传")
		return
	}

	webdir := fmt.Sprintf("/static/user/%s", name)
	_, err = swsqlx.Dbc.SQLDb.Exec("INSERT INTO t_lubopicb SET name=?,userid=?,md5=?,webdir=?", name, tools.SvrCfg.Server.UserID, md5, webdir)
	return
}

//DelLunBoPic 删除轮播图片
func DelLunBoPic(picid string) (err error) {
	webdir := ""
	err = swsqlx.Dbc.SQLDb.Get(&webdir, "select webdir from t_lubopicb where id=?", picid)
	if err == nil {
		tx, err := swsqlx.Dbc.SQLDb.Begin()
		if err == nil {
			//现在数据库事务中删除数据
			_, err = tx.Exec("delete from t_lubopicb WHERE id=?", picid)
			if err == nil {
				savepath := fmt.Sprintf("%s%s", tools.AppPath, webdir)
				_, err = os.Stat(savepath)
				if err == nil {
					err = os.Remove(savepath)
					if err == nil {
						tx.Commit()
					}

				} else if os.IsNotExist(err) {
					tx.Commit()
				} else {
					tx.Rollback()
				}
			} else {
				tx.Rollback()
			}
		}
	}
	return
}

//GetBakArtList 获取后台管理文章列表
func GetBakArtList(uid, index, pagesize string) (data *page.TableJSONData, err error) {
	indexNum, _ := strconv.Atoi(index)
	pagesizeNum, _ := strconv.Atoi(pagesize)
	curpage := (indexNum - 1) * pagesizeNum
	count := 0
	arts := make([]*artciles.Article, 0)
	sqlstr := `SELECT t1.id,t1.name,t4.name AS author,t2.name AS classify ,t3.name AS tag,t1.createtime 
	FROM t_articleb t1 LEFT JOIN t_classifyb t2 ON t1.classify=t2.id 
	LEFT JOIN t_classifyb t3 ON t1.tag=t3.id 
	LEFT JOIN t_userb t4 ON t1.userid=t4.id 
	WHERE t1.userid=?  
	LIMIT ?,?`
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		err = swsqlx.Dbc.SQLDb.Select(&arts, sqlstr, uid, curpage, pagesize)
		wg.Done()
	}()
	go func() {
		err = swsqlx.Dbc.SQLDb.Get(&count, "SELECT COUNT(id) as count FROM t_articleb WHERE userid=?", uid)
		wg.Done()
	}()
	wg.Wait()
	msg := ""
	code := 0
	if err != nil {
		msg = err.Error()
		code = -1
	}
	return &page.TableJSONData{
		Code:  code,
		Count: count,
		Msg:   msg,
		Data:  arts,
	}, err
}
