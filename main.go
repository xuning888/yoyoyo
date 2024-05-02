// Code generated by hertz generator.

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xuning888/ollama-hertz/config"
	"github.com/xuning888/ollama-hertz/internal/controller"
	"github.com/xuning888/ollama-hertz/internal/dal"
	"github.com/xuning888/ollama-hertz/internal/service"
	"github.com/xuning888/ollama-hertz/pkg/http"
	"github.com/xuning888/ollama-hertz/pkg/logger"
	"time"
)

func init() {
	config.Init()
	dal.Init()
	logger.InitLogger()
}

func Register(router *gin.Engine) {
	chatService := service.NewChatService()
	chatController := controller.NewChatController(chatService)
	chatController.Register(router)
}

func main() {
	router := gin.Default()
	// register html
	router.Static("/", "./static")

	// register router
	Register(router)

	server := http.NewServer(
		router,
		config.DefaultConfig.ServerPort,
		http.WithShutdownTimout(time.Second*time.Duration(25)),
	)
	server.Serve()
}
