package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/iskanye/utilities-payment-payment/internal/app"
	"github.com/iskanye/utilities-payment-payment/internal/config"
	pkgConfig "github.com/iskanye/utilities-payment/pkg/config"
	"github.com/iskanye/utilities-payment/pkg/logger"
)

func main() {
	cfg := pkgConfig.MustLoad[config.Config]()
	log := logger.SetupPrettySlog()
	app := app.New(log, cfg.Billing.Host, cfg.Billing.Port, cfg.Port)

	go func() {
		app.PaymentServer.MustRun()
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	app.PaymentServer.Stop()
	log.Info("Gracefully stopped")
}
