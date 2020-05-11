package controllers

import "swblog/models/page"

//GetBgIndex 获取后台默认页的数据
func GetBgIndex(uid string) *page.BackageStage {
	return &page.BackageStage{
		ArtCount:     10,
		CommentCount: 9,
		AllSee:       1000,
		TodaySee:     2,
	}
}
