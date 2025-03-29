package import_data

import (
	"context"
	"io"
	"log/slog"
)

type Controller struct {
	service Service
}

func (c Controller) Execute(ctx context.Context, in io.Reader) error {
	slog.Info("start")
	return nil
}

func NewController(service Service) *Controller {
	return &Controller{service: service}
}
