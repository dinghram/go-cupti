// +build !linux !cgo arm64

package cuptigrpc

import (
	cupti "github.com/c3sr/go-cupti"
	"google.golang.org/grpc"
)

func ServerUnaryInterceptor(_ ...cupti.Option) grpc.UnaryServerInterceptor {
	return noopUnaryServer
}
