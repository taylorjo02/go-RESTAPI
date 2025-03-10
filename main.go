package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/taylorjo02/go-RESTAPI/db"
	"github.com/taylorjo02/go-RESTAPI/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	err := server.Run(":8080")
	if err != nil {
		log.Fatal("unable to start server, %w", err)
	}
}
