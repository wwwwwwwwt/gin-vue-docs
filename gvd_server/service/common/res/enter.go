/*
 * @Author: zzzzztw
 * @Date: 2023-07-05 21:37:15
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-07-05 22:15:37
 * @FilePath: /gvd_server/service/common/res/enter.go
 */
package res

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	Success = 0
	Failed  = 7
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

type ListResponse[T any] struct {
	List  []T `json:"List"`
	Count int `json:"count"`
}

func OK(data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: Success,
		Data: data,
		Msg:  msg,
	})
}

//响应msg
func OKWithMsg(msg string, c *gin.Context) {
	OK(map[string]any{}, msg, c)
}

//响应数据
func OKWithData(data any, c *gin.Context) {
	OK(data, "查询成功", c)
}

//响应detail
func OKWithDetail[T any](List []T, count int, c *gin.Context) {

	if len(List) == 0 {
		List = []T{}
	}

	OK(ListResponse[T]{
		List:  List,
		Count: count,
	}, "成功", c)
}

//响应lsit
func OKWithList() {

}

func Fail(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func FailWithMsg(msg string, c *gin.Context) {
	Fail(Failed, map[string]any{}, msg, c)
}

func FailWithError(msg string, obj any, c *gin.Context) {

}

func FailWithData(data any, c *gin.Context) {
	Fail(Failed, data, "系统错误", c)
}
