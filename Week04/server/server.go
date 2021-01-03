package server

import (
	"github.com/gin-gonic/gin"
	"Week04/api"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {

	r := gin.Default()

	r.GET("query",api.QueryPicture)
	r.GET("all",api.AllPicture)

	return r
}

