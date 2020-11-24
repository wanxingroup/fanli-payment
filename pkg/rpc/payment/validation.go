package payment

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"dev-gitlab.wanxingrowth.com/fanli/payment/pkg/constant"
)

var PaymentIdRule = []validation.Rule{
	validation.Required.ErrorObject(
		validation.NewError(constant.ErrorCodePaymentIdEmpty, constant.ErrorMessagePaymentIdEmpty),
	),
}

var ShopIdRule = []validation.Rule{
	validation.Required.ErrorObject(
		validation.NewError(constant.ErrorCodeShopIdEmpty, constant.ErrorMessageShopIdEmpty),
	),
}

var UserIdRule = []validation.Rule{
	validation.Required.ErrorObject(
		validation.NewError(constant.ErrorCodeUserIdEmpty, constant.ErrorMessageUserIdEmpty),
	),
}

var AmountRule = []validation.Rule{
	validation.Required.ErrorObject(
		validation.NewError(constant.ErrorCodeAmountEmpty, constant.ErrorMessageAmountEmpty),
	),
}

var NoteRule = []validation.Rule{
	validation.Length(0, 200).ErrorObject(validation.NewError(constant.ErrorCodeNoteLengthOutOfRange, constant.ErrorMessageNoteLengthOutOfRange)),
}
