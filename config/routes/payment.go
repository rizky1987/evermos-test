package routes

import (
	"evermos-test/config/env"
	"evermos-test/database/connection/mongo"
	httpHelper "evermos-test/helper"
	"evermos-test/http/api"
	"github.com/labstack/echo/v4"
)

func RegisterPaymentRoutes(baseEndpointGroup *echo.Group, httpHelper httpHelper.HTTPHelper, config env.Config) {

	group := baseEndpointGroup.Group("payment")
	{
		handler := api.PaymentHandler{
			Helper							: httpHelper,
			Config							: config,
			PaymentRepository 				: mongo.PaymentRepository,
			CartRepository 					: mongo.CartRepository,
			ProductRepository 				: mongo.ProductRepository,
			InventoryAdjustmentRepository 	: mongo.InventoryAdjustmentRepository,
		}

		group.POST("/callback", handler.Callback)
	}

}
