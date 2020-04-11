package swsqlx

import (
	"fmt"
	"swblog/models/conf"

	//使用mysql第三方
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Dbc mysql 连接
var Dbc *MysqlDbc

// InitedMysql 初始化mysql连接
func initedMysql(dbcfg *conf.Databasecfg) (dbconn *sqlx.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbcfg.User, dbcfg.Password, dbcfg.IPAddress, dbcfg.Port, dbcfg.DbName)
	db, err := sqlx.Open("mysql", dsn)
	return db, err
}

//MysqlDbc mysql 数据库连接实例
type MysqlDbc struct {
	sqlDb *sqlx.DB
}

//NewMysqlDbc 构造函数
func NewMysqlDbc(dbcfg *conf.Databasecfg) (dbc *MysqlDbc, err error) {
	db, err := initedMysql(dbcfg)
	dbc = &MysqlDbc{
		sqlDb: db,
	}
	return dbc, err
}

//CreateDbcInstance 创建dbc实例
func CreateDbcInstance(dbcfg *conf.Databasecfg) {
	if nil == Dbc {
		db, err := initedMysql(dbcfg)
		if err == nil {
			dbc := &MysqlDbc{
				sqlDb: db,
			}
			Dbc = dbc
		}

	}

}
