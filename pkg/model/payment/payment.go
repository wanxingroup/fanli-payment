package payment

import (
	databases "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database/models"
)

const TableNamePayment = "payment"

type Payment struct {
	PaymentId uint64 `gorm:"column:paymentId;type:bigint unsigned;primary_key;comment:'ID'"`
	ShopId    uint64 `gorm:"column:shopId;type:bigint unsigned;not null;default:'0';index:shopId;comment:'店铺ID'"`
	UserId    uint64 `gorm:"column:userId;type:bigint unsigned;not null;default:'0';index:userId;comment:'用户ID'"`
	Amount    uint64 `gorm:"column:amount;type:bigint unsigned;not null;default:'0';comment:'金额'"`
	Note      string `gorm:"column:note;type:varchar(200);not null;default:'';comment:'备注'"`
	Status    Status `gorm:"column:status;type:tinyint unsigned;not null;default:'0';index:status;comment:'状态'"`
	databases.BasicTimeFields
}

func (payment *Payment) TableName() string {
	return TableNamePayment
}
