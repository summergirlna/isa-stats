package repository

import "context"

type Outbound interface {
	Save(ctx context.Context, out *Outbound) error
}
