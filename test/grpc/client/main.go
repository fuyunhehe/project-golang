package main

import (
	"os"
	"context"
	"time"
	pb "github.com/fuyunhehe/project-golang/test/grpc/helloworld"
	"google.golang.org/grpc"
	"log"
)

const (
	defaultName = "world"
)

func main() {
	t := time.Now().Unix()
	log.Println(t)
	conn, err := grpc.Dial("127.0.0.1:55001", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err.Error())
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	t = time.Now().Unix()
	log.Println(t)
	for i := 0; i < 10000; i++ {
		_, err := c.SayHello(ctx, &pb.HelloRequest{
			Name: name,
		})
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	t = time.Now().Unix()
	log.Println(t)


}
