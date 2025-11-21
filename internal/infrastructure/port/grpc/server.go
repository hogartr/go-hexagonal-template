package grpcport

import (
	"context"
	"fmt"
	"net"

	"github.com/hogartr/go-hexagonal-template/internal/infrastructure/config"
	"github.com/hogartr/go-hexagonal-template/proto"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

func NewGRPCServer(lc fx.Lifecycle, cfg config.Config, userHandler *UserServer) (*grpc.Server, error) {
	addr := fmt.Sprintf(":%s", cfg.GRPCPort)
	srv := grpc.NewServer()

	proto.RegisterUserServiceServer(srv, userHandler)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			lis, err := net.Listen("tcp", addr)
			if err != nil {
				return fmt.Errorf("failed to listen on %s: %w", addr, err)
			}
			fmt.Printf("gRPC server listening on %s (env: %s)\n", addr, cfg.Environment)
			go srv.Serve(lis)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			srv.GracefulStop()
			fmt.Println("gRPC server stopped")
			return nil
		},
	})

	return srv, nil
}
