package app

import (
	"log/slog"

	"github.com/iskanye/utilities-payment-payment/internal/app/grpc"
	"github.com/iskanye/utilities-payment-payment/internal/grpc/billing"
	"github.com/iskanye/utilities-payment-payment/internal/lib/payproc"
	"github.com/iskanye/utilities-payment-payment/internal/service/payment"
)

type App struct {
	PaymentServer *grpc.PaymentApp
	BillintClient billing.Billing
}

func New(
	log *slog.Logger,
	billingHost string,
	billingPort int,
	grpcPort int,
) *App {
	billingClient, err := billing.New(billingHost, billingPort)
	if err != nil {
	}
	pp := payproc.New(billingClient)

	paymentService := payment.New(log, pp)
	paymentApp := grpc.New(log, paymentService, grpcPort)

	return &App{
		PaymentServer: paymentApp,
		BillintClient: billingClient,
	}
}
