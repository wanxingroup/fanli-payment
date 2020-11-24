package payment

import (
	"context"
	"testing"

	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/api/gin/request/requestid"
	"github.com/stretchr/testify/assert"

	"dev-gitlab.wanxingrowth.com/fanli/payment/pkg/rpc/protos"
)

func TestCreatePayment(t *testing.T) {
	tests := []struct {
		ctx   context.Context
		input *protos.CreatePaymentRequest
		want  uint64
	}{
		{
			ctx: context.WithValue(context.Background(), requestid.Key, "test_request_id"),
			input: &protos.CreatePaymentRequest{
				ShopId: 12344,
				UserId: 123456,
				Amount: 1,
				Note:   "2个包子",
			},
			want: 0,
		},
	}

	for _, test := range tests {
		svc := &Controller{}
		reply, _ := svc.CreatePayment(test.ctx, test.input)
		if reply != nil {
			assert.Greater(t, reply.Payment.PaymentId, test.want, test)
		}

	}
}
