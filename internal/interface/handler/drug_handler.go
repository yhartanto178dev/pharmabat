// internal/interfaces/handlers/drug_handler.go
package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yhartanto178dev/pharmabot/internal/app/drug"
)

// Make Drugs godoc
//	@Summary		Create a new drug
//	@Description	Create a new drug
//	@Tags			drugs
//	@Accept			json
//	@Produce		json
//	@Param			drug	body		drug.CreateDrugRequest	true	"Create drug"
//	@Success		201		{object}	drug.Drug
//	@Router			/drugs [post]
type DrugHandler struct {
	service *drug.Service
}

func NewDrugHandler(service *drug.Service) *DrugHandler {
	return &DrugHandler{service: service}
}

func (h *DrugHandler) CreateDrug(c echo.Context) error {
	var req drug.CreateDrugRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	newDrug, err := h.service.CreateDrug(c.Request().Context(), req.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, newDrug)
}
