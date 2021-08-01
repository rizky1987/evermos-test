package routes

import (
	"evermos-test/config/env"
	"evermos-test/database/connection/mongo"
	httpHelper "evermos-test/helper"
	"evermos-test/http/api"
	"github.com/labstack/echo/v4"
)

func RegisterCartRoutes(baseEndpointGroup *echo.Group, httpHelper httpHelper.HTTPHelper, config env.Config) {

	group := baseEndpointGroup.Group("cart")
	{
		handler := api.CartHandler{
			Helper							: httpHelper,
			Config							: config,
			CartRepository					: mongo.CartRepository,
		}

		group.POST("/find-all", handler.FindAll)
		group.POST("", handler.CreateCart)
		group.PUT("/:id", handler.UpdateCart)
		group.GET("/:id", handler.FindById)
		group.POST("/checkout", handler.CheckoutCart)
	}

}
