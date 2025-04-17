package drug

import "context"

type Repository interface {
	Create(ctx context.Context, drug *Drug) error
	FindAll(ctx context.Context) ([]*Drug, error)
}
