package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type preferences struct {
	UserId	string	`json:"userId"`
	Homepage string `json:"homepage"`
	Currency string `json:"currency"`
}

var usersPreferences = []preferences{
	{UserId: "1", Homepage: "https://www.tissueinc.co.uk/home", Currency: "GBP"},
	{UserId: "2", Homepage: "https://www.tissueinc.co.uk/projects", Currency: "USD"},
	{UserId: "3", Homepage: "https://www.tissueinc.co.uk/home", Currency: "USD"},
}

func getAllUsersPreferences(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, usersPreferences)
}

func main() {
	router := gin.Default()
	router.GET("/preferences", getAllUsersPreferences)

	router.Run("localhost:8080")
}