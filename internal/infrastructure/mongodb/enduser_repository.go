// internal/infrastructure/mongodb/enduser_repository.go
package mongodb

import (
	"context"

	"github.com/yhartanto178dev/pharmabot/internal/domain/enduser"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type EndUserRepository struct {
	collection *mongo.Collection
}

func NewEndUserRepository(db *mongo.Database) *EndUserRepository {
	return &EndUserRepository{
		collection: db.Collection("end_users"),
	}
}

func (r *EndUserRepository) Create(ctx context.Context, eu *enduser.EndUser) error {
	_, err := r.collection.InsertOne(ctx, eu)
	return err
}

func (r *EndUserRepository) FindAll(ctx context.Context) ([]*enduser.EndUser, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var endUsers []*enduser.EndUser
	if err = cursor.All(ctx, &endUsers); err != nil {
		return nil, err
	}
	return endUsers, nil
}
