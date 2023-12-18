package route

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type RouterInitialize struct{}

func (r RouterInitialize) InitRouter() {
	engine := gin.Default()
	
	group := engine.Group("/api/cicoyi/v1")
	InitUserGroupRouter(group)
	
	// 启动http服务器
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", viper.GetString("port")),
		Handler: engine,
	}
	
	// 启动服务器（非阻塞）
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Failed to serve api server: %v", err)
		}
	}()
	
	log.Println("The userHandlers api service has been started, and the port is :", viper.GetString("port"))
	
	// 优雅关闭服务
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	
	// 阻塞等待信号
	<-quit
	
	// 设置超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	// 关闭服务器
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Shutdown(): %s", err)
	}
	log.Println("---------------------The userHandlers program has been closed---------------------------")
}
