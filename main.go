package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/taylorjo02/go-RESTAPI/models"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents) 
	server.POST("events", createEvents)

	err := server.Run(":8080"); if err != nil {
		log.Fatal("unable to start server, %w", err)
	} 
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvents(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event); if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	event.ID = 1
	event.UserID = 1
	event.Save()
	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}