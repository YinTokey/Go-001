package dao

import (
	"database/sql"
	"github.com/pkg/errors"

	"github.com/google/wire"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB 数据库链接单例
var sqlDB *sql.DB

var Provider = wire.NewSet(NewDB,NewDao)

type dao struct {
	db *sql.DB
}

func NewDao(db *sql.DB) (*dao, error)  {
	return &dao{db: db}, nil
}

func NewDB() (db *sql.DB, cf func(), err error) {
	// 连接数据库
	env := "root:12345678@tcp(127.0.0.1:3306)/week2_db?charset=utf8&parseTime=True&loc=Local"

	db, err = sql.Open("mysql",env)

	cf = func() {db.Close()}

	return
}

func Database(connString string) error {

	openDB, err := sql.Open("mysql",connString)

	// Error
	if err != nil {
		return errors.Wrap(err,"mysql connect failed")
	}

	sqlDB = openDB

	return nil
}

func InstanceDB() *sql.DB {
	return sqlDB
}
