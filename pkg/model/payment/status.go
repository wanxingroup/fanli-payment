package payment

//状态
type Status uint8

const (
	NotPay    Status = 1 //未支付
	Completed Status = 6 //订单完成
)
