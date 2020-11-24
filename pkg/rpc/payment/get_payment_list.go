package payment

import (
	rpcLog "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/api/rpc/utils/log"
	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"golang.org/x/net/context"

	"dev-gitlab.wanxingrowth.com/fanli/payment/pkg/constant"
	"dev-gitlab.wanxingrowth.com/fanli/payment/pkg/model/payment"
	"dev-gitlab.wanxingrowth.com/fanli/payment/pkg/rpc/protos"
	"dev-gitlab.wanxingrowth.com/fanli/payment/pkg/utils/log"
)

func (_ Controller) GetPaymentList(ctx context.Context, req *protos.GetPaymentListRequest) (*protos.GetPaymentListReply, error) {
	logger := rpcLog.WithRequestId(ctx, log.GetLogger())
	err := validateGetPaymentList(req)
	if err != nil {
		return &protos.GetPaymentListReply{
			Err: ConvertErrorToProtobuf(err),
		}, nil
	}

	var condition = payment.Payment{ShopId: req.GetShopId()}
	if req.GetStatus() > 0 {
		condition.Status = payment.Status(req.GetStatus())
	}

	var paymentList []*payment.Payment
	var count uint64

	db := database.GetDB(constant.DatabaseConfigKey).
		Model(&payment.Payment{}).Order("`PaymentId` DESC")

	err = db.Where(condition).Find(&paymentList).Count(&count).Error
	if err != nil {
		logger.WithError(err).Error("get payment list error")
		return &protos.GetPaymentListReply{
			Err: &protos.Error{
				Code:    constant.ErrorCodeGetPaymentListFailed,
				Message: constant.ErrorMessageGetPaymentListFailed,
			},
		}, nil
	}

	paymentInformationList := make([]*protos.PaymentInformation, 0, len(paymentList))
	for _, paymentStruct := range paymentList {
		paymentInformationList = append(paymentInformationList, convertToPaymentInformation(paymentStruct))
	}

	return &protos.GetPaymentListReply{
		PaymentInformationList: paymentInformationList,
		Count:                  count,
	}, nil
}

func validateGetPaymentList(req *protos.GetPaymentListRequest) error {
	return validation.ValidateStruct(req,
		validation.Field(&req.ShopId, ShopIdRule...),
	)
}
