package grpc

import (
	"context"
	"net"

	"github.com/hnguyen111/download-manager/internal/generated/grpc/idm"
	"google.golang.org/grpc"
)

type Server interface {
	Start(ctx context.Context) error
}

type server struct {
	handler idm.DownloadServiceServer
}

func NewServer(handler idm.DownloadServiceServer) Server {
	return &server{
		handler: handler,
	}
}

func (s *server) Start(ctx context.Context) error {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		return err
	}
	defer listener.Close()
	server := grpc.NewServer()
	idm.RegisterDownloadServiceServer(server, s.handler)
	return server.Serve(listener)
}
