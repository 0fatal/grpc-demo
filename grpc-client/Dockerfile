FROM golang:alpine AS builder

WORKDIR /app/server

COPY . ./

ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.io"

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-s -w --extldflags "-static -fpic"' .

# CGO_ENABLED 这个参数默认为1，开启CGO。需要指定为0来关闭，因为CGO不支持交叉编译。

# GOOS 和 GOARCH 用来指定要构建的平台为Linux

# 可选参数-ldflags 是编译选项：

# -s -w 去掉调试信息，可以减小构建后文件体积，
# --extldflags "-static -fpic" 完全静态编译[2]，这样编译生成的文件就可以任意放到指定平台下运行，而不需要运行环境配置。

FROM scratch as prod

WORKDIR /app/server

COPY --from=builder /app/server .

EXPOSE 8899

ENTRYPOINT ["./grpc-client"]


