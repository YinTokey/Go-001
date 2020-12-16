package conf

import (
	"Week04/dao"
)

// Init 初始化配置项
func Init() error {

	// 连接数据库
	env := "root:12345678@tcp(127.0.0.1:3306)/week2_db?charset=utf8&parseTime=True&loc=Local"
	return dao.Database(env)

}
