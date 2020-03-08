package main

import (
	"github.com/ipfans/project-layout/initializers"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			initializers.NewConfigure,
			initializers.NewLogger,
			initializers.NewXORM,
			initializers.NewRedis,
			initializers.NewHTTPService,
			initializers.NewGRPCService,
		),
		fx.Invoke(
			initializers.StartHTTP,
			initializers.StartGRPC,
		),
		fx.Logger(&log.Logger),
	)
	app.Run()
}
