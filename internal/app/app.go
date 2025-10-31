package app

import (
	"log/slog"

	"github.com/iskanye/utilities-payment-payment/internal/app/grpc"
	"github.com/iskanye/utilities-payment-payment/internal/lib/payproc"
	"github.com/iskanye/utilities-payment-payment/internal/service/payment"
)

type App struct {
	GRPCApp *grpc.PaymentApp
}

func New(
	log *slog.Logger,
	grpcPort int,
) *App {
	pp := payproc.New()

	paymentService := payment.New(log, pp)
	grpcApp := grpc.New(log, paymentService, grpcPort)

	return &App{
		GRPCApp: grpcApp,
	}
}
