package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	books "github.com/noaleibo1/grpc-workshop/step-4-stream-added-books/books.com/books"
	"golang.org/x/net/context"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"github.com/olebedev/emitter"
)

var (
	port = flag.Int("port", 50051, "The server port")
	booksList = []*books.Book{
		{
			Id: 123,
			Title: "A Tale of Two Cities",
			Author: "Charles Dickens",
		},
	}
	newBookEmitter = emitter.Emitter{}
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

func (s *service) List(context.Context, *books.Empty) (*books.BookList, error){
	return &books.BookList{Books: booksList}, nil
}

func (s *service) Insert(ctx context.Context, book *books.Book) (*books.Empty, error) {
	booksList = append(booksList, book)
	newBookEmitter.Emit("NewBook")
	return &books.Empty{}, nil
}

func (s *service) Get(ctx context.Context, req *books.BookIdRequest) (*books.Book, error){
	for i := 0; i < len(booksList); i++ {
		if booksList[i].Id == req.Id {
			return booksList[i], nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "Not found")
}

func (s *service) Delete (ctx context.Context, req *books.BookIdRequest) (*books.Empty, error) {
	for i := 0; i < len(booksList); i++ {
		if booksList[i].Id == req.Id {
			booksList = append(booksList[:i], booksList[i+1:]...)
			return &books.Empty{}, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "Not found")
}

func (s *service) Watch(empty *books.Empty, stream books.BookService_WatchServer) error {
	c := newBookEmitter.On("NewBook")
	for {
		<-c
		booksListLength := len(booksList)
		stream.Send(booksList[booksListLength-1])
	}
	newBookEmitter.Off("NewBook")
	return nil
}