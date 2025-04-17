// internal/domain/drug/entity.go
package drug

import "go.mongodb.org/mongo-driver/bson/primitive"

type Drug struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
}

func NewDrug(name string) *Drug {
	return &Drug{
		ID:   primitive.NewObjectID(),
		Name: name,
	}
}
