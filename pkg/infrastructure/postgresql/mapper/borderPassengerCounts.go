package mapper

import (
	"isa-stats/pkg/model"
	"strconv"
	"time"
)

type BorderPassengerCountsMapper struct {
	Id          int       `db:"id"`
	Direction   int       `db:"direction"`
	Year        int       `db:"year"`
	Month       int       `db:"month"`
	AirportName string    `db:"airport_name"`
	JapanNum    int       `db:"japan_num"`
	ForeignNum  int       `db:"foreign_num"`
	TreatyNum   int       `db:"treaty_num"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

const (
	Inbound = iota + 1
	Outbound
)

func (m BorderPassengerCountsMapper) FromInboundModel(in *model.Inbound) *BorderPassengerCountsMapper {
	year, _ := strconv.Atoi(in.Date.Format("2006"))
	month, _ := strconv.Atoi(in.Date.Format("01"))
	return &BorderPassengerCountsMapper{
		Direction:   Inbound,
		Year:        year,
		Month:       month,
		AirportName: in.Airport,
		JapanNum:    in.Person.Japan,
		ForeignNum:  in.Person.Foreign,
		TreatyNum:   in.Person.Treaty,
	}
}

func (m BorderPassengerCountsMapper) FromOutboundModel(out *model.Outbound) *BorderPassengerCountsMapper {
	year, _ := strconv.Atoi(out.Date.Format("2006"))
	month, _ := strconv.Atoi(out.Date.Format("01"))
	return &BorderPassengerCountsMapper{
		Direction:   Outbound,
		Year:        year,
		Month:       month,
		AirportName: out.Airport,
		JapanNum:    out.Person.Japan,
		ForeignNum:  out.Person.Foreign,
		TreatyNum:   out.Person.Treaty,
	}
}
func NewBorderPassengerCountsMapper() *BorderPassengerCountsMapper {
	return &BorderPassengerCountsMapper{}
}
