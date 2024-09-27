package main

import (
	"net/http"
	"strconv"

	"github.com/douglasmg7/gin_rest_api.git/db"
	"github.com/douglasmg7/gin_rest_api.git/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDb()
	server := gin.Default()
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events, try again latter."})
		return
	}
	c.JSON(http.StatusOK, events)
}

func getEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}
}

func createEvent(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	event.ID = 1
	event.UserID = 1
	err = event.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save the event, try again latter."})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "event created", "event": event})
}
