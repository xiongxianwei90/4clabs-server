package ports

import "context"

type Script interface {
	RegisterUpdate(context.Context, string, string, string) error
	ComicWorksUpdate(context.Context, bool) error
	ComicWorksSold(ctx context.Context, to string, tokenId int64) error
}
