package hub

import (
	"encoding/json"
	"gochat/internal/models"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	conn   *websocket.Conn
	userId uuid.UUID
	send   chan *models.Message
}

func (c *Client) ReadLoop(h *Hub) {
	defer func() {
		h.unregister <- c
		c.conn.Close()
	}()

	for {
		_, payload, err := c.conn.ReadMessage()
		if err != nil {
			break
		}

		message := new(models.Message)
		if err := json.Unmarshal(payload, message); err != nil {
			log.Println("invalid message format:", err)
			continue
		}

		message.From = c.userId
		message.Timestamp = time.Now()

		h.direct <- message
	}
}

func (c *Client) WriteLoop() {
	defer c.conn.Close()

	for message := range c.send {
		data, err := json.Marshal(message)
		if err != nil {
			// handle error
			break
		}
		c.conn.WriteMessage(websocket.TextMessage, data)
	}
}

type Hub struct {
	clients    map[uuid.UUID]*Client
	Register   chan *Client
	unregister chan *Client
	direct     chan *models.Message
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.clients[client.userId] = client
		case client := <-h.unregister:
			delete(h.clients, client.userId)
			close(client.send)
		case message := <-h.direct:
			if client, ok := h.clients[message.To]; ok {
				client.send <- message
			} else {
				log.Println("user offline")
			}
		}
	}
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[uuid.UUID]*Client),
		Register:   make(chan *Client),
		unregister: make(chan *Client),
		direct:     make(chan *models.Message),
	}
}

func NewClient(c *websocket.Conn, u uuid.UUID) *Client {
	return &Client{
		conn:   c,
		userId: u,
		send:   make(chan *models.Message, 5),
	}
}
