package user

import (
	"swblog/swsqlx"
)

//LoginInfo 登录时使用
type LoginInfo struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

//User 用户类
type User struct {
	Name      string      `json:"username" db:"name"`
	LoginName string      `json:"loginname" db:"loginname" form:"username"`
	PassWord  string      `json:"password" db:"password" form:"password"`
	CheckSvae bool        `json:"chksave" form:"chksave"`
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
	var tUser User = User{
		Name:      "",
		LoginName: "",
		PassWord:  "",
		CheckSvae: false,
		State: &LoginState{
			State: false,
			Msg:   "",
			Code:  -1,
		},
	}

	err := swsqlx.Dbc.SQLDb.Get(&tUser, "select name,loginname,password from t_userb where loginname=?", user.LoginName)

	if err != nil {
		user.State.State = false
		user.State.Code = 1
		user.State.Msg = "用户未注册！"
		return err
	}

	if tUser.LoginName == user.LoginName {
		if tUser.PassWord == user.PassWord {
			user.Name = tUser.Name
			user.State.State = true
			user.State.Code = 0
		} else {
			user.State.State = false
			user.State.Code = 2
			user.State.Msg = "密码不正确！"
		}
	}

	return nil
}
