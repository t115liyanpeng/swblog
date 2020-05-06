package controllers

import (
	"fmt"
	"swblog/models/artciles"
	"swblog/models/page"
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

//GetArtList 获取文章列表
func GetArtList(pageindex, pagesize int, classify, tag, uid, name string) *artciles.ArtList {
	curpage := (pageindex - 1) * pagesize
	items := []*artciles.Article{}
	count := 0
	var artsql string = fmt.Sprintf(`SELECT t1.id,t1.name,LEFT(t1.content,200) as subsummary,up,t1.ulike ,
	t1.click,t2.name as classify,t3.name as tag,t4.name as author,t1.createtime FROM t_articleb t1 inner JOIN t_classifyb t2
	 ON t1.userid=t2.userid inner JOIN t_classifyb t3 ON t1.userid=t3.userid INNER JOIN t_userb t4 ON t1.userid=t4.id 
	 WHERE t2.id=t3.pid  and t1.tag=t3.id and t1.userid='%s' `, uid)
	var artcount string = fmt.Sprintf(`SELECT COUNT(t1.id) as count FROM t_articleb t1 LEFT JOIN t_classifyb t2 
	ON t1.classify=t2.id LEFT JOIN t_classifyb t3 ON t1.tag =t3.id WHERE t1.userid='%s'`, uid)
	if classify != "" {
		artsql = fmt.Sprintf("%s and t2.id=%s ", artsql, classify)
		artcount = fmt.Sprintf("%s and t2.id=%s ", artcount, classify)
	}
	if tag != "" {
		artsql = fmt.Sprintf("%s and t3.id=%s ", artsql, tag)
		artcount = fmt.Sprintf("%s and t3.id=%s ", artcount, tag)
	}
	if name != "" {
		artsql = fmt.Sprintf("%s and t1.name like '%s"+"%s"+"%s' ", artsql, "%", name, "%")
		artcount = fmt.Sprintf("%s and t1.name like '%s"+"%s"+"%s' ", artcount, "%", name, "%")
	}

	artsql = fmt.Sprintf("%s order by createtime desc LIMIT %d,%d", artsql, curpage, pagesize)
	err := swsqlx.Dbc.SQLDb.Select(&items, artsql)
	err = swsqlx.Dbc.SQLDb.Get(&count, artcount)
	if err == nil {
		list := artciles.ArtList{
			List:     items,
			ArtCount: count,
			PageSize: pagesize,
		}
		return &list
	}
	return nil
}

//GetClassOrTags 获取分类或标签
func GetClassOrTags(pageindex, pagesize, classOrtag int, uid string) (data []*page.TreeSource, count int) {
	curpage := (pageindex - 1) * pagesize
	data = []*page.TreeSource{}
	count = 1
	sqlstr := ""
	sqlstrcnt := ""
	if classOrtag == 0 {
		sqlstr = "select id,pid,name,link,icon from t_classifyb where userid=? and pid=0 LIMIT ?,?"
		sqlstrcnt = "select count(id) as count from t_classifyb where userid=? and pid=0"
	} else {
		sqlstr = "select id,pid,name,link,icon from t_classifyb where userid=? and pid!=0 LIMIT ?,?"
		sqlstrcnt = "select count(id) as count from t_classifyb where userid=? and pid!=0"
	}
	err := swsqlx.Dbc.SQLDb.Select(&data, sqlstr, uid, curpage, pagesize)
	err = swsqlx.Dbc.SQLDb.Get(&count, sqlstrcnt, uid)
	if err == nil {
		return data, count
	}
	return nil, 1
}
