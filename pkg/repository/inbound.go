package repository

import "context"

type Inbound interface {
	Save(ctx context.Context, in *Inbound) error
}
