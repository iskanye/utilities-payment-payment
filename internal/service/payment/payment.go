package payment

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/iskanye/utilities-payment-payment/internal/grpc/payment"
	"github.com/iskanye/utilities-payment/pkg/logger"
)

type Payment struct {
	log              *slog.Logger
	paymentProcessor PaymentProcessor
}

type PaymentProcessor interface {
	ProcessPayment(
		ctx context.Context,
		amount int,
	) (payment.PaymentStatus, error)
}

func New(
	log *slog.Logger,
	paymentProcessor PaymentProcessor,
) *Payment {
	return &Payment{
		log:              log,
		paymentProcessor: paymentProcessor,
	}
}

func (b *Payment) ProcessPayment(
	ctx context.Context,
	amount int,
) (payment.PaymentStatus, error) {
	const op = "Payment.ProcessPayment"

	log := b.log.With(
		slog.String("op", op),
		slog.Int("amount", amount),
	)

	log.Info("attempting to process payment")

	status, err := b.paymentProcessor.ProcessPayment(ctx, amount)
	if err != nil {
		log.Error("failed to process payment", logger.Err(err))
		return payment.PAYMENT_FAILED, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("payment processed",
		slog.String("status", status.String()),
	)

	return status, nil
}
