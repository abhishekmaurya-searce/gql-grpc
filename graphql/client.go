package graph

import (
	pb "greeting/proto"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func SetupGQL() {
	var ser Resolver
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	ser.Client = pb.NewGreetServiceClient(conn)
	if err != nil {
		log.Fatalf("Failed to migrate schema: %v", err)
	}
	serve := handler.NewDefaultServer(
		NewExecutableSchema(
			Config{Resolvers: &mutationResolver{Resolver: &ser}},
		),
	)
	serve.AddTransport(transport.POST{})

	http.Handle("/greet", serve)
	http.ListenAndServe(":8080", nil)
}
