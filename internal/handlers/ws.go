package handlers

import (
	"errors"
	"gochat/internal/hub"
	"gochat/internal/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func extractUserId(c *gin.Context) (uuid.UUID, error) {
	fromContext, ok := c.Get("userId")
	if !ok {
		return uuid.Nil, errors.New("")
	}

	asString, ok := fromContext.(string)
	if !ok {
		return uuid.Nil, errors.New("")
	}

	userId, err := uuid.Parse(asString)
	if err != nil {
		return uuid.Nil, err
	}

	return userId, nil
}

func (h *Handler) WebSocket(c *gin.Context) {
	userId, err := extractUserId(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "internal error"})
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := hub.NewClient(conn, userId)
	h.hub.Register <- client

	go client.ReadLoop(h.hub)
	go client.WriteLoop()
}
