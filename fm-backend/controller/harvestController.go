package controller

import (
	"github.com/FusionMate/fm-backend/common"
	"github.com/FusionMate/fm-backend/common/log"
	"github.com/FusionMate/fm-backend/conf"
	"github.com/FusionMate/fm-backend/service"
	"github.com/gin-gonic/gin"
	"math/big"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

type Harvest struct {
}

// GetHarvestTokenNum 获取Harvest的token数量和签名
// GET
// @Headers
//   - token string
//
// @Query Params：
//   - tba: common.address, NFT TBA地址
//
// @input none
// @output
//   - tokenNum: erc20奖励数量
func (*Harvest) GetHarvestTokenNum(c *gin.Context) {
	appG := common.Gin{C: c}
	collectionIDStr := c.Param("collectionID")
	collectionID, err := strconv.Atoi(collectionIDStr)
	if err != nil {
		log.Error("[Harvest Controller][GetHarvestTokenNum] get parameter of collectionID [%s] fail:%s", collectionIDStr, err.Error())
		appG.Response(http.StatusBadRequest, common.INVALID_PARAM, err.Error())
		return
	}

	tokenIDStr := c.Param("tokenID")
	tokenID, err := strconv.Atoi(tokenIDStr)
	if err != nil {
		log.Error("[Harvest Controller][GetHarvestTokenNum] get parameter of tokenID [%s] fail:%s", tokenIDStr, err.Error())
		appG.Response(http.StatusBadRequest, common.INVALID_PARAM, err.Error())
		return
	}
	// todo check if user owns this token
	// todo harvest := dao.CountAndUptAIMessageHarvest(collectionID, tokenID)

	harvest := rand.Intn(100)
	h := big.NewInt(int64(harvest))
	bigMultiplier := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	h.Mul(h, bigMultiplier)
	harvestStr := h.String()
	privateKey := os.Getenv("PRIVATE_KEY")
	factoryContract := conf.GConfig.GetString("contract.assistantFactoryContractAddress")
	signature, err := service.HarvestSign(int64(collectionID), h, uint(tokenID), factoryContract, privateKey)
	if err != nil {
		log.Info("[Assistant Controller][Create] generate signature fail: %s", err)
		appG.Response(http.StatusInternalServerError, common.INTERNAL_ERROR, err.Error())
		return
	}
	appG.SuccessResponse(gin.H{
		"collectionID": collectionID,
		"tokenID":      tokenID,
		"tokenNum":     harvestStr,
		"signature":    signature,
	})
}
