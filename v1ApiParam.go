package c_code

import "github.com/gin-gonic/gin"

func V1GinSuccess(data interface{}, message ...string) gin.H {
	me := "OK"
	u := ""
	if len(message) >= 1 {
		me = message[0]
	}
	if len(message) == 2 {
		u = message[1]
	}
	return gin.H{
		"code":    1,
		"data":    data,
		"message": me,
		"u":       u,
	}
}

func V1GinError(code int, message ...string) gin.H {
	me := "一个错误的请求"
	u := ""
	data := ""
	if len(message) >= 1 {
		me = message[0]
	}
	if len(message) == 2 {
		data = message[0]
		u = message[1]
	}
	return gin.H{
		"code":    code,
		"data":    data,
		"message": me,
		"u":       u,
	}
}
