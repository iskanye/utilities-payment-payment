// Short for Payment Processor
package payproc

import (
	"context"
	"fmt"

	"github.com/iskanye/utilities-payment-payment/internal/grpc/billing"
	"github.com/iskanye/utilities-payment-payment/internal/grpc/payment"
)

type PaymentProcessor struct {
	billing billing.Billing
}

func New(
	billing billing.Billing,
) *PaymentProcessor {
	return &PaymentProcessor{
		billing: billing,
	}
}

func (p *PaymentProcessor) ProcessPayment(
	ctx context.Context,
	amount int,
	billID int64,
) (payment.PaymentStatus, error) {
	const op = "lib.payproc.ProcessPayment"

	err := p.billing.PayBill(ctx, billID)
	if err != nil {
		return payment.PAYMENT_FAILED, fmt.Errorf("%s: %w", op, err)
	}

	return payment.PAYMENT_OK, nil
}
