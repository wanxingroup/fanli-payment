package payment

import (
	"fmt"

	rpcLog "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/api/rpc/utils/log"
	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"
	idCreator "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/utils/idcreator/snowflake"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"golang.org/x/net/context"

	"dev-gitlab.wanxingrowth.com/fanli/payment/pkg/constant"
	"dev-gitlab.wanxingrowth.com/fanli/payment/pkg/model/payment"
	"dev-gitlab.wanxingrowth.com/fanli/payment/pkg/rpc/protos"
	"dev-gitlab.wanxingrowth.com/fanli/payment/pkg/utils/log"
)

func (_ Controller) CreatePayment(ctx context.Context, req *protos.CreatePaymentRequest) (*protos.CreatePaymentReply, error) {
	logger := rpcLog.WithRequestId(ctx, log.GetLogger())
	if req == nil {
		logger.Error("request data is nil")
		return nil, fmt.Errorf("request data is nil")
	}

	err := validateCreatePayment(req)
	if err != nil {
		return &protos.CreatePaymentReply{
			Err: ConvertErrorToProtobuf(err),
		}, nil
	}

	var paymentData *payment.Payment
	paymentData, err = createPayment(req)
	if err != nil {
		logger.WithError(err).Error("create payment error")
		return &protos.CreatePaymentReply{
			Err: &protos.Error{
				Code:    constant.ErrorCodeCreatePaymentFailed,
				Message: constant.ErrorMessageCreatePaymentFailed,
				Stack:   nil,
			},
		}, nil
	}

	return &protos.CreatePaymentReply{
		Payment: convertToPaymentInformation(paymentData),
	}, nil
}

func createPayment(req *protos.CreatePaymentRequest) (*payment.Payment, error) {
	record := &payment.Payment{
		PaymentId: idCreator.NextID(),
		ShopId:    req.ShopId,
		UserId:    req.UserId,
		Amount:    req.Amount,
		Note:      req.Note,
		Status:    payment.NotPay,
	}
	err := database.GetDB(constant.DatabaseConfigKey).Create(record).Error
	if err != nil {
		log.GetLogger().WithField("payment", record).WithError(err).Error("create record error")
		return nil, err
	}

	return record, nil
}

func validateCreatePayment(req *protos.CreatePaymentRequest) error {
	return validation.ValidateStruct(req,
		validation.Field(&req.ShopId, ShopIdRule...),
		validation.Field(&req.UserId, UserIdRule...),
		validation.Field(&req.Amount, AmountRule...),
		validation.Field(&req.Note, NoteRule...),
	)
}
