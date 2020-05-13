package user

import (
	"fmt"
	"swblog/swsqlx"
)

//DbUser 操作数据库 用户表
type DbUser struct {
	ID        string `db:"id"`
	Name      string `db:"name"`
	LoginName string `db:"loginname"`
	PassWord  string `db:"password"`
	Brief     string `db:"brief"`
	Email     string `db:"email"`
}

//UpdateDbUser 更新数据库用户信息
func (dbu *DbUser) UpdateDbUser() (ret bool, msg string) {
	tx, err := swsqlx.Dbc.SQLDb.Begin()
	if err != nil {
		return false, err.Error()
	}
	sqlstr := `UPDATE  t_userb SET name = ?, loginname = ?,  brief = ?, email = ? WHERE id = ?`
	_, err = tx.Exec(sqlstr, dbu.Name, dbu.LoginName, dbu.Brief, dbu.Email, dbu.ID)
	_, err = tx.Query("select password from t_userb where id=?", dbu.ID)

	if err == nil {
		ret = true
		msg = ""
	} else {
		ret = false
		msg = err.Error()
	}
	return
}

//GetUserByID 获取用户信息
func (dbu *DbUser) GetUserByID(id string) error {
	sqlstr := `SELECT name,loginname,password,brief,email FROM t_userb WHERE id=?`
	err := swsqlx.Dbc.SQLDb.Get(dbu, sqlstr, id)
	fmt.Printf("get user errr %v\n", err)
	return err
}
