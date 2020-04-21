package page

import "swblog/swsqlx"

//UserModule 右上角用户信息
type UserModule struct {
	UserName string `db:"username"` //显示的首页的用户的名称
	Brief    string `db:"brief"`    //个性签名
	Email    string `db:"email"`    //邮件地址
	LogNum   int    `db:"lognum"`   //文章总数
	Classify int    `db:"classify"` //分类数
	Tags     int    `db:"tags"`     //标签数

}

//FirstPage 首页对应实体类
type FirstPage struct {
	Title    string      //标题
	UserInfo *UserModule //用户模块信息
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
	} else {
		return nil
	}
}
