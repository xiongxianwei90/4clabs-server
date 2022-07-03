package ports

import (
 "4clabs-server/app/4clabs-server/internal/domain/entity"
 "context"
)

type Comic interface {
	List(ctx context.Context, userAddress string, limit uint32, nextScore int64) ([]entity.Comic, int64, uint32, bool, error)
    Create(ctx context.Context, comic entity.Comic)
}
