package postgresql

import (
	"context"
	"github.com/jmoiron/sqlx"
	"isa-stats/pkg/infrastructure/postgresql/mapper"
	"isa-stats/pkg/model"
	"isa-stats/pkg/repository"
)

type Inbound struct {
	db *sqlx.DB
}

func (i Inbound) Save(ctx context.Context, in *model.Inbound) error {
	mp := mapper.NewBorderPassengerCountsMapper().FromInboundModel(in)

	tx, ok := GetTx(ctx)
	if !ok {
		tx = i.db.MustBegin()
	}

	_, err := tx.NamedExec(`
INSERT INTO border_passenger_counts (direction, year, month, airport_name, japan_num, foreign_num, treaty_num)
VALUES (:direction, :year, :month, :airport_name, :japan_num, :foreign_num, :treaty_num)
`, mp)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func NewInbound(db *sqlx.DB) repository.Inbound {
	return &Inbound{db: db}
}
