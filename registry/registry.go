package registry

import (
	"fmt"
	"github.com/feixiaobo/go-micro-registry/client"
	"github.com/feixiaobo/go-micro-registry/option"
	"github.com/feixiaobo/go-plugins/registry/eureka"
	"github.com/feixiaobo/go-plugins/server/http"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/server"
	"github.com/prometheus/common/log"
	"net"
	"time"
)

type Server struct {
	opts     option.Options
	registry registry.Registry
}

func (s *Server) Start() {
	go register(s)
}


func register(s *Server) {
	opts := s.opts

	if len(opts.RegistryAddress) == 0 {
		log.Errorf("the register address is required")
		panic("[error] the register address can't be null")
	}

	name := opts.Name
	if name == "" {
		log.Errorf("the server name is required")
		panic("[error] the server name can't be null")
	}
	ip := getLocalIP()
	port := opts.Port
	if port == 0 {
		log.Errorf("the server port is required")
		panic("[error] the server port can't be null")
	}
	ttl := opts.RegisterTTL
	if ttl == time.Duration(0) {
		ttl = time.Second * 30
	}

	addr := fmt.Sprintf("%s:%d", ip, port)
	instanceId := fmt.Sprintf("%s:%s:%d", ip, name, port)

	metaMap := opts.Metadata
	if metaMap == nil {
		metaMap = make(map[string]string)
	}
	metaMap["instanceId"] = instanceId

	ser := http.NewServer(
		server.Metadata(metaMap),
		server.Id(instanceId),
		server.Registry(s.registry),
		server.Address(addr),
		server.Name(name),
		server.Advertise(addr),
	)

	selector := selector.NewSelector(
		selector.Registry(s.registry),
		selector.SetStrategy(selector.RoundRobin),
	)

	client := client.InitClient(&s.registry, &selector, 3)

	service := micro.NewService(
		micro.Name(name),
		micro.Registry(s.registry),
		micro.Server(ser),
		micro.Address(addr),
		micro.Selector(selector),
		micro.Client(*client),
		micro.RegisterInterval(ttl),
	)

	service.Init()
	service.Run()
}

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	panic("Unable to determine local IP address (non loopback). Exiting.")
}
