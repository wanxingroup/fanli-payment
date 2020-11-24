package payment

import (
	rpcLog "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/api/rpc/utils/log"
	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"

	"dev-gitlab.wanxingrowth.com/fanli/payment/pkg/constant"
	"dev-gitlab.wanxingrowth.com/fanli/payment/pkg/model/payment"
	"dev-gitlab.wanxingrowth.com/fanli/payment/pkg/rpc/protos"
	"dev-gitlab.wanxingrowth.com/fanli/payment/pkg/utils/log"
)

func (_ Controller) FinishPayment(ctx context.Context, req *protos.FinishPaymentRequest) (*protos.FinishPaymentReply, error) {
	logger := rpcLog.WithRequestId(ctx, log.GetLogger())
	err := validateFinishPayment(req)
	if err != nil {
		return &protos.FinishPaymentReply{
			Err: ConvertErrorToProtobuf(err),
		}, nil
	}

	var paymentData *payment.Payment
	paymentData, err = getPayment(req.GetPaymentId())
	if err != nil {
		logger.WithError(err).Error("get payment error")
		return &protos.FinishPaymentReply{
			Err: &protos.Error{
				Code:    constant.ErrorCodeGetPaymentFailed,
				Message: constant.ErrorMessageGetPaymentFailed,
			},
		}, nil
	}

	if paymentData.Status == payment.Completed {
		return &protos.FinishPaymentReply{
			Payment: convertToPaymentInformation(paymentData),
		}, nil
	}

	if paymentData.Status != payment.NotPay {
		logger.WithError(err).Error("payment status is not NotPay")
		return &protos.FinishPaymentReply{
			Err: &protos.Error{
				Code:    constant.ErrorCodePaymentStatus,
				Message: constant.ErrorMessagePaymentStatus,
			},
		}, nil
	}

	_, updateStatusErr := updateStatus(paymentData, logger)
	if updateStatusErr != nil {
		return &protos.FinishPaymentReply{
			Err: ConvertErrorToProtobuf(err),
		}, nil
	}

	return &protos.FinishPaymentReply{
		Payment: convertToPaymentInformation(paymentData),
	}, nil
}

func validateFinishPayment(req *protos.FinishPaymentRequest) error {
	return validation.ValidateStruct(req,
		validation.Field(&req.PaymentId, PaymentIdRule...),
	)
}

func updateStatus(paymentData *payment.Payment, logger *logrus.Entry) (bool, error) {
	err := database.GetDB(constant.DatabaseConfigKey).Model(paymentData).Where(&payment.Payment{PaymentId: paymentData.PaymentId, Status: payment.NotPay}).UpdateColumn("status", payment.Completed).Error
	if err != nil {
		logger.WithError(err).Error("finish payment error")
		return false, err
	}

	paymentData.Status = payment.Completed
	return true, nil
}
