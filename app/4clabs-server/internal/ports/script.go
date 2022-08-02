package ports

import "context"

type Script interface {
	RegisterUpdate(context.Context, string, string, string) error
}
