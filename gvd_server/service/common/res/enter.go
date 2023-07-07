/*
 * @Author: zzzzztw
 * @Date: 2023-07-05 21:37:15
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-07-08 00:18:31
 * @FilePath: /gin-vue-docs/gvd_server/service/common/res/enter.go
 */
package res

import (
	"gvd_server/utils/valid"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Code int

const (
	Success   Code = 0
	ErrorCode Code = 7 //系统错误
	ValidCode Code = 9 // 校验错误
)

type Response struct {
	Code Code   `json:"code"`
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

func Fail(code Code, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func FailWithMsg(msg string, c *gin.Context) {
	Fail(ErrorCode, map[string]any{}, msg, c)
}

func FailWithError(err error, obj any, c *gin.Context) {
	errmsg := valid.Error(err)
	FailWithMsg(errmsg, c)
}

func FailWithData(data any, c *gin.Context) {
	Fail(ErrorCode, data, "系统错误", c)
}

func FailWithValidError(err error, obj any, c *gin.Context) {

}
