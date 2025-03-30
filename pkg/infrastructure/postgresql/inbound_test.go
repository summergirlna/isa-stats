package postgresql

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"isa-stats/pkg/model"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

func TestInbound_Save(t *testing.T) {
	db, err := sqlx.Connect("postgres",
		"host=localhost user=app password=app12345@ dbname=isa_stats sslmode=disable")
	assert.NoError(t, err)

	r := NewInbound(db)
	txManager := NewTransactionManager(db)

	in := &model.Inbound{
		Airport: "北海道_旭川（空港）",
		Person: &model.Person{
			Japan:   7330,
			Foreign: 3834,
			Treaty:  202,
		},
		Date: time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC),
	}

	ctx := context.TODO()
	_, err = txManager.Execute(ctx, func(ctx context.Context) (any, error) {
		return nil, r.Save(context.TODO(), in)
	})
	assert.NoError(t, err)
}
