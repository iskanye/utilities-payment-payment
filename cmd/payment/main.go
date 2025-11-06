package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/iskanye/utilities-payment-payment/internal/app"
	"github.com/iskanye/utilities-payment-payment/internal/config"
	pkgConfig "github.com/iskanye/utilities-payment-utils/pkg/config"
	"github.com/iskanye/utilities-payment-utils/pkg/logger"
)

func main() {
	cfg := pkgConfig.MustLoad[config.Config]()
	log := logger.SetupPrettySlog()
	app := app.New(log, cfg.GRPC.Port)

	go func() {
		app.GRPCApp.MustRun()
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	app.GRPCApp.Stop()
	log.Info("Gracefully stopped")
}
