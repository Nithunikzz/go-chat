package main

import (
	"log"

	"github.com/go-chat/config"
	"github.com/go-chat/database"
	"github.com/go-chat/routes"
)

func main() {
	cfg := config.LoadConfig()

	database.Connect(cfg.DatabaseURL)

	r := routes.SetupRouter()
	log.Println("Server running on port:", cfg.ServerPort)
	r.Run(":" + cfg.ServerPort)
}
