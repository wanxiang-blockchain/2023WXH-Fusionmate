package main

import (
	"context"
	"github.com/FusionMate/fm-backend/common/log"
	"github.com/FusionMate/fm-backend/conf"
	"github.com/FusionMate/fm-backend/controller"
	"github.com/FusionMate/fm-backend/dao"
	"github.com/FusionMate/fm-backend/fmGrpc"
	"github.com/FusionMate/fm-backend/router"
	_ "github.com/FusionMate/fm-backend/service"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	log.Info("[init] server started...")

	// connect to db
	dao.InitDB()

	// start a gin server for API.
	gin.SetMode(conf.GConfig.GetString("server.runMode"))
	endPoint := conf.GConfig.GetString("server.addr")
	handle := router.InitRouter()
	server := &http.Server{
		Addr:    endPoint,
		Handler: handle,
	}
	log.Info("[init] start http server listening %s", endPoint)
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("[init] server is not started, error is %s", err)
		}
	}()

	// fetch and update nft contract address
	existCh := make(chan bool)
	go controller.FetchAndUptContractAddress(existCh)

	// Following code will start a http service for prometheus metrics export.
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		err := http.ListenAndServe(":9111", nil)
		if err != nil {
			log.Fatal(" %v", err)
		}
	}()
	log.Info("[init] prometheus exporter listening...")

	// start a grpc server for CleanupD communicate with Freezer
	lis, err := net.Listen("tcp", conf.GConfig.GetString("server.grpc"))
	if err != nil {
		log.Fatal("failed to listen on grpc port: %v", err)
	}
	g := grpc.NewServer()
	fmGrpc.RegisterBotServiceServer(g, &fmGrpc.BotServiceServerImpl{})
	log.Info("[init] grpc server listening at %v", lis.Addr())
	go func() {
		if err = g.Serve(lis); err != nil {
			log.Fatal("failed to serve: %v", err)
		}
	}()

	// capture the terminating signal to stop gracefully.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGSTOP)

	<-quit
	existCh <- true

	// stop the http server
	log.Info("[quit] stopping the http server...")
	httpServerContext, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(httpServerContext); err != nil {
		log.Fatal("[quit] server forced to shutdown:", err)
	}
	log.Info(" http server stopped")
	log.Info("[quit] server shutting...")

	// stop the grpc server
	g.GracefulStop()
	log.Info(" grpc server stopped")

	// stop db
	if err := dao.CloseDB(); err != nil {
		log.Fatal("[quit] close db fail:", err)
	}
	log.Info("db closed")
}
