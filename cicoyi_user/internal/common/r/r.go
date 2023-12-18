package r

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
)

const (
	SUCCESS = iota
	ERROR
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func OK(msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		SUCCESS,
		"",
		msg,
	})
}

func OKWithData(data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		SUCCESS,
		data,
		msg,
	})
}

func Failed(msg string, c *gin.Context) {
	c.JSON(http.StatusBadRequest, Response{
		ERROR,
		"",
		msg,
	})
}

func FailedWithError(msg string, err error, c *gin.Context) {
	c.JSON(http.StatusBadRequest, Response{
		ERROR,
		err,
		msg,
	})
}
