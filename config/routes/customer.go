package routes

import (
	"evermos-test/config/env"
	"evermos-test/database/connection/mongo"
	httpHelper "evermos-test/helper"
	"evermos-test/http/api"
	"github.com/labstack/echo/v4"
)

func RegisterCustomerRoutes(baseEndpointGroup *echo.Group, httpHelper httpHelper.HTTPHelper, config env.Config) {

	group := baseEndpointGroup.Group("customer")
	{
		handler := api.CustomerHandler{
			Helper							: httpHelper,
			Config							: config,
			CustomerRepository				: mongo.CustomerRepository,
			InventoryAdjustmentRepository 	: mongo.InventoryAdjustmentRepository,
		}

		group.POST("/find-all", handler.FindAll)
		group.POST("", handler.CreateCustomer)
		group.PUT("/:id", handler.UpdateCustomer)
		group.GET("/:id", handler.FindById)
	}

}
