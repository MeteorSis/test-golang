package client

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/MeteorSis/test-golang/grpc-interceptor/proto"
)

func DoRequest(port uint16) error {
	conn, err := grpc.Dial(
		fmt.Sprintf("localhost:%d", port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(
			func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
				log.Println("client interceptor 1_before")
				err := invoker(ctx, method, req, reply, cc, opts...)
				log.Println("client interceptor 1_after")
				return err
			},
			func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
				log.Println("client interceptor 2_before")
				err := invoker(ctx, method, req, reply, cc, opts...)
				log.Println("client interceptor 2_after")
				return err
			},
			func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
				log.Println("client interceptor 3_before")
				err := invoker(ctx, method, req, reply, cc, opts...)
				log.Println("client interceptor 3_after")
				return err
			},
		),
	)
	if err != nil {
		return fmt.Errorf("failed to dial: %w", err)
	}
	cli := proto.NewEchoClient(conn)
	_, err = cli.SendMessage(context.Background(), &proto.SendMessageRequest{Message: "hello"})
	if err != nil {
		return fmt.Errorf("failed to request: %w", err)
	}
	log.Println("client response")
	return nil
}
