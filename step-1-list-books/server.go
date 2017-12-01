package step_1_list_books

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	books "github.com/noaleibo1/grpc-workshop/start/books"
	"golang.org/x/net/context"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	fmt.Println("Server running at http://0.0.0.0:50051")
	service := service{}
	books.RegisterBookServiceServer(grpcServer, &service)
	grpcServer.Serve(lis)
}

type service struct {
}

func (s *service) List(context.Context, *books.Empty) (*books.Empty, error){
	return &books.Empty{}, status.Error(codes.Unimplemented, "The server does not implement this method")
}