package import_data

import "context"

type Input struct{}

type Output struct{}

type Service interface {
	Import(ctx context.Context, in *Input) (*Output, error)
}
