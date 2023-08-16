package main

import (
	pb "greeting/proto"
	"greeting/server"
	"log"
	"net"

	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

func main() {
	var ser server.Server
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to Listen: %v", err)
	}
	log.Printf("Listning to: %s", addr)
	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &ser)
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to Serve: %v", err)
	}

}
