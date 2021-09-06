import rpc.pp.pp_spider_pb2_grpc as PPSpiderGrpc
import spider.pp_spider as PPSpider
import rpc.pp.pp_spider_pb2 as PPSpiderD

import rpc.zhihu.zhihu_spider_pb2_grpc as ZhiHuSpiderGrpc
import spider.zhihu_spider as ZhiHuSpider
import rpc.zhihu.zhihu_spider_pb2 as ZhiHuSpiderD

import grpc
from concurrent import futures
import time

class PPSpiderServer(PPSpiderGrpc.PPSpiderServicer):
    def FetchPPW(self, request, context):
        page = request.page if request.page > 0 else 1
        section_type = PPSpider.SectionType(request.section_type) if request.section_type != 0 else PPSpider.SectionType.NEWS
        news_type = PPSpider.NewsType(request.news_type)
        print(PPSpider.fetchPPW(page,section_type,news_type))

        return PPSpiderD.PPSpiderResp(data=PPSpider.fetchPPW(page,section_type,news_type))

class ZhiHuSpiderServer(ZhiHuSpiderGrpc.ZhiHuSpiderServicer):
    def FetchHot(self, request, context):
        return ZhiHuSpiderD.ZhiHuSpiderResp(data=ZhiHuSpider.getZhiHuHot())

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    PPSpiderGrpc.add_PPSpiderServicer_to_server(PPSpiderServer(),server)
    ZhiHuSpiderGrpc.add_ZhiHuSpiderServicer_to_server(ZhiHuSpiderServer(),server)
    server.add_insecure_port('[::]:50051')
    server.start()
    try:
        while True:
            time.sleep(60 * 60 * 24)  # one day in seconds
    except KeyboardInterrupt:
        server.stop(0)