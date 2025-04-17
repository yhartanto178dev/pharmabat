// internal/infrastructure/mongodb/expiration_repository.go
package mongodb

import (
	"context"

	"github.com/yhartanto178dev/pharmabot/internal/domain/expiration"
	"go.mongodb.org/mongo-driver/mongo"
)

type ExpirationRepository struct {
	collection *mongo.Collection
}

func NewExpirationRepository(db *mongo.Database) *ExpirationRepository {
	return &ExpirationRepository{
		collection: db.Collection("expirations"),
	}
}

func (r *ExpirationRepository) Create(ctx context.Context, exp *expiration.Expiration) error {
	_, err := r.collection.InsertOne(ctx, exp)
	return err
}

func (r *ExpirationRepository) Aggregate(ctx context.Context, pipeline mongo.Pipeline) (*mongo.Cursor, error) {
	return r.collection.Aggregate(ctx, pipeline)
}
