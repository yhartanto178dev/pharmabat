package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yhartanto178dev/pharmabot/internal/app/expiration"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Make Expiration godoc
//
//	@Summary		Create a new Expiration
//	@Description	Create a new Expiration
//	@Tags			expirations
//	@Accept			json
//	@Produce		json
//	@Param			expiration	body		expiration.CreateExpirationRequest	true	"Create Expiration"
//	@Success		201			{object}	expiration.Expiration
//	@Router			/expirations [post]

type ExpirationHandler struct {
	service *expiration.Service
}

func NewExpirationHandler(service *expiration.Service) *ExpirationHandler {
	return &ExpirationHandler{service: service}
}

func (h *ExpirationHandler) CreateExpiration(c echo.Context) error {
	var req expiration.CreateExpirationRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request format",
		})
	}

	// Parse expiration date
	expDate, err := req.ParseExpirationDate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": fmt.Sprintf("Invalid date format. Use YYYY-MM-DD. Error: %v", err),
		})
	}

	// Convert string IDs to ObjectID
	drugID, err := primitive.ObjectIDFromHex(req.DrugID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid drug ID format",
		})
	}

	endUserID, err := primitive.ObjectIDFromHex(req.EndUserID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid end user ID format",
		})
	}

	// Panggil service
	newExp, err := h.service.CreateExpiration(
		c.Request().Context(),
		drugID,
		endUserID,
		expDate,
		req.Quantity,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, newExp)
}
