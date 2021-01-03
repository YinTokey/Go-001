// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package server

import (
	"Week04/dao"
	"Week04/services"
	"github.com/google/wire"
)

func InitApp() (*services.PictureService) {

	// 配置初始化
	wire.Build(services.NewPictureService,dao.Provider)

	return &services.PictureService{}
}
