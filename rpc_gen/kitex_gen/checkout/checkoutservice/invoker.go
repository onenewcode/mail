// Code generated by Kitex v0.9.1. DO NOT EDIT.

package checkoutservice

import (
	server "github.com/cloudwego/kitex/server"
	checkout "rpc_gen/kitex_gen/checkout"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler checkout.CheckoutService, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
