package mucp

import (
	"context"
	"sync"

	"github.com/stack-labs/stack-rpc/server"
)

type serverKey struct{}

func wait(ctx context.Context) *sync.WaitGroup {
	if ctx == nil {
		return nil
	}
	wg, ok := ctx.Value("wait").(*sync.WaitGroup)
	if !ok {
		return nil
	}
	return wg
}

func FromContext(ctx context.Context) (server.Server, bool) {
	c, ok := ctx.Value(serverKey{}).(server.Server)
	return c, ok
}

func NewContext(ctx context.Context, s server.Server) context.Context {
	return context.WithValue(ctx, serverKey{}, s)
}
