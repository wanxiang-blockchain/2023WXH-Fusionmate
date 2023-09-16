package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	ReturnCode int         `json:"returnCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode int, errCode int, data interface{}) {
	msg, _ := GetMsg(errCode)
	g.C.JSON(httpCode, Response{
		ReturnCode: errCode,
		Message:    msg,
		Data:       data,
	})
	return
}

func (g *Gin) SuccessResponse(data interface{}) {
	msg, _ := GetMsg(http.StatusOK)
	g.C.JSON(http.StatusOK, Response{
		ReturnCode: SUCCESS,
		Message:    msg,
		Data:       data,
	})
	return
}
