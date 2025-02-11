package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func WebSocketHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	clientID := c.Query("id")
	client := room.Join(clientID)

	// Send old messages first
	for _, msg := range client.History {
		conn.WriteMessage(websocket.TextMessage, []byte(msg))
	}

	// Start listening for new messages
	go func() {
		for msg := range client.Message {
			conn.WriteMessage(websocket.TextMessage, []byte(msg))
		}
	}()

	// Handle incoming WebSocket messages
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			room.Leave(clientID)
			break
		}
		room.Send(clientID, string(message))
	}
}
