package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	SUCCESS = 0
	ERROR = -1
)

func Result(code int, data interface{},msg string, c *gin.Context) {
	c.JSON(http.StatusOK,Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS,map[string]interface{}{},"操作成功",c)
}

func OkWithMessage(msg string,c *gin.Context) {
	Result(SUCCESS,map[string]interface{}{}, msg,c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS,data,"操作成功",c)
}

func FailWithMessage(msg string,c *gin.Context) {
	Result(ERROR,map[string]interface{}{},msg,c)
}

