> Exception calling application: 'ascii' codec can't encode characters in ...
- 设置环境变量 `ENV PYTHONIOENCODING=utf-8`
> pipreqs ./ --encoding=utf-8
- 生成requirements.txt文件
> docker容器中，下载fake_useragent被墙
- 手动下载json文件，挂载到 /tmp
> docker exec -it grpc-test
> 
> docker build -t grpc-server .
> 
> docker run -it --name grpc-test grpc-server
- 常用docker命令
> ubuntu16自带python3.5, 18自带py3.6.9
---
> apt-get install inetutils-ping
- ping工具
>  python -m grpc_tools.protoc --python_out=. --grpc_python_out=. -I spider.proto
- python grpc 生成