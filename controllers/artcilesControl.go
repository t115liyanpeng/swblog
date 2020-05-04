package controllers

import (
	"swblog/models/artciles"
	"swblog/swsqlx"
)

//GetArticleDetail 根据文章id获取文章详情
func GetArticleDetail(aid, uid string) *artciles.Article {
	artdetail := artciles.Article{}
	err := swsqlx.Dbc.SQLDb.Get(&artdetail, `SELECT t1.id,t1.name,t1.content as subsummary,t1.up,t1.ulike ,
	t1.click,t2.name as classify,t3.name as tag,t4.name as author,t1.createtime FROM t_articleb t1 inner JOIN t_classifyb t2
	 ON t1.userid=t2.userid inner JOIN t_classifyb t3 ON t1.userid=t3.userid INNER JOIN t_userb t4 ON t1.userid=t4.id 
	 WHERE t2.id=t3.pid  and t1.tag=t3.id and t1.id=? and t1.userid=?`, aid, uid)
	if err == nil {
		return &artdetail
	}
	return nil
}

//GetTimeLineArticle 获取时间线数据 先不做分页
func GetTimeLineArticle(uid string, pagesize, pageindex int) *artciles.ArtTimeLine {

	timeline := []*artciles.TimeLine{}
	curpage := (pageindex - 1) * pagesize
	count := 0
	err := swsqlx.Dbc.SQLDb.Select(&timeline, "SELECT id,name,createtime FROM t_articleb WHERE userid=? ORDER BY createtime DESC  LIMIT ?,?", uid, curpage, pagesize)
	err = swsqlx.Dbc.SQLDb.Get(&count, "SELECT COUNT(id) as count FROM t_articleb WHERE userid=?", uid)
	if err == nil {
		arts := artciles.ArtTimeLine{
			LineItems: timeline,
			ArtCount:  count,
			PageSize:  pagesize,
		}
		return &arts
	}
	return nil
}
