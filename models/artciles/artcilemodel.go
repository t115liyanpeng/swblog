package artciles

import "html/template"

//Article 文章
type Article struct {
	ID          int           `db:"id" json:"id"`              //id
	Name        string        `db:"name" json:"name"`          //文章名称
	Content     string        `db:"subsummary" json:"content"` //文章摘要 或主体
	ContentHTML template.HTML //文章html
	Top         bool          `db:"up" json:"up"`                 //是否置顶
	Like        int           `db:"ulike" json:"ulike"`           //喜欢
	Click       int           `db:"click" json:"click"`           //点击次数
	Classify    string        `db:"classify" json:"classify"`     //分类
	Tag         string        `db:"tag" json:"tag"`               //标签
	Author      string        `db:"author" json:"author"`         //作者
	CreateTime  string        `db:"createtime" json:"createtime"` //创建时间
	//Comid      int    `db:"commentid" json:"commentid"`   //评论id
}

//ArticleHTML 文章html处理
type ArticleHTML struct {
	ID         int           //id
	Name       string        //文章名称
	Content    template.HTML //文章摘要 或主体
	Top        bool          //是否置顶
	Like       int           //喜欢
	Click      int           //点击次数
	Classify   string        //分类
	Tag        string        //标签
	Author     string        //作者
	CreateTime string        //创建时间
	Comid      int           //评论id
}

//ArtSummary 文章汇总标题
type ArtSummary struct {
	ID    int    `db:"id"`    //id
	Num   int    `db:"num"`   //文章序号
	Title string `db:"title"` //文章标题
}

//TimeLine 文章时间线item
type TimeLine struct {
	ID         int    `db:"id"`         //文章id
	Name       string `db:"name"`       //文章名称
	CreateTime string `db:"createtime"` //创建时间
}

//ArtTimeLine 文章时间线
type ArtTimeLine struct {
	LineItems []*TimeLine //时间线items
	ArtCount  int         //总条数
	PageSize  int         //每页大小
}

//ArtListParam 参数
type ArtListParam struct {
	Classify string
	Tag      string
	Name     string
}

//ArtList 文章列表
type ArtList struct {
	Title    string        //分类标题
	List     []*Article    // 文章列表
	ArtCount int           //总条数
	PageSize int           //每页大小
	Param    *ArtListParam //参数
}
