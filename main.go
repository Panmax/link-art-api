package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"link-art-api/application"
	"link-art-api/application/middleware"
	"link-art-api/domain/model"
	"link-art-api/infrastructure/config"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func init() {
	config.Setup()
	model.Setup()
	middleware.SetupAuth()
}

func main() {
	gin.SetMode(config.ServerConfig.Mode)
	engine := gin.Default()
	application.SetupRoute(engine)
	application.SetupSchedule()

	Port := fmt.Sprintf(":%d", config.ServerConfig.Port)
	server := &http.Server{
		Addr:         Port,
		Handler:      engine,
		ReadTimeout:  config.ServerConfig.ReadTimeout,
		WriteTimeout: config.ServerConfig.WriteTimeout,
	}

	fmt.Println("|-----------------------------------|")
	fmt.Println("|            go-gin-api             |")
	fmt.Println("|-----------------------------------|")
	fmt.Println("|  Go Http Server Start Successful  |")
	fmt.Println("|    Port" + Port + "     Pid:" + fmt.Sprintf("%d", os.Getpid()) + "        |")
	fmt.Println("|-----------------------------------|")
	fmt.Println("")

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, os.Interrupt)
	sig := <-signalChan
	log.Println("Get Signal:", sig)
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
