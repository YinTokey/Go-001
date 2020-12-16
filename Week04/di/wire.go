package di

import (
	"Week04/dao"
	"github.com/google/wire"
)


func InitApp() ( func(), error) {

	// 配置初始化
	panic(wire.Build(dao.Provider))
}
