package repository

import (
	"context"
	"isa-stats/pkg/model"
)

type Outbound interface {
	Save(ctx context.Context, out *model.Outbound) error
}
