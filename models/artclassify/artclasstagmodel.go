package artclassify

import "swblog/models/page"

//ArtClassTag 分类和标签
type ArtClassTag struct {
	Title      string             //标题名称
	List       []*page.TreeSource //数据
	ArtCount   int                //总条数
	PageSize   int                //每页大小
	ClassOrTag int                //查询分类还是tag 0：分类 1：tag
}
