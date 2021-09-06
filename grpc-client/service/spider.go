package service

import (
	"grpc-client/client"
	"grpc-client/rpc/pp"
	"grpc-client/rpc/zhihu"
)

type SpiderService struct {

}


func (ss *SpiderService) FetchZhiHuHot() (resp *zhihu.ZhiHuSpiderResp,err error){
	return client.MyClient.FetchZhiHuHot()
}

func (ss *SpiderService) FetchPPW(page int, sectionType pp.SectionType, newsType pp.NewsType)(resp *pp.PPSpiderResp,err error) {
	return client.MyClient.FetchPPW(&pp.PPSpiderReq{
		Page: uint32(page),
		SectionType: sectionType,
		NewsType: newsType,
	})
}
