package expiration

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Expiration struct {
	ID             primitive.ObjectID `bson:"_id"`
	DrugID         primitive.ObjectID `bson:"drug_id"`
	EndUserID      primitive.ObjectID `bson:"end_user_id"`
	ExpirationDate time.Time          `bson:"expiration_date"`
	Quantity       int                `bson:"quantity"`
}

func NewExpiration(drugID, endUserID primitive.ObjectID, expDate time.Time, quantity int) *Expiration {
	return &Expiration{
		ID:             primitive.NewObjectID(),
		DrugID:         drugID,
		EndUserID:      endUserID,
		ExpirationDate: expDate,
		Quantity:       quantity,
	}
}
