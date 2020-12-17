package rpc

import (
	"context"
	"log"

	"grpce/gen/helloworld"
)

// RPC implements GreeterService
type RPC struct{}

func (*RPC) SayHello(_ context.Context, req *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	log.Printf("got %q request", req.ProtoReflect().Descriptor().FullName())

	resp := &helloworld.HelloReply{
		Message: "Hello, " + req.Name,
	}

	return resp, nil
}
