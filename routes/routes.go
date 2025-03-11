package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents) 
	server.GET("/events/:id", getSingleEvent) 
	server.POST("events", createEvents)
	server.PUT("/events/:id", updateSingleEvent)
	server.DELETE("/events/:id", deleteEvent)
	server.POST("/signup", signup)
	server.POST("/login", login)
}