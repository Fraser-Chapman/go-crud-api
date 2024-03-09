package main

import (
	"fraser-chapman/go-crud-api/preferences"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/preferences", preferences.GetAllUsersPreferences)

	router.Run("localhost:8080")
}