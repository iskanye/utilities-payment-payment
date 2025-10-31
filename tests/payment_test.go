package tests

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/iskanye/utilities-payment-payment/tests/suite"
	"github.com/iskanye/utilities-payment-proto/payment"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProcessPayment_Success(t *testing.T) {
	ctx, s := suite.New(t)

	amount := gofakeit.Number(100, 100000)

	resp, err := s.PaymentClient.ProcessPayment(ctx, &payment.PaymentRequest{
		Amount: int32(amount),
	})
	require.NoError(t, err)
	assert.NotEmpty(t, resp)

	assert.Equal(t, payment.PaymentStatus_OK, resp.GetStatus())
}
