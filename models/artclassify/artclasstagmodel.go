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

//Classify 分类或标签
type Classify struct {
	ID    int    `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Icon  string `db:"icon" json:"icon"`
	Brief string `db:"brief" json:"brief"`
}
