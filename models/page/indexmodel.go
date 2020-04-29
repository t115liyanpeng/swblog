package page

import (
	"database/sql"
	"swblog/swsqlx"
)

//TreeSource 树行结构的数据源
type TreeSource struct {
	ID   int            `db:"id"`
	Pid  int            `db:"pid"`
	Name string         `db:"name"`
	Link sql.NullString `db:"link"`
	Icon sql.NullString `db:"icon"`
}

//UserModule 右上角用户信息
type UserModule struct {
	UserName string `db:"username"` //显示的首页的用户的名称
	Brief    string `db:"brief"`    //个性签名
	Email    string `db:"email"`    //邮件地址
	LogNum   int    `db:"lognum"`   //文章总数
	Classify int    `db:"classify"` //分类数
	Tags     int    `db:"tags"`     //标签数

}

//SubTags 子分类
type SubTags struct {
	NavName string
	Link    string
}

//LeftTags 左侧导航大类
type LeftTags struct {
	ID      int
	Summary string     //汇总类
	Icon    string     //图标
	Sub     []*SubTags //子分类
}

//Article 文章
type Article struct {
	ID         int    `db:"id"`         //id
	Name       string `db:"name"`       //文章名称
	Content    string `db:"subsummary"` //文章摘要
	Top        bool   `db:"up"`         //是否置顶
	Like       int    `db:"like"`       //喜欢
	Click      int    `db:"click"`      //点击次数
	Classify   string `db:"classify"`   //分类
	Tag        string `db:"tag"`        //标签
	Author     string `db:"author"`     //作者
	CreateTime string `db:"createtime"` //创建时间
}

//FirstPage 首页对应实体类
type FirstPage struct {
	Title    string      //标题
	UserInfo *UserModule //用户模块信息
	Left     []*LeftTags //左侧导航
	Articles []*Article  //文章
}

//GetWebSietUserInfo 获取网站用户模块的信息
func GetWebSietUserInfo(userid string) *UserModule {
	var um UserModule = UserModule{
		UserName: "",
		Brief:    "",
		Email:    "",
		LogNum:   0,
		Classify: 0,
		Tags:     0,
	}
	//没有办法了 好久没写过sql 自己写不出来了 只好用了 子查询
	err := swsqlx.Dbc.SQLDb.Get(&um, `SELECT  t1.name as username,t1.brief,t1.email,
	(SELECT COUNT(id) FROM  t_classifyb WHERE pid=0 and userid=?) as classify,
	 (SELECT COUNT(id) FROM t_classifyb WHERE pid!=0 AND userid=?) as tags ,
	 (SELECT COUNT(id) FROM t_articleb WHERE userid=?)as lognum FROM t_userb t1`, userid, userid, userid)
	if err == nil {
		//fmt.Printf("user info  %v\n", um)
		return &um
	}
	//fmt.Printf("user err %v\n", err)
	return nil

}

//查询数据库中的分类数据
func getLeftDataSource(userid string) []*TreeSource {
	data := []*TreeSource{}
	err := swsqlx.Dbc.SQLDb.Select(&data, "select id,pid,name,link,icon from t_classifyb where userid=?", userid)
	if err != nil {
		//fmt.Printf("err : %v\n", err)
		return nil
	}
	return data
}

//GetLeftNavData 获取左侧导航数据
func GetLeftNavData(uid string) []*LeftTags {
	source := getLeftDataSource(uid)
	var tags []*LeftTags = []*LeftTags{}
	//只生成二级树
	for _, v := range source {
		lt := &LeftTags{}
		if v.Pid == 0 {
			lt.ID = v.ID
			lt.Icon = v.Icon.String
			lt.Summary = v.Name
			lt.Sub = []*SubTags{}
			tags = append(tags, lt)
		}

	}

	for _, v := range source {
		for _, s := range tags {
			if v.ID != 0 && v.Pid == s.ID {
				sub := &SubTags{
					NavName: v.Name,
					Link:    v.Link.String,
				}
				s.Sub = append(s.Sub, sub)
			}
		}
	}
	return tags
}

//GetContent 获取文章主体
func GetContent(uid string) []*Article {
	articles := []*Article{}
	err := swsqlx.Dbc.SQLDb.Select(&articles, `SELECT t1.id,t1.name,LEFT(t1.content,200) as subsummary,up,t1.like ,
	t1.click,t2.name as classify,t3.name as tag,t4.name as author,t1.createtime FROM t_articleb t1 inner JOIN t_classifyb t2
	 ON t1.userid=t2.userid inner JOIN t_classifyb t3 ON t1.userid=t3.userid INNER JOIN t_userb t4 ON t1.userid=t4.id 
	 WHERE t2.id=t3.pid  and t1.tag=t3.id and t1.userid=?`, uid)
	if err != nil {
		return nil
	}
	return articles
}
