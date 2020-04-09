package swsqlx

import (
	"database/sql"
	"fmt"
	"swblog/models/conf"

	//使用mysql第三方
	_ "github.com/go-sql-driver/mysql"
)

// Dbc mysql 连接
var Dbc *sql.Conn

// InitedMysql 初始化mysql连接
func InitedMysql(dbcfg *conf.Databasecfg) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbcfg.User, dbcfg.Password, dbcfg.IPAddress, dbcfg.Port, dbcfg.DbName)
	Dbc, err := sql.Open("mysql", dsn)
	return err
}
