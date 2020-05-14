package user

import (
	"swblog/swsqlx"
	"swblog/tools"
)

//DbUser 操作数据库 用户表
type DbUser struct {
	ID        string `db:"id" json:"id"`
	Name      string `db:"name" json:"name"`
	LoginName string `db:"loginname" json:"loginname"`
	PassWord  string `db:"password" json:"password"`
	Brief     string `db:"brief" json:"brief"`
	Email     string `db:"email" json:"email"`
}

//UpdateDbUser 更新数据库用户信息
func (dbu *DbUser) UpdateDbUser() (ret bool, msg string) {
	//查询当前密码
	pwd := ""
	err := swsqlx.Dbc.SQLDb.Get(&pwd, "select password from t_userb where id=?", dbu.ID)
	if err != nil {
		return false, "get user pwd faild"
	}
	tx, err := swsqlx.Dbc.SQLDb.Begin()
	if err != nil {
		return false, err.Error()
	}
	sqlstr := `UPDATE  t_userb SET name = ?, loginname = ?,  brief = ?, email = ? WHERE id = ?`
	_, err = tx.Exec(sqlstr, dbu.Name, dbu.LoginName, dbu.Brief, dbu.Email, dbu.ID)
	//如果当前密码md5和数据库密码的md5值是不一样 则需要更新密码
	if pwd != tools.GetMd5Str(dbu.PassWord) {
		_, err = tx.Exec("UPDATE t_userb SET password=? WHERE id = ? ", dbu.PassWord, dbu.ID)
	}

	if err == nil {
		err = tx.Commit()
		if err == nil {
			ret = true
			msg = ""
		} else {
			tx.Rollback()
			ret = false
			msg = err.Error()
		}
	} else {
		tx.Rollback()
		ret = false
		msg = err.Error()
	}
	return
}

//GetUserByID 获取用户信息
func (dbu *DbUser) GetUserByID(id string) error {
	sqlstr := `SELECT id,name,loginname,password,brief,email FROM t_userb WHERE id=?`
	err := swsqlx.Dbc.SQLDb.Get(dbu, sqlstr, id)
	//fmt.Printf("user id is %v\n", dbu.ID)
	return err
}
