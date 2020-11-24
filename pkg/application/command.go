package application

import (
	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/api/launcher"
	idCreator "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/utils/idcreator/snowflake"
	"github.com/spf13/viper"

	"dev-gitlab.wanxingrowth.com/fanli/payment/pkg/rpc/payment"

	"dev-gitlab.wanxingrowth.com/fanli/payment/pkg/config"
	"dev-gitlab.wanxingrowth.com/fanli/payment/pkg/rpc/protos"
	"dev-gitlab.wanxingrowth.com/fanli/payment/pkg/utils/log"
)

func Start() {

	app := launcher.NewApplication(
		launcher.SetApplicationDescription(
			&launcher.ApplicationDescription{
				ShortDescription: "payment service",
				LongDescription:  "support payment data management function.",
			},
		),
		launcher.SetApplicationLogger(log.GetLogger()),
		launcher.SetApplicationEvents(
			launcher.NewApplicationEvents(
				launcher.SetOnInitEvent(func(app *launcher.Application) {
					unmarshalConfiguration()
					registerPaymentRPCRouter(app)
					idCreator.InitCreator(app.GetServiceId())
				}),
				launcher.SetOnStartEvent(func(app *launcher.Application) {
					autoMigration()
				}),
			),
		),
	)

	app.Launch()
}

func registerPaymentRPCRouter(app *launcher.Application) {
	rpcService := app.GetRPCService()
	if rpcService == nil {
		log.GetLogger().WithField("stage", "onInit").Error("get rpc service is nil")
		return
	}

	protos.RegisterPaymentControllerServer(rpcService.GetRPCConnection(), &payment.Controller{})
}

func unmarshalConfiguration() {
	err := viper.Unmarshal(config.Config)
	if err != nil {
		log.GetLogger().WithError(err).Error("unmarshal config error")
	}
}
