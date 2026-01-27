package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

/*func readerLoop(conn *websocket.Conn, ch chan []byte) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			conn.Close()
			break
		}

		ch <- message
	}
}*/

func (h *Handler) WebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}

		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Println(err)
			break
		}
	}

}
