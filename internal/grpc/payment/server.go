package payment

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	protoPayment "github.com/iskanye/utilities-payment-proto/payment"
)

type PaymentStatus = protoPayment.PaymentStatus

const (
	PAYMENT_PENDING PaymentStatus = iota
	PAYMENT_OK
	PAYMENT_FAILED
)

type serverAPI struct {
	protoPayment.UnimplementedPaymentServer
	payment Payment
}

type Payment interface {
	ProcessPayment(
		ctx context.Context,
		amount int,
	) (protoPayment.PaymentStatus, error)
}

func Register(gRPCServer *grpc.Server, payment Payment) {
	protoPayment.RegisterPaymentServer(gRPCServer, &serverAPI{payment: payment})
}

func (s *serverAPI) ProcessPayment(
	ctx context.Context,
	in *protoPayment.PaymentRequest,
) (*protoPayment.PaymentResponse, error) {
	if in.Amount <= 0 {
		return &protoPayment.PaymentResponse{
			Status: PAYMENT_FAILED,
		}, status.Error(codes.InvalidArgument, "amount must be positive")
	}

	statusCode, err := s.payment.ProcessPayment(ctx, int(in.GetAmount()))
	if err != nil {
		return &protoPayment.PaymentResponse{
			Status: PAYMENT_FAILED,
		}, status.Error(codes.Internal, err.Error())
	}

	return &protoPayment.PaymentResponse{
		Status: statusCode,
	}, nil
}
