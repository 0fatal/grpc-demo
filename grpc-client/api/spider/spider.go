package spider

import (
	"github.com/gin-gonic/gin"
	"grpc-client/api/response"
	"grpc-client/rpc/pp"
	"grpc-client/service"
	"grpc-client/utils"
	"log"
	"strconv"
)

func FetchZhiHuHot(c *gin.Context){
	res, err := service.SSpider.FetchZhiHuHot()
	if err != nil {
		log.Println(err)
		response.FailWithMessage(err.Error(),c)
		return
	}
	response.OkWithData(res.Data, c)
}

func FetchPP(c *gin.Context) {
	var sectionType pp.SectionType
	var newsType pp.NewsType

	page,err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}

	var tmp int

	utils.ProtectRun(func() {
		tmp,err = strconv.Atoi(c.Query("section_type"))
		if err != nil || pp.SectionType_name[int32(tmp)] == ""  {
			panic(err)
		}
		sectionType = pp.SectionType(tmp)
	}, func(){
		sectionType = pp.SectionType_ST_NEWS
	})

	utils.ProtectRun(func() {
		tmp,err = strconv.Atoi(c.Query("news_type"))
		if err != nil || pp.NewsType_name[int32(tmp)] == "" {
			panic(err)
		}
		newsType = pp.NewsType(tmp)
	}, func(){
		newsType = pp.NewsType_NT_ALL
	})

	log.Println(page, sectionType, newsType)
	res, err := service.SSpider.FetchPPW(page,sectionType,newsType)

	if err != nil {
		log.Println(err)
		response.FailWithMessage(err.Error(),c)
		return
	}
	response.OkWithData(res.Data, c)
}