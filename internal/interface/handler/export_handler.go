package handlers

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/yhartanto178dev/pharmabot/internal/app/export"
)

// Make Export godoc
//
//	@Summary		Export data to CSV
//	@Description	Export data to CSV
//	@Tags			export
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	export.ExportReport
//	@Router			/export [get]
//	@Param			drugID			body	string	false	"Drug ID"
//	@Param			endUserID		body	string	false	"End User ID"
//	@Param			expirationDate	body	string	false	"Expiration Date"
//	@Param			quantity		body	int		false	"Quantity"

type ExportHandler struct {
	service *export.Service
}

func NewExportHandler(service *export.Service) *ExportHandler {
	return &ExportHandler{service: service}
}

func (h *ExportHandler) ExportCSV(c echo.Context) error {
	report, err := h.service.GenerateCSVReport(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": fmt.Sprintf("Failed to generate report: %v", err),
		})
	}

	// Set CSV headers
	c.Response().Header().Set(echo.HeaderContentType, "text/csv")
	c.Response().Header().Set(echo.HeaderContentDisposition, "attachment; filename=export.csv")

	writer := csv.NewWriter(c.Response())
	defer writer.Flush()

	// Write headers
	if err := writer.Write(report.Headers); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to write CSV headers",
		})
	}

	// Write rows
	for _, row := range report.Rows {
		csvRow := make([]string, len(report.Headers))
		csvRow[0] = row.DrugName

		for i, header := range report.Headers[1:] {
			if i%2 == 0 { // Date column
				endUser := strings.TrimSuffix(header, " Date")
				csvRow[i+1] = row.Expirations[endUser].Date
			} else { // Quantity column
				endUser := strings.TrimSuffix(header, " Quantity")
				csvRow[i+1] = strconv.Itoa(row.Expirations[endUser].Quantity)
			}
		}

		if err := writer.Write(csvRow); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to write CSV row",
			})
		}
	}

	return nil
}
