package expiration

import (
	"context"
	"time"

	"github.com/yhartanto178dev/pharmabot/internal/domain/drug"
	"github.com/yhartanto178dev/pharmabot/internal/domain/enduser"
	"github.com/yhartanto178dev/pharmabot/internal/domain/expiration"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (s *Service) CreateExpiration(ctx context.Context, drugID, endUserID primitive.ObjectID, expDate time.Time, quantity int) (*expiration.Expiration, error) {

	newExp := expiration.NewExpiration(drugID, endUserID, expDate, quantity)
	if err := s.expRepo.Create(ctx, newExp); err != nil {
		return nil, err
	}
	return newExp, nil
}
