package artclassify

import (
	"swblog/models/artciles"
	"swblog/models/page"
)

//ArtClassTag 分类和标签
type ArtClassTag struct {
	Title      string             //标题名称
	List       []*page.TreeSource //数据
	ArtCount   int                //总条数
	PageSize   int                //每页大小
	ClassOrTag int                //查询分类还是tag 0：分类 1：tag
}

//Classify 分类
type Classify struct {
	ID    int    `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Icon  string `db:"icon" json:"icon"`
	Brief string `db:"brief" json:"brief"`
}

//Tags 标签
type Tags struct {
	ID        int    `db:"id" json:"id"`
	PID       int    `db:"pid" json:"pid"`
	Name      string `db:"name" json:"name"`
	ClassName string `db:"classname" json:"classname"`
	Brief     string `db:"brief" json:"brief"`
}

//ClassSimple 下拉列表使用
type ClassSimple struct {
	ID   string `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

//UserSimple 下拉列表使用
type UserSimple struct {
	ID   string `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

//AddArtPage 添加文章页的模型
type AddArtPage struct {
	ClassDropList []*ClassSimple //分类下拉列表
	UserDropList  []*UserSimple  //用户下拉列表

}

//EditArtData 更新文章页面打开的时候需要的数据
type EditArtData struct {
	ArtInfo      *artciles.ArticleEdit //文章信息
	ClassifyDrop []*ClassSimple        //分类下拉
	TagDrop      []*ClassSimple        //标签下拉
	UserDrop     []*UserSimple         //作者下拉

}
