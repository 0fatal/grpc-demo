package client

import (
	"context"
	"google.golang.org/grpc"
	"grpc-client/config"
	"grpc-client/rpc/pp"
	"grpc-client/rpc/zhihu"
	"log"
	"strconv"
)


type Client struct {
	option      Options
	Conn        *grpc.ClientConn
	PPClient    pp.PPSpiderClient
	ZhiHuClient zhihu.ZhiHuSpiderClient
}

type Options struct {
	Port        int
	Host 		string
}

type Option func(*Options)

func Port(port int)Option {
	return func(this *Options) {
		this.Port = port
	}
}

func Host(host string)Option {
	return func(this *Options) {
		this.Host = host
	}
}

func newOptions(opts ...Option) Options{
	opt := Options{
		Port: config.AllConfig.Grpc.Port,
		Host: config.AllConfig.Grpc.Host,
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

func (c *Client) Init(opts ...Option) (err error) {
	c.option = newOptions(opts...)

	conn, err := grpc.Dial(c.option.Host+":"+strconv.Itoa(c.option.Port), grpc.WithInsecure())
	if err != nil {
		return err
	}
	c.Conn = conn

	c.ConnectToClient()
	return
}

func (c *Client)ConnectToClient() {
	c.PPClient = pp.NewPPSpiderClient(c.Conn)
	c.ZhiHuClient = zhihu.NewZhiHuSpiderClient(c.Conn)
}

func (c *Client) Close() (err error) {
	err = c.Conn.Close()
	return
}

func (c *Client) FetchPPW(req *pp.PPSpiderReq) (resp *pp.PPSpiderResp,err error){
	resp, err = c.PPClient.FetchPPW(context.Background(),req)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func (c *Client) FetchZhiHuHot()(resp *zhihu.ZhiHuSpiderResp,err error) {
	resp, err = c.ZhiHuClient.FetchHot(context.Background(), &zhihu.ZhiHuSpiderReq{})
	if err != nil {
		log.Println(err)
		return
	}
	return
}
