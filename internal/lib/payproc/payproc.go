// Short for Payment Processor
package payproc

import (
	"context"

	"github.com/iskanye/utilities-payment-payment/internal/grpc/payment"
)

type PaymentProcessor struct{}

func New() *PaymentProcessor {
	return &PaymentProcessor{}
}

func (p *PaymentProcessor) ProcessPayment(
	ctx context.Context,
	amount int,
) (payment.PaymentStatus, error) {
	const op = "lib.payproc.ProcessPayment"

	// Some payment processing here...

	return payment.PAYMENT_OK, nil
}
