package routes

import (
	"net/http"
	"strconv"

	"github.com/douglasmg7/gin_rest_api.git/models"
	"github.com/gin-gonic/gin"
)

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
	event, err := models.GetEventById(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}
	c.JSON(http.StatusOK, event)

}

func createEvent(c *gin.Context) {
	// token := c.Request.Header.Get("Authorization")
	// if token == "" {
	// c.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
	// return
	// }

	// userId, err := utils.VerifyToken(token)
	// if err != nil {
	// c.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
	// return
	// }

	var event models.Event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	event.UserID = c.GetInt64("userId")
	err = event.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save the event, try again latter."})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "event created", "event": event})
}

func updateEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}
	_, err = models.GetEventById(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}

	var eventUpdate models.Event
	err = c.ShouldBindJSON(&eventUpdate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	eventUpdate.ID = eventId
	err = eventUpdate.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "event updated"})
}

func deleteEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}

	err = event.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "event deleted"})
}
