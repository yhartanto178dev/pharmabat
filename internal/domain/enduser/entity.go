package enduser

import "go.mongodb.org/mongo-driver/bson/primitive"

type EndUser struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
}

func NewEndUser(name string) *EndUser {
	return &EndUser{
		ID:   primitive.NewObjectID(),
		Name: name,
	}
}
