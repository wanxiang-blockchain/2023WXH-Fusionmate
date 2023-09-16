package router

import (
	"github.com/FusionMate/fm-backend/controller"
	"github.com/FusionMate/fm-backend/middleware"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	pprof.Register(r)
	r.Use(middleware.Cors)
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/api/v1", hailing)
	authGroup := r.Group("/api/v1/auth")
	var auth controller.Auth
	{
		authGroup.POST("/login", auth.Login)
	}

	assistantGroup := r.Group("/api/v1/assistant")
	var assistant controller.Assistant
	{
		assistantGroup.POST("/genImgURI", middleware.TokenCheck, assistant.GenImgURI)
		assistantGroup.POST("/create", middleware.TokenCheck, assistant.Create)
		assistantGroup.POST("/notifyCreateResult", middleware.TokenCheck, assistant.NotifyCreateResult)

		assistantGroup.GET("/collections", assistant.QueryCollectionsByPage)
		assistantGroup.GET("/collection/:collectionID", assistant.GetCollectionById)
	}

	metadataGroup := r.Group("/api/v1/metadata")
	var metadata controller.Metadata
	{
		metadataGroup.GET("/:collectionID/:tokenID", metadata.GetTokenMetadata)
	}
	harvestGroup := r.Group("/api/v1/harvest")
	var harvest controller.Harvest
	{
		harvestGroup.GET("/:collectionID/:tokenID", middleware.TokenCheck, harvest.GetHarvestTokenNum)
	}
	return r
}

func hailing(c *gin.Context) {
	c.String(http.StatusOK, "It works!")
}
