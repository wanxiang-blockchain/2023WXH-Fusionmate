package controller

import (
	"encoding/json"
	"github.com/FusionMate/fm-backend/common"
	"github.com/FusionMate/fm-backend/common/log"
	"github.com/FusionMate/fm-backend/service"
	"github.com/gin-gonic/gin"
	"io"
	"math/big"
	"net/http"
	"strconv"
)

var testSignerTypeContractCreation int = 1
var testSignerTypeHarvest int = 2
var testPrivateKey string = "92f0952a43c92e1c69d942065f3568416854c4bd898cbb42cf23fef9c4ebc5d0"

type contractCreationSignDataRequest struct {
	Name           string `json:"name"`
	Symbol         string `json:"symbol"`
	BaseUri        string `json:"baseUri"`
	CollectionId   int64  `json:"collectionId"`
	MakerAddress   string `json:"makerAddress"`
	FactoryAddress string `json:"factoryAddress"`
	MaxSupply      uint   `json:"maxSupply"`
	MintPrice      string `json:"mintPrice"`
}

type harvestSignDataRequest struct {
	CollectionId   int64  `json:"collectionId"`
	Harvested      string `json:"harvested"`
	TokenId        uint   `json:"tokenId"`
	FactoryAddress string `json:"factoryAddress"`
}

type SignerTest struct {
}

func (s *SignerTest) Sign(c *gin.Context) {
	var err error
	var signType int
	var bodyBytes []byte
	appG := common.Gin{C: c}
	t, ok := c.Params.Get("signType")
	if ok {
		signType, err = strconv.Atoi(t)
	}
	if !ok || err != nil || (signType != testSignerTypeHarvest && signType != testSignerTypeContractCreation) {
		log.Info("[Signer Test Controller][Sign] sign type error")
		appG.Response(http.StatusOK, common.INVALID_PARAM, nil)
		return
	}
	bodyBytes, err = io.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("[Signer Test Controller][Sign] Read request body error: %s", err)
		appG.Response(http.StatusInternalServerError, common.INTERNAL_ERROR, nil)
		return
	}
	if signType == testSignerTypeContractCreation {
		var request contractCreationSignDataRequest
		err = json.Unmarshal(bodyBytes, &request)
		if err != nil {
			log.Info("[Signer Test Controller][Sign] Json unmarshal contractCreationSignDataRequest error: %s", err)
			appG.Response(http.StatusOK, common.INVALID_PARAM, nil)
			return
		}
		sig, err := service.NFTContractCreateSign(
			request.Name,
			request.Symbol,
			request.BaseUri,
			request.MakerAddress,
			testPrivateKey,
			request.CollectionId,
			request.MaxSupply,
			request.MintPrice,
			request.FactoryAddress,
		)
		if err != nil {
			log.Info("[Signer Test Controller][Sign] Sign contractCreationSignDataRequest error: %s", err)
			appG.Response(http.StatusOK, common.INVALID_PARAM, nil)
			return
		}
		appG.SuccessResponse(gin.H{
			"sig": sig,
		})
		return
	} else {
		var request harvestSignDataRequest
		err = json.Unmarshal(bodyBytes, &request)
		if err != nil {
			log.Info("[Signer Test Controller][Sign] Json unmarshal harvestSignDataRequest error: %s", err)
			appG.Response(http.StatusOK, common.INVALID_PARAM, nil)
			return
		}
		harvest := new(big.Int)
		_, ok := harvest.SetString(request.Harvested, 10)
		if !ok {
			log.Info("[Signer Test Controller][Sign] parse harvest fail: %s", err)
			appG.Response(http.StatusOK, common.INVALID_PARAM, nil)
			return
		}

		sig, err := service.HarvestSign(
			request.CollectionId,
			harvest,
			request.TokenId,
			request.FactoryAddress,
			testPrivateKey)
		if err != nil {
			log.Info("[Signer Test Controller][Sign] Sign harvestSignDataRequest error: %s", err)
			appG.Response(http.StatusOK, common.INVALID_PARAM, nil)
			return
		}
		appG.SuccessResponse(gin.H{
			"sig": sig,
		})
		return
	}
}
