package routes

import (
	"net/http"

	"github.com/douglasmg7/gin_rest_api.git/models"
	"github.com/gin-gonic/gin"
)

func signup(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = user.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user, try again latter.", "error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}
