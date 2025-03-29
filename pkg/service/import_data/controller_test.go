package import_data

import (
	"context"
	"go.uber.org/mock/gomock"
	"strings"
	"testing"
)

func TestController_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	svc := NewMockService(ctrl)

	c := NewController(svc)
	c.Execute(context.TODO(), strings.NewReader(""))
}
