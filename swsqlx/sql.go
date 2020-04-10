package swsqlx

import (
	"database/sql"
	"fmt"
	"swblog/models/conf"

	//使用mysql第三方
	_ "github.com/go-sql-driver/mysql"
)

// Dbc mysql 连接
var Dbc *MysqlDbc

// InitedMysql 初始化mysql连接
func initedMysql(dbcfg *conf.Databasecfg) (dbconn *sql.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbcfg.User, dbcfg.Password, dbcfg.IPAddress, dbcfg.Port, dbcfg.DbName)
	db, err := sql.Open("mysql", dsn)
	return db, err
}

//MysqlDbc mysql 数据库连接实例
type MysqlDbc struct {
	sqlDb *sql.DB
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
func CreateDbcInstance(dbcfg *conf.Databasecfg) *MysqlDbc {
	if nil == Dbc {
		db, err := initedMysql(dbcfg)
		if err == nil {
			dbc := &MysqlDbc{
				sqlDb: db,
			}
			Dbc = dbc
		}

	}
	return Dbc

}

//ExecNonQuery 无返回
func (dbc *MysqlDbc) ExecNonQuery() {
	if dbc != nil {
		dbc.sqlDb.Ping()
	}
}
