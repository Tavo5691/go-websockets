package main

import (
	"fmt"
	"gochat/internal/config"
	"gochat/internal/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config := config.Load()
	router := gin.Default()
	handlers := handlers.New()

	router.GET("/health", handlers.Health)
	router.GET("/ws", handlers.WebSocket)
	router.StaticFile("/", "./web/index.html")

	if config.Port != "" {
		log.Fatal(router.Run(fmt.Sprintf(":%v", config.Port)))
	} else {
		log.Fatal(router.Run())
	}
}
