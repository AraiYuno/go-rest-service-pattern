package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"golang-rest-api-starter/services"
)

type UserController struct{}

func (h UserController) Post(c *gin.Context) {
	var input services.RegisterUserInput

	// Bind request body to input struct
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service := &services.Service{
		Driver: "postgres",
	}

	user, err := service.RegisterUser(&input)
	if err != nil {
		fmt.Println("Failed to register user:", err)
		return
	}

	fmt.Println("Registered user:", user)

	c.JSON(http.StatusOK, user)
}
