package registry

import (
	"context"
	"github.com/feixiaobo/go-micro-registry/option"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
)

func ConsulServer(opts ...option.Option) Server {
	return newConsulServer(opts...)
}

func newConsulServer(opts ...option.Option) Server {
	ser := &Server{
		opts: option.Options{
			Context: context.Background(),
		},
	}

	for _, o := range opts {
		o(&ser.opts)
	}

	ser.registry = consul.NewRegistry(
		registry.Addrs(ser.opts.RegistryAddress...),
	)
	return *ser
}