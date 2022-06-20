package ports

import "context"

type Auth interface {
	GetNonce(ctx context.Context, address string) (string, error)
	RefreshNonce(ctx context.Context, address string) error
}
