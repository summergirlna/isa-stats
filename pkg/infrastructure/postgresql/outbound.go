package postgresql

import (
	"context"
	"github.com/jmoiron/sqlx"
	"isa-stats/pkg/infrastructure/postgresql/mapper"
	"isa-stats/pkg/model"
	"isa-stats/pkg/repository"
)

type Outbound struct {
	db *sqlx.DB
}

func (o Outbound) Save(ctx context.Context, out *model.Outbound) error {
	mp := mapper.NewBorderPassengerCountsMapper().FromOutboundModel(out)

	tx, ok := GetTx(ctx)
	if !ok {
		tx = o.db.MustBegin()
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

func NewOutbound(db *sqlx.DB) repository.Outbound {
	return &Outbound{db: db}
}
