package payment

import (
	rpcLog "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/api/rpc/utils/log"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"golang.org/x/net/context"

	"dev-gitlab.wanxingrowth.com/fanli/payment/pkg/constant"
	"dev-gitlab.wanxingrowth.com/fanli/payment/pkg/model/payment"
	"dev-gitlab.wanxingrowth.com/fanli/payment/pkg/rpc/protos"
	"dev-gitlab.wanxingrowth.com/fanli/payment/pkg/utils/log"
)

func (_ Controller) GetPayment(ctx context.Context, req *protos.GetPaymentRequest) (*protos.GetPaymentReply, error) {

	logger := rpcLog.WithRequestId(ctx, log.GetLogger())

	err := validateGetPayment(req)
	if err != nil {
		return &protos.GetPaymentReply{
			Err: ConvertErrorToProtobuf(err),
		}, nil
	}

	var paymentData *payment.Payment

	paymentData, err = getPayment(req.GetPaymentId())
	if err != nil {
		logger.WithError(err).Error("get payment error")
		return &protos.GetPaymentReply{
			Err: &protos.Error{
				Code:    constant.ErrorCodeGetPaymentFailed,
				Message: constant.ErrorMessageGetPaymentFailed,
			},
		}, nil
	}

	return &protos.GetPaymentReply{
		PaymentInformation: convertToPaymentInformation(paymentData),
	}, nil
}

func validateGetPayment(req *protos.GetPaymentRequest) error {
	return validation.ValidateStruct(req,
		validation.Field(&req.PaymentId, PaymentIdRule...),
		validation.Field(&req.ShopId, ShopIdRule...),
	)
}
