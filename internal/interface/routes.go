// internal/interfaces/routes.go
package interfaces

import (
	"github.com/labstack/echo/v4"
	"github.com/yhartanto178dev/pharmabot/internal/app/drug"
	"github.com/yhartanto178dev/pharmabot/internal/app/enduser"
	"github.com/yhartanto178dev/pharmabot/internal/app/expiration"
	"github.com/yhartanto178dev/pharmabot/internal/app/export"
	"github.com/yhartanto178dev/pharmabot/internal/infrastructure/mongodb"
	handlers "github.com/yhartanto178dev/pharmabot/internal/interface/handler"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterRoutes(e *echo.Echo, db *mongo.Database) {
	// Initialize repositories
	drugRepo := mongodb.NewDrugRepository(db)
	endUserRepo := mongodb.NewEndUserRepository(db)
	expRepo := mongodb.NewExpirationRepository(db)

	// Initialize services
	drugService := drug.NewService(drugRepo)
	endUserService := enduser.NewService(endUserRepo)
	expService := expiration.NewService(expRepo, drugRepo, endUserRepo)

	// Initialize handlers
	drugHandler := handlers.NewDrugHandler(drugService)
	endUserHandler := handlers.NewEndUserHandler(endUserService)
	expHandler := handlers.NewExpirationHandler(expService)
	// Initialize services
	exportService := export.NewService(expRepo, drugRepo, endUserRepo)

	// Initialize handlers
	exportHandler := handlers.NewExportHandler(exportService)

	// Register routes
	ApiV1 := e.Group("/api/v1")
	// Register routes
	ApiV1.POST("/drugs", drugHandler.CreateDrug)
	ApiV1.POST("/end-users", endUserHandler.CreateEndUser)
	ApiV1.POST("/expirations", expHandler.CreateExpiration)
	ApiV1.GET("/export", exportHandler.ExportCSV)

}
