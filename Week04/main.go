package main

import (
	"Week04/server"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	// TODO  wire 生成代码报错，还没理解透
	//_, err := di.InitApp()
	//if err != nil {
	//	panic(err)
	//}

	// server初始化
	server := &http.Server{
		Addr:    ":8080",
		Handler: server.NewRouter(),
	}

	// 信号监听与退出，使用gin 文档提供的方法
	// https://github.com/gin-gonic/gin#graceful-shutdown-or-restart
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")

}

