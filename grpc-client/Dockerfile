FROM golang:alpine

WORKDIR /app/server

COPY . ./

ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.io"

RUN go mod tidy
RUN go build .
EXPOSE 8899

ENTRYPOINT ["./grpc-client"]
