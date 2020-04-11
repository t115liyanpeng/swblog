package user

//User 用户类
type User struct {
	Name      string `json:"name"`
	LoginName string `json:"loginname"`
	PassWord  string `json:"password"`
}

//AddUser 添加用户
func (user *User) AddUser() {

}
