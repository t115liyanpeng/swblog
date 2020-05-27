package artclassify

import (
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

//ClassSimple 类别下拉列表使用
type ClassSimple struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}
