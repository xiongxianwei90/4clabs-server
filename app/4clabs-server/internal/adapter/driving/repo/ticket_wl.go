package repo

import (
	"4clabs-server/app/4clabs-server/internal/adapter/data/query"
	"4clabs-server/app/4clabs-server/internal/data"
	"4clabs-server/app/4clabs-server/internal/ports"
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var _ ports.TicketRepo = &Ticket{}

type Ticket struct {
	data *data.Data
}

func (t Ticket) Query(ctx context.Context, address string) (bool, uint32, error) {
	wl := query.Use(t.data.DB).TicketWl
	q, err := wl.WithContext(ctx).Where(wl.Address.Eq(address)).First()
	if err != nil && !errors.Is(gorm.ErrRecordNotFound, err) {
		return false, 0, errors.Wrapf(err, "address : %s", address)
	}
	if err != nil {
		return false, 0, nil
	}
	return true, uint32(q.Level), nil
}

func NewTicket(data *data.Data) *Ticket {
	return &Ticket{data: data}
}
