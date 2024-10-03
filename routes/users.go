package routes

import (
	"net/http"

	"github.com/douglasmg7/gin_rest_api.git/models"
	"github.com/douglasmg7/gin_rest_api.git/utils"
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

func login(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if user.ValidateCredentials() != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Credentials invalid"})
		return
	}

	token, err := utils.GenerateToke(user.Email, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not authenticate user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User autenticated", "token": token})
}
