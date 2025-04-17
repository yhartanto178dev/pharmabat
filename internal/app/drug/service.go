// internal/app/drug/service.go
package drug

import (
	"context"

	"github.com/yhartanto178dev/pharmabot/internal/domain/drug"
)

type Service struct {
	repo drug.Repository
}

func NewService(repo drug.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateDrug(ctx context.Context, name string) (*drug.Drug, error) {
	newDrug := drug.NewDrug(name)
	if err := s.repo.Create(ctx, newDrug); err != nil {
		return nil, err
	}
	return newDrug, nil
}
