package initializers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

func httpServe(lc fx.Lifecycle, srvType, addr string, logger zerolog.Logger, h http.Handler) {
	srv := http.Server{
		Addr:    addr,
		Handler: h,
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info().Str("addr", addr).Str("type", srvType).Msg("Starting services...")
			go func() {
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					logger.Error().Err(err).Msg("Listen failed.")
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			err := srv.Shutdown(ctx)
			return err
		},
	})
}

func StartHTTP(lc fx.Lifecycle, v *viper.Viper, logger zerolog.Logger, e *gin.Engine) {
	addr := v.GetString("http.listen")
	httpServe(lc, "HTTP", addr, logger, e)
}

func StartGRPC(lc fx.Lifecycle, v *viper.Viper, logger zerolog.Logger, r *grpc.Server) {
	addr := v.GetString("grpc.listen")
	httpServe(lc, "gRPC", addr, logger, r)
}
