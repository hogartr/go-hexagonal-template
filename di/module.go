package di

import (
	"github.com/hogartr/go-hexagonal-template/internal/application/usecase"
	"github.com/hogartr/go-hexagonal-template/internal/infrastructure/clock"
	"github.com/hogartr/go-hexagonal-template/internal/infrastructure/config"
	"github.com/hogartr/go-hexagonal-template/internal/infrastructure/db"
	grpcport "github.com/hogartr/go-hexagonal-template/internal/infrastructure/port/grpc"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

var Module = fx.Options(
	// 1. Config first
	fx.Provide(config.Load),

	// 2. Clock
	fx.Provide(clock.NewRealClock),

	// 3. Database (now uses config)
	fx.Provide(db.NewUserRepoFromConfig), // see below

	// 4. Use Cases
	fx.Provide(
		usecase.NewCreateUserUseCase,
	),

	// 5. Handlers & Server
	fx.Provide(grpcport.NewUserServer),
	fx.Provide(grpcport.NewGRPCServer),

	// 6. Start server
	fx.Invoke(func(*grpc.Server) {}),
)
