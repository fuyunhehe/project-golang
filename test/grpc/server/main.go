package main

import (
	"google.golang.org/grpc"
	pb "github.com/fuyunhehe/project-golang/test/grpc/helloworld"
	"google.golang.org/grpc/reflection"
	"net"
	"log"
	"context"
)

//go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto

type greeter struct {
}

func (s *greeter) SayHello(ctx context.Context, r *pb.HelloRequest) (resp *pb.HelloReply, err error) {
	resp = &pb.HelloReply{
		Message: "hello " + r.Name,
	}
	return resp, nil
}

func main() {
	var (
		port = 55001
	)
	lis, err := net.ListenTCP("tcp", &net.TCPAddr{
		Port: port,
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	s := grpc.NewServer()
	g := &greeter{}
	pb.RegisterGreeterServer(s, g)
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatal(err.Error())
	}
}
