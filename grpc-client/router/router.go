package router

import (
	"github.com/gin-gonic/gin"
	spider2 "grpc-client/api/spider"
)

func RouterInit() *gin.Engine{
	r := gin.Default()
	spider := r.Group("spider")
	{
		spider.GET("zhihu",spider2.FetchZhiHuHot)
		spider.GET("pp",spider2.FetchPP)
	}
	return r
}
