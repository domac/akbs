package core

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/phillihq/akbs/logger"
	"gopkg.in/gorp.v1"
)

var session *Session

var (
	db_host     = ""
	db_port     = ""
	db_database = ""
	db_user     = ""
	db_password = ""
)

//获取session
func OpenSession() *Session {
	if session == nil {
		session = openConnection()
	}
	return session
}

//资源结构
type Session struct {
	Map *gorp.DbMap
}

func openConnection() *Session {
	connect_info := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?autocommit=true", db_user, db_password, db_host, db_port, db_database)
	db, err := sql.Open("mysql", connect_info)
	if err != nil {
		logger.GetLogger().Errorln("mysql创建失败::", err)
		return nil
	}

	//Ping 测试
	if err := db.Ping(); err != nil {
		logger.GetLogger().Errorln("mysql连接失败::", err)
		return nil
	} else {
		logger.GetLogger().Infoln("mysql连接成功!")
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "UTF8"}}
	return &Session{Map: dbmap}
}

func (resource *Session) DB() *sql.DB {
	return resource.Map.Db
}

func (resource *Session) Close() {
	resource.Map.Db.Close()
	resource = nil
}

//语句查询
func (resource *Session) Select(i interface{}, sql string) (interface{}, error) {
	_, err := resource.Map.Select(i, sql)
	if err != nil {
		return nil, err
	}
	return i, nil
}
