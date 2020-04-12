package user

import (
	"swblog/swsqlx"
)

//User 用户类
type User struct {
	Name      string      `json:"name" db:"name"`
	LoginName string      `json:"loginname" db:"loginname" form:"name"`
	PassWord  string      `json:"password" db:"password" form:"password"`
	CheckSvae bool        `form:"chksave"`
	State     *LoginState `json:"loginstate"`
}

//LoginState 用户登录状态
type LoginState struct {
	State bool   `json:"state"`
	Msg   string `json:"msg"`
	Code  int    `json:"code"`
}

//AddUser 添加用户
func (user *User) AddUser() bool {
	return false
}

//UserLogin 用户登录
func (user *User) UserLogin() error {
	var tUser User
	err := swsqlx.Dbc.SQLDb.Get(&tUser, "select name,loginname,password from t_userb where name=?", user.Name)
	if err != nil {
		return err
	}
	if tUser.Name == "" {
		user.State.State = false
		user.State.Code = 1
		user.State.Msg = "user not found"
	} else if tUser.Name == user.Name {
		if tUser.PassWord == user.PassWord {
			user.State.State = true
			user.State.Code = 0
		} else {
			user.State.State = false
			user.State.Code = 2
			user.State.Msg = "password error"
		}
	}

	return nil
}
