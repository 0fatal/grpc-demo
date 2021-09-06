import grpc
import rpc.spider_pb2 as SpiderStruct
import rpc.spider_pb2_grpc as SpiderGrpc

def run():
    channel = grpc.insecure_channel('localhost:50051')
    stub = SpiderGrpc.NewsServiceStub(channel)
    resp = stub.GetNews(SpiderStruct.SpiderReq())
    for i in resp.news:
        print(i.title)

run()