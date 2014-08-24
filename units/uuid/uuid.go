package uuid

import (
	"code.google.com/p/go-uuid/uuid"
	"code.google.com/p/go.net/context"
)

func Call(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, "id", uuid.New())

	return ctx
}
