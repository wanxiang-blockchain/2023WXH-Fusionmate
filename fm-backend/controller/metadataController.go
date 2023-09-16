package controller

import (
	"github.com/FusionMate/fm-backend/common"
	"github.com/FusionMate/fm-backend/common/log"
	"github.com/FusionMate/fm-backend/dao"
	"github.com/FusionMate/fm-backend/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Metadata struct {
}

// GetTokenMetadata 通过NFT合约的BaseURI，可以查看到合约的metadata
// GET
// @output
// - model.TokenMetaData
func (m *Metadata) GetTokenMetadata(c *gin.Context) {
	appG := common.Gin{C: c}
	collectionIDStr := c.Param("collectionID")
	collectionID, err := strconv.Atoi(collectionIDStr)
	if err != nil {
		log.Error("[Metadata Controller][GetTokenMetadata] get parameter of collectionID [%s] fail:%s", collectionIDStr, err.Error())
		appG.Response(http.StatusBadRequest, common.INVALID_PARAM, err.Error())
		return
	}

	tokenIDStr := c.Param("tokenID")
	tokenID, err := strconv.Atoi(tokenIDStr)
	if err != nil {
		log.Error("[Metadata Controller][GetTokenMetadata] get parameter of tokenID [%s] fail:%s", tokenIDStr, err.Error())
		appG.Response(http.StatusBadRequest, common.INVALID_PARAM, err.Error())
		return
	}

	collection, err := dao.GetConnectionByID(int64(collectionID))
	if err != nil {
		log.Error("[Metadata Controller][GetTokenMetadata] query collectionID:%d fail:%s",
			collectionID, err.Error())
		appG.Response(http.StatusNotFound, common.CANNOT_ACCESS_TO_RESOURCES, err.Error())
		return
	}

	metaData := model.TokenMetaData{
		CollectionID: collectionID,
		TokenID:      tokenID,
		Description:  collection.Description,
		Image:        collection.ImgURI,
		Name:         collection.Name,
		Attributes: []model.Attribute{
			{TraitType: "mintPrice", Value: collection.MintPrice},
			{TraitType: "type", Value: collection.Typ},
		},
	}

	appG.SuccessResponse(metaData)
}
