package expiration

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Create(ctx context.Context, exp *Expiration) error
	Aggregate(ctx context.Context, pipeline mongo.Pipeline) (*mongo.Cursor, error)
}
