package core

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/phillihq/akbs/logger"
	"gopkg.in/gorp.v1"
)

var session *Session

func init() {
	logger.GetLogger().Infoln("init session")
	session = openConnection()
}

//获取session
func OpenSession() *Session {
	if session == nil {
		logger.GetLogger().Infoln("reopen the session")
		session = openConnection()
	}
	return session
}

//资源结构
type Session struct {
	Map *gorp.DbMap
}

func openConnection() *Session {
	db, _ := sql.Open("mysql", "mydb@/akbs")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "UTF8"}}
	return &Session{Map: dbmap}
}

func (resource *Session) DB() *sql.DB {
	return resource.Map.Db
}

//语句查询
func (resource *Session) Select(i interface{}, sql string) (interface{}, error) {
	_, err := resource.Map.Select(i, sql)
	if err != nil {
		return nil, err
	}
	return i, nil
}
