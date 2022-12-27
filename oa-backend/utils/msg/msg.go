package msg

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SUCCESS = 1
	FAIL    = 2
	ERROR   = -1
)

var codeMsg = map[int]string{
	SUCCESS: "请求成功！",
	FAIL:    "请求失败！请检查请求参数。",
	ERROR:   "请求异常！请联系管理员。",
}

func GetMsg(code int) string {
	return codeMsg[code]
}

func Message(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": GetMsg(code),
		"data":    data,
	})
}
