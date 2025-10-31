package billing

import (
	"context"
	"net"
	"strconv"

	"github.com/iskanye/utilities-payment-proto/billing"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type clientApi struct {
	billing billing.BillingClient
}

type Billing interface {
	PayBill(
		ctx context.Context,
		billID int64,
	) error
}

func New(
	host string,
	port int,
) (*clientApi, error) {
	cc, err := grpc.NewClient(
		net.JoinHostPort(host, strconv.Itoa(port)),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return &clientApi{}, status.Error(codes.Unavailable, err.Error())
	}

	return &clientApi{billing.NewBillingClient(cc)}, nil
}

func (c *clientApi) PayBill(
	ctx context.Context,
	billID int64,
) error {
	_, err := c.billing.PayBill(ctx, &billing.PayRequest{BillId: billID})
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}
