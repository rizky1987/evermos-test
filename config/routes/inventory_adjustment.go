package routes

import (
	"evermos-test/config/env"
	"evermos-test/database/connection/mongo"
	httpHelper "evermos-test/helper"
	"evermos-test/http/api"
	"github.com/labstack/echo/v4"
)

func RegisterInventoryAdjustmentRoutes(baseEndpointGroup *echo.Group, httpHelper httpHelper.HTTPHelper, config env.Config) {

	group := baseEndpointGroup.Group("inventory-adjustment")
	{
		handler := api.InventoryAdjustmentHandler{
			Helper							: httpHelper,
			Config							: config,
			InventoryAdjustmentRepository	: mongo.InventoryAdjustmentRepository,
			ProductRepository 				: mongo.ProductRepository,
		}

		group.POST("/find-all", handler.FindAll)
		group.POST("", handler.CreateInventoryAdjustment)
		group.GET("/:id", handler.FindById)
	}

}
