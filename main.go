package main

import (
	"context"
	"ecloudsystem/configs"
	"ecloudsystem/middleware"
	"ecloudsystem/pkg/cache"
	"ecloudsystem/pkg/db"
	"ecloudsystem/routers"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"time"
)


func main() {

	/**
	 * 初始化日志
	 */
	configs.InitLogs()

	defer func() {
		_ = configs.Logger.Sync()
	}()

	gin.DisableConsoleColor()
	gin.SetMode(gin.DebugMode)

	var r *gin.Engine

	if configs.Get().App.Debug == true {
		r = gin.Default()
	} else {
		r = gin.New()
		r.Use(middleware.GinLoggerZap(configs.Logger,"10", true),
			middleware.GinRecoveryWithZap(configs.Logger, true))
	}

	r = routers.InitRouter(r)
	db.InitDb()
	cache.InitCache()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", configs.Get().App.Port),
		Handler:        r,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			configs.Logger.Error(fmt.Sprintf("Listen: %s\n", err))
		}
	}()

	defer func() {
		db.Client.DbWClose()
	}()

	defer func() {
		cache.Client.RedisClose()
	}()

	// 优雅的退出
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<- quit
	configs.Logger.Error("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		configs.Logger.Error(fmt.Sprintf("erver Shutdown: %s", err))
	}

	//r.Run(":8000")
}
