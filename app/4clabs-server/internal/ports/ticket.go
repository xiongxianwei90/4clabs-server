package ports

import "context"

type TicketRepo interface {
	Query(ctx context.Context, address string) (bool, uint32, error)
}
