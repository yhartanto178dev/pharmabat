package expiration

import (
	"strings"
	"time"
)

type CreateExpirationRequest struct {
	DrugID         string `json:"drug_id" validate:"required"`
	EndUserID      string `json:"end_user_id" validate:"required"`
	ExpirationDate string `json:"expiration_date" validate:"required"`
	Quantity       int    `json:"quantity" validate:"required,min=1"`
}

func (r *CreateExpirationRequest) ParseExpirationDate() (time.Time, error) {
	// Trim whitespace and parse date
	dateStr := strings.TrimSpace(r.ExpirationDate)

	// Coba format YYYY-MM-DD
	parsedDate, err := time.Parse("2006-01-02", dateStr)
	if err == nil {
		return parsedDate, nil
	}

	// Coba format lain jika perlu
	return time.Parse(time.RFC3339, dateStr)
}

// internal/app/export/service.go
// type ExportRow struct {
// 	DrugName    string
// 	Expirations map[string]ExpirationDetail
// }

// type ExpirationDetail struct {
// 	Date     string
// 	Quantity int
// }
