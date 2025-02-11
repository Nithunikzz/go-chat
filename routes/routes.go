package routes

import (
	"github.com/go-chat/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/join", handlers.JoinChat)
	r.GET("/send", handlers.SendMessage)
	r.GET("/leave", handlers.LeaveChat)
	r.GET("/messages", handlers.GetMessages)
	r.GET("/ws", handlers.WebSocketHandler)
	return r
}
