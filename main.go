package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/taylorjo02/go-RESTAPI/db"
	"github.com/taylorjo02/go-RESTAPI/models"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents) 
	server.POST("events", createEvents)

	err := server.Run(":8080"); if err != nil {
		log.Fatal("unable to start server, %w", err)
	} 
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events."})
		return
	}
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
	if err = event.Save(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event."})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}