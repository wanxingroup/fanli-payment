package constant

const (
	ErrorCodePaymentIdEmpty    = "417001"
	ErrorMessagePaymentIdEmpty = "订单 ID 为必填"
)

const (
	ErrorCodeShopIdEmpty    = "417002"
	ErrorMessageShopIdEmpty = "店铺 ID 为必填"
)

const (
	ErrorCodeUserIdEmpty    = "417003"
	ErrorMessageUserIdEmpty = "用户 ID 为必填"
)

const (
	ErrorCodeAmountEmpty    = "417004"
	ErrorMessageAmountEmpty = "金额为必填"
)

const (
	ErrorCodeNoteLengthOutOfRange    = "417005"
	ErrorMessageNoteLengthOutOfRange = "备注长度超出长度"
)

const (
	ErrorCodeCreatePaymentFailed    = 517001
	ErrorMessageCreatePaymentFailed = "创建支付失败，内部服务暂时不可用"
)

const (
	ErrorCodeGetPaymentFailed    = 517002
	ErrorMessageGetPaymentFailed = "获取订单失败，内部服务暂时不可用"
)

const (
	ErrorCodeGetPaymentListFailed    = 517003
	ErrorMessageGetPaymentListFailed = "获取订单列表失败，内部服务暂时不可用"
)

const (
	ErrorCodePaymentNotExist    = 517004
	ErrorMessagePaymentNotExist = "订单不存在"
)

const (
	ErrorCodePaymentStatus    = 517005
	ErrorMessagePaymentStatus = "订单状态不正确"
)
