package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yhartanto178dev/pharmabot/internal/app/enduser"
)

// Make End User godoc
//
//	@Summary		Create a new End User
//	@Description	Create a new End User
//	@Tags			end-users
//	@Accept			json
//	@Produce		json
//	@Param			end-user	body		enduser.CreateEndUserRequest	true	"Create End User"
//	@Success		201			{object}	enduser.EndUser
//	@Router			/end-users [post]
type EndUserHandler struct {
	service *enduser.Service
}

func NewEndUserHandler(service *enduser.Service) *EndUserHandler {
	return &EndUserHandler{service: service}
}

func (h *EndUserHandler) CreateEndUser(c echo.Context) error {
	var req enduser.CreateEndUserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	newEndUser, err := h.service.CreateEndUser(c.Request().Context(), req.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, newEndUser)
}
