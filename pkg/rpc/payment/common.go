package payment

import (
	"strconv"

	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/api/errors"
	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jinzhu/gorm"

	"dev-gitlab.wanxingrowth.com/fanli/payment/pkg/constant"
	"dev-gitlab.wanxingrowth.com/fanli/payment/pkg/model/payment"
	"dev-gitlab.wanxingrowth.com/fanli/payment/pkg/rpc/protos"
)

func convertToPaymentInformation(paymentStruct *payment.Payment) *protos.PaymentInformation {
	return &protos.PaymentInformation{
		PaymentId:  paymentStruct.PaymentId,
		ShopId:     paymentStruct.ShopId,
		UserId:     paymentStruct.UserId,
		Amount:     paymentStruct.Amount,
		Status:     convertModelPaymentStatusToProtobuf(paymentStruct.Status),
		CreateTime: paymentStruct.CreatedAt.Unix(),
	}
}

func convertModelPaymentStatusToProtobuf(status payment.Status) protos.PaymentStatus {
	switch status {
	case payment.Completed:
		return protos.PaymentStatus_Completed
	default:
		return protos.PaymentStatus_NotPay
	}
}

func getPayment(paymentId uint64) (paymentData *payment.Payment, err error) {
	paymentData = new(payment.Payment)
	db := database.GetDB(constant.DatabaseConfigKey).
		Model(&payment.Payment{})

	err = db.Where(payment.Payment{PaymentId: paymentId}).First(&paymentData).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, err
		}
		return nil, err
	}

	return
}

func ConvertErrorToProtobuf(err error) *protos.Error {
	if validationError, ok := err.(validation.Error); ok {
		errorCode, convertError := strconv.Atoi(validationError.Code())
		if convertError != nil {
			errorCode = errors.CodeServerInternalError
		}
		return &protos.Error{
			Code:    int64(errorCode),
			Message: validationError.Error(),
		}
	}

	return &protos.Error{
		Code:    errors.CodeServerInternalError,
		Message: err.Error(),
	}
}
