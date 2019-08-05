package registry

import (
	"context"
	"github.com/feixiaobo/go-micro-registry/option"
	"github.com/feixiaobo/go-plugins/registry/eureka"
	"github.com/micro/go-micro/registry"
)

func EurekaServer(opts ...option.Option) Server {
	return newEurekaServer(opts...)
}

func newEurekaServer(opts ...option.Option) Server {
	ser := &Server{
		opts: option.Options{
			Context: context.Background(),
		},
	}

	for _, o := range opts {
		o(&ser.opts)
	}

	ser.registry = eureka.NewRegistry(
		registry.Addrs(ser.opts.RegistryAddress...),
	)
	return *ser
}
