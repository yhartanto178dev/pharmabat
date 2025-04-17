package enduser

import (
	"context"

	"github.com/yhartanto178dev/pharmabot/internal/domain/enduser"
)

type Service struct {
	repo enduser.Repository
}

func NewService(repo enduser.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateEndUser(ctx context.Context, name string) (*enduser.EndUser, error) {
	newEndUser := enduser.NewEndUser(name)
	if err := s.repo.Create(ctx, newEndUser); err != nil {
		return nil, err
	}
	return newEndUser, nil
}
