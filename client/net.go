package client

import (
	"context"
	"fmt"
	http2 "github.com/feixiaobo/go-plugins/client/http"
	"github.com/google/martian/log"
	"github.com/micro/go-micro/client"
	client2 "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"time"
)

var httpClient client.Client

/**
	type User struct {
		Id           int32  `json:"id"`
		UserName     string `json:"userName"`
	}
	req ex: &User{Id: 123, UserName: "飞晓波"}
 	res ex: new(User)
*/
func Call(serviceName, path string, req interface{}, res interface{}, opts ...client.CallOption) (err error) {
	if httpClient == nil {
		return fmt.Errorf("error: client must be init before call request")
	}
	request := httpClient.NewRequest(serviceName, path, req)
	err = httpClient.Call(context.Background(), request, res, opts...)
	if err != nil {
		log.Errorf("call server error", err)
		return err
	}
	return err
}

func InitClient(register *registry.Registry, s *selector.Selector, retries int, timeout time.Duration) *client.Client {
	httpClient = http2.NewClient(
		client2.Retries(retries),
		client2.Registry(*register),
		client2.ContentType("application/json"),
		client2.Selector(*s),
		client2.RequestTimeout(timeout),
	)
	return &httpClient
}

