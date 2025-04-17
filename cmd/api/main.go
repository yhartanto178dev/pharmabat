package main

import (
	"context"

	"github.com/labstack/echo/v4"
	interfaces "github.com/yhartanto178dev/pharmabot/internal/interface"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	e := echo.New()

	// MongoDB connection
	client, _ := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://192.168.75.1:27017"))
	db := client.Database("pharmacy")

	// Register routes
	interfaces.RegisterRoutes(e, db)

	e.Logger.Fatal(e.Start(":1323"))
}
