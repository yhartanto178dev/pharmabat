// internal/infrastructure/mongodb/drug_repository.go
package mongodb

import (
	"context"

	"github.com/yhartanto178dev/pharmabot/internal/domain/drug"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DrugRepository struct {
	collection *mongo.Collection
}

// FindAll implements drug.Repository.
func (r *DrugRepository) FindAll(ctx context.Context) ([]*drug.Drug, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var drugs []*drug.Drug
	for cursor.Next(ctx) {
		var d drug.Drug
		if err := cursor.Decode(&d); err != nil {
			return nil, err
		}
		drugs = append(drugs, &d)
	}

	return drugs, nil
}
func NewDrugRepository(db *mongo.Database) *DrugRepository {
	return &DrugRepository{
		collection: db.Collection("drugs"),
	}
}

func (r *DrugRepository) Create(ctx context.Context, d *drug.Drug) error {
	_, err := r.collection.InsertOne(ctx, d)
	return err
}

// func (r *DrugRepository) FindAll(ctx context.Context) ([]*drug.Drug, error) {
// 	// Implementation
// }
