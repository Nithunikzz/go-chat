package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chat/chat"

	"github.com/gin-gonic/gin"
)

var room = chat.NewChatRoom()

func JoinChat(c *gin.Context) {
	id := c.Query("id")
	client := room.Join(id)

	c.JSON(http.StatusOK, gin.H{"message": id + " joined"})
	go func() {
		for msg := range client.Message {
			// Store messages in DB (future enhancement)
			c.Writer.Write([]byte(msg + "\n"))

		}
	}()

}

func SendMessage(c *gin.Context) {
	id := c.Query("id")
	message := c.Query("message")
	room.Send(id, message)
	c.JSON(http.StatusOK, gin.H{"status": "sent"})
}

func LeaveChat(c *gin.Context) {
	id := c.Query("id")
	room.Leave(id)
	c.JSON(http.StatusOK, gin.H{"message": id + " left"})
}
func GetMessages(c *gin.Context) {
	clientID := c.Query("id")

	messages := room.GetMessages(clientID)
	if messages == nil {
		log.Println("❌ Client not found:", clientID)
		c.JSON(http.StatusNotFound, gin.H{"error": "Client not found"})
		return
	}

	if len(messages) > 0 {
		fmt.Println("✅ Messages retrieved for", clientID, ":", messages)
		c.JSON(http.StatusOK, gin.H{"messages": messages})
	} else {
		fmt.Println("⚠️ No new messages for", clientID)
		c.JSON(http.StatusNoContent, gin.H{"message": "No new messages"})
	}
}
