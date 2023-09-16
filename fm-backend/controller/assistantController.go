package controller

import (
	"encoding/json"
	"fmt"
	"github.com/FusionMate/fm-backend/common"
	"github.com/FusionMate/fm-backend/common/log"
	"github.com/FusionMate/fm-backend/conf"
	"github.com/FusionMate/fm-backend/dao"
	"github.com/FusionMate/fm-backend/fmGrpc"
	"github.com/FusionMate/fm-backend/model"
	"github.com/FusionMate/fm-backend/service"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	NFTContractFetchChLen        = 1000 // 队列获取NFT contract address
	NFTContractFetchTimeInterval = 3    // second
)

var (
	nftContractCh chan int64
	ticker        *time.Ticker
)

func init() {
	nftContractCh = make(chan int64, NFTContractFetchChLen)
	ticker = time.NewTicker(NFTContractFetchTimeInterval * time.Second)
}

type Assistant struct {
}

type genImgRequest struct {
	Prompt string `json:"prompt,required"`
}

// 创建AI助手的表单数据
type createRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Derive      int64  `json:"derive"`
	Replication int64  `json:"replication"`
	MintPrice   string `json:"mintPrice"`
	Prompt      string `json:"prompt"`
	ImgURI      string `json:"imgURI"`
	Typ         string `json:"type"`
}

// 创建AI助手时返回前端的
type createResponse struct {
	CollectionID int64  `json:"collectionID"`
	Name         string `json:"name"`
	Symbol       string `json:"symbol"`
	BaseURI      string `json:"baseURI"`
	MaxSupply    int64  `json:"maxSupply"`
	MintPrice    string `json:"mintPrice"`
	Signature    string `json:"signature"`
}

type notifyRequest struct {
	CollectionID    int64  `json:"collectionID"`
	ContractAddress string `json:"contractAddress"`
	Result          bool   `json:"result"`
}

/*创建类*/

// GenImgURI 创建AI助手时，根据prompt生成图片URI
// POST
// @Headers
//   - token: string M 调用者token
//
// @input
//   - prompt: string M
//
// @output
//   - imgURI: string M, e.g. https://static.fusionmate.xyz/img/:collectionId/
func (*Assistant) GenImgURI(c *gin.Context) {
	appG := common.Gin{C: c}

	var bodyBytes []byte // prompts
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("[Assistant Controller][GenImgURI] Read request body error: %s", err)
		appG.Response(http.StatusInternalServerError, common.INTERNAL_ERROR, err.Error())
		return
	}
	var request genImgRequest
	err = json.Unmarshal(bodyBytes, &request)
	if err != nil {
		log.Info("[Assistant Controller][GenImgURI] Json unmarshal error: %s", err)
		appG.Response(http.StatusInternalServerError, common.INVALID_PARAM, err.Error())
		return
	}

	lastCollectionID, err := dao.GetCollectionLastID()
	if err != nil {
		lastCollectionID = 0
	}
	curID := lastCollectionID + 1
	imgURLPrefix := conf.GConfig.GetString("matadata.imgURL")
	imgUri := fmt.Sprintf("%s/%d/img_%d.png", imgURLPrefix, curID, curID)
	appG.SuccessResponse(gin.H{
		"imgURI": imgUri,
	})
}

// Create AI Assistant 为每个新create的NFT合约准备BaseURI和metadata数据
// POST
// @Headers
//   - token: string M 调用者token
//
// @input
//   - createRequest
//
// @output
//   - createResponse
func (*Assistant) Create(c *gin.Context) {
	appG := common.Gin{C: c}
	var bodyBytes []byte
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("[Assistant Controller][Create] Read request body error: %s", err)
		appG.Response(http.StatusBadRequest, common.INVALID_PARAM, err.Error())
		return
	}
	var request createRequest
	err = json.Unmarshal(bodyBytes, &request)
	if err != nil {
		log.Info("[Assistant Controller][Create] Json unmarshal error: %s", err)
		appG.Response(http.StatusBadRequest, common.INTERNAL_ERROR, err.Error())
		return
	}

	// 存储Collection info
	symbol := genSymbol(request.Name)

	maker, ok := c.Get("address")
	if !ok {
		log.Info("[Assistant Controller][Create] get address fail")
		appG.Response(http.StatusNonAuthoritativeInfo, common.CANNOT_ACCESS_TO_RESOURCES, err.Error())
		return
	}
	collection := model.Collection{
		Name:         request.Name,
		Symbol:       symbol,
		BaseURI:      "",
		MaxSupply:    request.Replication,
		MintPrice:    request.MintPrice,
		Description:  request.Description,
		Derive:       request.Derive,
		Prompts:      request.Prompt,
		ContractAddr: "", // 后续查询
		Maker:        maker.(string),
		ImgURI:       request.ImgURI,
		Typ:          request.Typ,
	}
	err = dao.InsertCollection(&collection)
	if err != nil {
		log.Info("[Assistant Controller][Create] insert collection error: %s", err)
		appG.Response(http.StatusInternalServerError, common.INTERNAL_ERROR, err.Error())
		return
	}

	// 更新baseURI
	baseURI := genBaseURI(collection.CollectionID)
	err = dao.UptCollectionBaseURI(&model.Collection{
		CollectionID: collection.CollectionID,
		BaseURI:      baseURI,
	})
	if err != nil {
		log.Info("[Assistant Controller][Create] update baseURI fail: %s", err)
		appG.Response(http.StatusInternalServerError, common.INTERNAL_ERROR, err.Error())
		return
	}

	// 获取contractAddress
	nftContractCh <- collection.CollectionID

	// 返回数据
	privateKey := os.Getenv("PRIVATE_KEY")
	factoryContract := conf.GConfig.GetString("contract.assistantFactoryContractAddress")
	signature, err := service.NFTContractCreateSign(collection.Name, collection.Symbol, baseURI,
		collection.Maker, privateKey, collection.CollectionID, uint(collection.MaxSupply), request.MintPrice, factoryContract)
	if err != nil {
		log.Info("[Assistant Controller][Create] generate signature fail: %s", err)
		appG.Response(http.StatusInternalServerError, common.INTERNAL_ERROR, err.Error())
		return
	}
	var resp = createResponse{
		CollectionID: collection.CollectionID,
		Name:         collection.Name,
		Symbol:       collection.Symbol,
		BaseURI:      baseURI,
		MaxSupply:    collection.MaxSupply,
		MintPrice:    collection.MintPrice,
		Signature:    signature,
	}

	err = fmGrpc.PoeBotCreation(collection.CollectionID, collection.Maker, collection.Prompts)
	if err != nil && err != common.ErrDuplicatedBotCreation {
		log.Error("[Assistant Controller][Create] bot created fail: %s", err)
		appG.Response(http.StatusInternalServerError, common.INTERNAL_ERROR, err.Error())
		return
	} else if err == common.ErrDuplicatedBotCreation {
		log.Warn("[Assistant Controller][Create] bot duplicated, ignoring creation...")
	}
	appG.SuccessResponse(resp)
}

// NotifyCreateResult NFT合约部署反馈接口
// POST
// @Headers
//   - token: string M 调用者token
//
// @input
//   - collectionID
//   - contractAddr: NFT合约地址，M
//   - result: 部署结果， bool， M
//
// @output none
func (*Assistant) NotifyCreateResult(c *gin.Context) {
	appG := common.Gin{C: c}
	var bodyBytes []byte
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("[Assistant Controller][NotifyCreateResult] Read request body error: %s", err)
		appG.Response(http.StatusInternalServerError, common.INVALID_PARAM, err.Error())
		return
	}
	var request notifyRequest
	err = json.Unmarshal(bodyBytes, &request)
	if err != nil {
		log.Info("[Assistant Controller][NotifyCreateResult] Json unmarshal error: %s", err)
		appG.Response(http.StatusInternalServerError, common.INTERNAL_ERROR, err.Error())
		return
	}

	// 判断结果
	if !request.Result {
		appG.SuccessResponse("nothing update")
		return
	}

	err = dao.UptCollectionAddress(&model.Collection{
		CollectionID: request.CollectionID,
		ContractAddr: request.ContractAddress,
	})
	if err != nil {
		log.Info("[Assistant Controller][NotifyCreateResult] UptCollectionAddress fail: %s", err)
		appG.Response(http.StatusNotModified, common.INTERNAL_ERROR, err.Error())
		return
	}
	appG.SuccessResponse(nil)
}

/*查询类*/

// QueryCollectionsByPage 分页获取符合条件的AI助手
// GET
// @Query Params：
//   - type：AI助手类型，非必填
//   - makerAddr: AI助手的Maker地址，非必填
//   - page：分页的页号，从0开始，必填
//   - perPage: 每页数据量，必填
//
// @input none
// @output
//   - [] model.Collection
func (*Assistant) QueryCollectionsByPage(c *gin.Context) {
	appG := common.Gin{C: c}
	typ := c.Query("type")
	pageIdxStr := c.Query("page")
	perPageStr := c.Query("perPage")
	addr := c.Query("makerAddr")

	pageIdx, err := strconv.Atoi(pageIdxStr)
	if err != nil {
		log.Error("[Assistant Controller][QueryAssistantsByPage] get parameter of page [%s] fail:%s", pageIdxStr, err.Error())
		appG.Response(http.StatusBadRequest, common.INVALID_PARAM, err.Error())
		return
	}
	perPage, err := strconv.Atoi(perPageStr)
	if err != nil {
		log.Error("[Assistant Controller][QueryAssistantsByPage] get parameter of perPage [%s] fail:%s", perPageStr, err.Error())
		appG.Response(http.StatusBadRequest, common.INVALID_PARAM, err.Error())
		return
	}
	log.Info("[Assistant Controller][QueryByPage] typ: %s, pageIdx:%d, perPage:%d, addr:%s", typ, pageIdx, perPage, addr)

	collections, err := dao.GetCollectionsByPage(typ, addr, pageIdx, perPage)
	if err != nil {
		log.Error("[Assistant Controller][QueryAssistantsByPage] query fail:%s", err.Error())
		appG.Response(http.StatusNotFound, common.CANNOT_ACCESS_TO_RESOURCES, err.Error())
		return
	}
	appG.SuccessResponse(collections)
}

// GetCollectionById 根据合约地址检索AI助手详情
// GET
// @input none
// @output
//   - model.Collection
func (*Assistant) GetCollectionById(c *gin.Context) {
	appG := common.Gin{C: c}
	collectionIDStr := c.Param("collectionID")
	collectionID, err := strconv.Atoi(collectionIDStr)
	if err != nil {
		log.Error("[Assistant Controller][GetCollectionById] get parameter of collectionID [%s] fail:%s", collectionIDStr, err.Error())
		appG.Response(http.StatusBadRequest, common.INVALID_PARAM, err.Error())
		return
	}

	collection, err := dao.GetConnectionByID(int64(collectionID))
	if err != nil {
		log.Error("[Assistant Controller][GetCollectionById] GetConnectionByID collectionID [%s] fail:%s", collectionIDStr, err.Error())
		appG.Response(http.StatusNotFound, common.CANNOT_ACCESS_TO_RESOURCES, err.Error())
		return
	}

	// upt contract address
	if len(collection.ContractAddr) == 0 {
		if err = service.GetAndUptNFTContractAddress(int64(collectionID)); err != nil {
			log.Error("[Assistant Controller][GetCollectionById] GetAndUptNFTContractAddress collectionID [%s] fail:%s", collectionIDStr, err.Error())
			appG.Response(http.StatusNotFound, common.CANNOT_ACCESS_TO_RESOURCES, err.Error())
			return
		}
	}

	appG.SuccessResponse(collection)
}

// 拼接base uri
func genBaseURI(collectionID int64) string {
	return fmt.Sprintf("%s/%d", conf.GConfig.GetString("matadata.baseURI"), collectionID)
}

func genSymbol(name string) string {
	if len(name) >= 3 {
		return strings.ToUpper(name[:3])
	}
	return strings.ToUpper(name)
}

// FetchAndUptContractAddress get and upt contract address periodically
func FetchAndUptContractAddress(existCh chan bool) {
	for {
		select {
		case collectionID := <-nftContractCh:
			if err := service.GetAndUptNFTContractAddress(collectionID); err != nil {
				log.Error("[FetchAndUptContractAddress]fail to GetAndUptNFTContractAddress err:%s", err)
				nftContractCh <- collectionID
			}
		case <-ticker.C:
			continue
		case <-existCh:
			log.Info("stop FetchAndUptContractAddress")
			return
		}
	}
}
