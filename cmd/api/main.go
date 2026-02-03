package main

import (
	"fmt"
	"gochat/internal/config"
	"gochat/internal/handlers"
	"gochat/internal/hub"
	"gochat/internal/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config := config.Load()
	hub := hub.NewHub()
	go hub.Run()

	router := gin.Default()
	protected := router.Group("/")
	protected.Use(middleware.Auth(config.JwtKey))

	handlers := handlers.New(hub, config.JwtKey)
	router.GET("/health", handlers.Health)
	router.POST("/dev/token", handlers.DevToken)
	router.StaticFile("/", "./web/index.html")

	protected.GET("/ws", handlers.WebSocket)

	if config.Port != "" {
		log.Fatal(router.Run(fmt.Sprintf(":%v", config.Port)))
	} else {
		log.Fatal(router.Run())
	}
}
