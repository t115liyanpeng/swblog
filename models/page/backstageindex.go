package page

//BackageStage 后台访问
type BackageStage struct {
	ArtCount     int //文章总数
	CommentCount int //评论总数
	AllSee       int //总访问量
	TodaySee     int //今日访问
}

//TableJSONData 表json通用结构体
type TableJSONData struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Count int         `json:"count"`
	Data  interface{} `json:"data"`
}
