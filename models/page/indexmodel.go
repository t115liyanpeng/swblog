package page

import (
	"database/sql"
	"fmt"
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
	Summary string     //汇总类
	Icon    string     //图标
	Sub     []*SubTags //子分类
}

//FirstPage 首页对应实体类
type FirstPage struct {
	Title    string      //标题
	UserInfo *UserModule //用户模块信息
	Left     []*LeftTags //左侧导航
}

//GetWebSietUserInfo 获取网站用户模块的信息
func GetWebSietUserInfo() *UserModule {
	var um UserModule = UserModule{
		UserName: "",
		Brief:    "",
		Email:    "",
		LogNum:   0,
		Classify: 0,
		Tags:     0,
	}

	err := swsqlx.Dbc.SQLDb.Get(&um, "select username,brief,email,lognum,classify,tags from t_websiteb")
	if err == nil {
		return &um
	}
	return nil

}

//查询数据库中的分类数据
func getLeftDataSource() []*TreeSource {
	data := []*TreeSource{}
	err := swsqlx.Dbc.SQLDb.Select(&data, "select id,pid,name,link,icon from t_classifyb")
	if err != nil {
		fmt.Printf("err : %v\n", err)
		return nil
	}
	for i, v := range data {
		fmt.Printf("i=%d,v=%s\n", i, v.Name)
	}
	return data
}

//GetLeftNavData 获取左侧导航数据
func GetLeftNavData() []*LeftTags {
	source := getLeftDataSource()
	var tags []*LeftTags = make([]*LeftTags, 0)
	//只生成二级树
	for _, v := range source {
		lt := LeftTags{
			Summary: "",
			Sub:     make([]*SubTags, 0),
		}
		if v.ID == 0 {
			lt.Summary = v.Name
			for _, v2 := range source {
				if v.ID == v2.Pid {
					sub := SubTags{
						NavName: v2.Name,
						Link:    v2.Link.String,
					}
					lt.Sub = append(lt.Sub, &sub)
				}

			}
		}
		tags = append(tags, &lt)
	}

	return tags
}
