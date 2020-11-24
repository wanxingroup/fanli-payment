package application

import (
	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"

	"dev-gitlab.wanxingrowth.com/fanli/payment/pkg/constant"
	"dev-gitlab.wanxingrowth.com/fanli/payment/pkg/model/payment"
)

func autoMigration() {
	db := database.GetDB(constant.DatabaseConfigKey)
	db.AutoMigrate(payment.Payment{})
}
