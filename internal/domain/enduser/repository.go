package enduser

import "context"

type Repository interface {
	Create(ctx context.Context, endUser *EndUser) error
	FindAll(ctx context.Context) ([]*EndUser, error)
}
