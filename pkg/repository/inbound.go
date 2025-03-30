package repository

import (
	"context"
	"isa-stats/pkg/model"
)

type Inbound interface {
	Save(ctx context.Context, in *model.Inbound) error
}
