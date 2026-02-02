package hub

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	conn *websocket.Conn
	send chan []byte
}

func (c *Client) ReadLoop(h *Hub) {
	defer func() {
		h.unregister <- c
		c.conn.Close()
	}()

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			break
		}

		h.broadcast <- message
	}
}

func (c *Client) WriteLoop() {
	defer c.conn.Close()

	for message := range c.send {
		if err := c.conn.WriteMessage(websocket.TextMessage, message); err != nil {
			break
		}
	}
}

type Hub struct {
	clients    map[*Client]bool // poor man's Set, we ensure uniqueness of clients
	Register   chan *Client
	unregister chan *Client
	broadcast  chan []byte
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.clients[client] = true
		case client := <-h.unregister:
			delete(h.clients, client)
			close(client.send)
		case message := <-h.broadcast:
			for client := range h.clients {
				client.send <- message
			}
		}
	}
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		Register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan []byte),
	}
}

func NewClient(c *websocket.Conn) *Client {
	return &Client{
		conn: c,
		send: make(chan []byte, 5),
	}
}
