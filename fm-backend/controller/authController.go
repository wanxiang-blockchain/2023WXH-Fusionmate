package controller

import (
	"encoding/json"
	"github.com/FusionMate/fm-backend/common"
	"github.com/FusionMate/fm-backend/common/log"
	"github.com/FusionMate/fm-backend/service"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type loginRequest struct {
	Address   string `json:"address"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
	Signature string `json:"signature"`
}

type Auth struct {
}

func (a *Auth) Login(c *gin.Context) {
	appG := common.Gin{C: c}
	var bodyBytes []byte
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("[Auth Controller][Login] Read request body error: %s", err)
		appG.Response(http.StatusInternalServerError, common.INTERNAL_ERROR, nil)
		return
	}
	var request loginRequest
	err = json.Unmarshal(bodyBytes, &request)
	if err != nil {
		log.Info("[Auth Controller][Login] Json unmarshal error: %s", err)
		appG.Response(http.StatusOK, common.INVALID_PARAM, nil)
		return
	}
	token, rtCode, err := service.Web3Login(request.Address, request.Message, request.Timestamp, request.Signature)
	if err != nil {
		log.Error("[Auth Controller][Login] Web3 login error: %s, address is %s", err, request.Address)
	}
	if rtCode != common.SUCCESS {
		log.Info("[Auth Controller][Login] Web3 login failed, address is %s, return code is %d",
			request.Address, rtCode)
		appG.Response(http.StatusOK, rtCode, nil)
		return
	}
	log.Debug("[Auth Controller][Login] Web3 login successfully, address is %s", request.Address)
	appG.SuccessResponse(gin.H{
		"token": token,
	})
}
