package graph

import (
	pb "greeting/proto"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Client pb.GreetServiceClient
}
