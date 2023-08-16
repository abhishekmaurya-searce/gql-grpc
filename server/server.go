package server

import (
	"context"
	"greeting/greet"
	pb "greeting/proto"
)

type Server struct {
	pb.GreetServiceServer
}

func (s *Server) Greet(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	return &pb.Response{
		Res: greet.Greet(req.Req),
	}, nil
}
