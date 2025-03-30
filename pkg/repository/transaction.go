package repository

import "context"

type Transaction interface {
	Execute(ctx context.Context, f func(ctx context.Context) (any, error)) (any, error)
}
