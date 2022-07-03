package usecase

import (
	"4clabs-server/app/4clabs-server/internal/ports"
	"context"
)

type Ticket struct {
	ports.TicketRepo
}

func NewTicket(ticketRepo ports.TicketRepo) *Ticket {
	return &Ticket{TicketRepo: ticketRepo}
}

func (t *Ticket) CanMint(ctx context.Context, address string) (bool, uint32, error) {
	return t.TicketRepo.Query(ctx, address)
}
