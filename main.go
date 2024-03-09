package main

import (
	"context"
	"fmt"

	"fraser-chapman/go-crud-api/preferences"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://root:example@localhost:27017"

func main() {
	testDatabaseConnection()
	initialiseRestServer()
}

func testDatabaseConnection() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Send a ping to confirm a successful connection
	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to MongoDB!")
}

func initialiseRestServer() {
	router := gin.Default()
	router.GET("/preferences", preferences.GetAllUsersPreferences)
	router.Run("localhost:8080")
}