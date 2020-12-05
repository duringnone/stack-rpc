// Package mucp provides an mucp server
package mucp

import (
	"github.com/stack-labs/stack-rpc/server"
)

var (
	DefaultRouter = newRpcRouter()
)

// NewServer returns a stack server interface
func NewServer(opts ...server.Option) server.Server {
	return newRpcServer(opts...)
}
