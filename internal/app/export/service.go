package export

import (
	"context"
	"fmt"

	"github.com/yhartanto178dev/pharmabot/internal/domain/drug"
	"github.com/yhartanto178dev/pharmabot/internal/domain/enduser"
	"github.com/yhartanto178dev/pharmabot/internal/domain/expiration"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service struct {
	expRepo     expiration.Repository
	drugRepo    drug.Repository
	endUserRepo enduser.Repository
}

func NewService(
	expRepo expiration.Repository,
	drugRepo drug.Repository,
	endUserRepo enduser.Repository,
) *Service {
	return &Service{
		expRepo:     expRepo,
		drugRepo:    drugRepo,
		endUserRepo: endUserRepo,
	}
}

type ExportReport struct {
	Headers []string
	Rows    []ExportRow
}

type ExportRow struct {
	DrugName    string
	Expirations map[string]ExpirationDetail
}

type ExpirationDetail struct {
	Date     string
	Quantity int
}

func (s *Service) GenerateCSVReport(ctx context.Context) (*ExportReport, error) {
	pipeline := mongo.Pipeline{
		{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "drugs"},
			{Key: "localField", Value: "drug_id"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "drug"},
		}}},
		{{Key: "$unwind", Value: bson.D{
			{Key: "path", Value: "$drug"},
			{Key: "preserveNullAndEmptyArrays", Value: false},
		}}},
		{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "end_users"},
			{Key: "localField", Value: "end_user_id"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "end_user"},
		}}},
		{{Key: "$unwind", Value: bson.D{
			{Key: "path", Value: "$end_user"},
			{Key: "preserveNullAndEmptyArrays", Value: false},
		}}},
		{{Key: "$project", Value: bson.D{
			{Key: "drug_name", Value: "$drug.name"},
			{Key: "end_user", Value: "$end_user.name"},
			{Key: "expiration_date", Value: bson.D{
				{Key: "$dateToString", Value: bson.D{
					{Key: "format", Value: "%Y-%m-%d"},
					{Key: "date", Value: "$expiration_date"},
				}},
			}},
			{Key: "quantity", Value: "$quantity"},
		}}},
	}

	cursor, err := s.expRepo.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("failed to aggregate data: %v", err) // Tambahkan detail error
	}
	defer cursor.Close(ctx)

	var results []struct {
		DrugName       string `bson:"drug_name"`
		EndUser        string `bson:"end_user"`
		ExpirationDate string `bson:"expiration_date"`
		Quantity       int    `bson:"quantity"`
	}

	if err = cursor.All(ctx, &results); err != nil {
		return nil, fmt.Errorf("failed to decode results: %v", err) // Tambahkan detail error
	}

	report := &ExportReport{
		Headers: []string{"Drug Name"},
		Rows:    make([]ExportRow, 0),
	}

	endUsers := make(map[string]struct{})
	drugMap := make(map[string]*ExportRow)

	for _, item := range results {
		// Track unique end users
		endUsers[item.EndUser] = struct{}{}

		// Initialize drug row if not exists
		if _, exists := drugMap[item.DrugName]; !exists {
			drugMap[item.DrugName] = &ExportRow{
				DrugName:    item.DrugName,
				Expirations: make(map[string]ExpirationDetail),
			}
		}

		// Add expiration detail
		drugMap[item.DrugName].Expirations[item.EndUser] = ExpirationDetail{
			Date:     item.ExpirationDate,
			Quantity: item.Quantity,
		}
	}

	// Build headers with date and quantity columns
	for endUser := range endUsers {
		report.Headers = append(report.Headers,
			fmt.Sprintf("%s Date", endUser),
			fmt.Sprintf("%s Quantity", endUser),
		)
	}

	// Convert map to slice
	for _, row := range drugMap {
		report.Rows = append(report.Rows, *row)
	}

	return report, nil
}
