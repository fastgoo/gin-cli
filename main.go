package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	cRouter "gin-cli/router"

	_ "gin-cli/plugins/env"   //这里是载入环境变量的配置信息，必须优先引入
	_ "gin-cli/plugins/redis" //这里载入redis插件

	_ "gin-cli/modules/api_v1" //这里是引入api_v1模块
)

func main() {
	srv := &http.Server{
		Addr:    os.Getenv("HTTP_HOST") + ":" + os.Getenv("HTTP_PORT"),
		Handler: cRouter.Router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
