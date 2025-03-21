package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Code int         `json:"code"`           // 状态码
	Msg  string      `json:"msg"`            // 提示信息
	Data interface{} `json:"data,omitempty"` // 返回数据（可选）
}

// JSON 是最底层响应封装方法
func JSON(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, Result{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

// Success 成功响应（含数据）
func Success(c *gin.Context, data interface{}) {
	JSON(c, CodeSuccess, "success", data)
}

// Fail 失败响应（含错误码与信息）
func Fail(c *gin.Context, code int, msg string) {
	JSON(c, code, msg, nil)
}
