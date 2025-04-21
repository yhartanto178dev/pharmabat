package main

import (
	"context"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/yhartanto178dev/pharmabot/config"
	interfaces "github.com/yhartanto178dev/pharmabot/internal/interface"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//	@title			Pharmabot API
//	@version		1.0
//	@description	API for Pharmabot
//	@termsOfService
//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:1323
//	@BasePath	/api/v1

//	@schemes					http
//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	e := echo.New()

	// MongoDB connection
	ctx, cancel := context.WithTimeout(context.Background(), cfg.ServerTimeout)
	defer cancel()
	// MongoDB connection
	client, _ := mongo.Connect(context.Background(), options.Client().ApplyURI(cfg.MongoURI))

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	db := client.Database(cfg.DatabaseName)

	// Register routes
	interfaces.RegisterRoutes(e, db)

	e.Logger.Fatal(e.Start(":1323"))
}
