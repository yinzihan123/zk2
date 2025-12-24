package main

import (
	"flag"
	"google.golang.org/grpc"
	"log"
	"net"
	_ "shop-srv/basic/init"
	__ "shop-srv/basic/proto"
	"shop-srv/handler/service"
)

// server is used to implement helloworld.GreeterServer.

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	__.RegisterShopServer(s, &service.Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
