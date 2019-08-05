package registry

import (
	"context"
	"github.com/feixiaobo/go-micro-registry/option"
	"github.com/feixiaobo/go-plugins/registry/eureka"
	"github.com/micro/go-micro/registry"
	"github.com/prometheus/common/log"
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

	if len(ser.opts.RegistryAddress) == 0 {
		log.Errorf("the register address is required")
		panic("[error] the register address can't be null")
	}
	ser.registry = eureka.NewRegistry(
		registry.Addrs(ser.opts.RegistryAddress...),
	)
	return *ser
}
