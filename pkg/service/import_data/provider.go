package import_data

import (
	"context"
)

type Provider struct{}

func (p Provider) Import(ctx context.Context, in *Input) (*Output, error) {
	//TODO implement me
	panic("implement me")
}

func NewProvider() Service {
	return &Provider{}
}
