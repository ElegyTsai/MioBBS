package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type response struct {
	code int         `json:"code"`
	data interface{} `json:"data"`
	msg  string      `json:"msg"`
}

const (
	FAIL    = 0
	SUCCESS = 7
)

func result(c *gin.Context, code int, data interface{}, msg string) {
	c.JSON(http.StatusOK, response{
		code,
		data,
		msg,
	})
}
func Ok(c *gin.Context) {
	result(c, SUCCESS, map[string]interface{}{}, "Success")
}

func OkResponseWithMsg(c *gin.Context, msg string) {
	result(c, SUCCESS, map[string]interface{}{}, msg)
}

func ResponseWithData(c *gin.Context, data interface{}) {
	result(c, SUCCESS, data, "Success")
}

func OkWithAll(c *gin.Context, data interface{}, msg string) {
	result(c, SUCCESS, data, msg)
}

func Fail(c *gin.Context) {
	result(c, FAIL, map[string]interface{}{}, "Fail")
}

func FailResponseWithMsg(c *gin.Context, msg string) {
	result(c, FAIL, map[string]interface{}{}, msg)
}

func FailResponseWithData(c *gin.Context, data interface{}) {
	result(c, FAIL, data, "Fail")
}

func FailWithAll(c *gin.Context, data interface{}, msg string) {
	result(c, FAIL, data, msg)
}
