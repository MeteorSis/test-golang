package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/MeteorSis/test-golang/grpc-interceptor/proto"
)

type Server struct {
	srv *grpc.Server
	lis net.Listener
}

func New(port uint16) (*Server, error) {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		return nil, fmt.Errorf("failed to listen: %w", err)
	}

	srv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			func(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
				log.Println("server interceptor 1_before")
				res, err := handler(ctx, req)
				log.Println("server interceptor 1_after")
				return res, err
			},
			func(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
				log.Println("server interceptor 2_before")
				res, err := handler(ctx, req)
				log.Println("server interceptor 2_after")
				return res, err
			},
			func(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
				log.Println("server interceptor 3_before")
				res, err := handler(ctx, req)
				log.Println("server interceptor 3_after")
				return res, err
			},
		),
	)
	proto.RegisterEchoServer(srv, &server{})

	return &Server{srv: srv, lis: lis}, nil
}

func (s *Server) Serve() error {
	if err := s.srv.Serve(s.lis); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}

	return nil
}

func (s *Server) GracefulStop() {
	s.srv.GracefulStop()
}

type server struct {
	proto.UnimplementedEchoServer
}

func (*server) SendMessage(ctx context.Context, req *proto.SendMessageRequest) (*proto.SendMessageResponse, error) {
	log.Println("server handler")
	return &proto.SendMessageResponse{Message: req.Message}, nil
}
