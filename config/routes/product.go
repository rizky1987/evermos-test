package routes

import (
	"evermos-test/config/env"
	"evermos-test/database/connection/mongo"
	httpHelper "evermos-test/helper"
	"evermos-test/http/api"
	"github.com/labstack/echo/v4"
)

func RegisterProductRoutes(baseEndpointGroup *echo.Group, httpHelper httpHelper.HTTPHelper, config env.Config) {

	group := baseEndpointGroup.Group("product")
	{
		handler := api.ProductHandler{
			Helper:         httpHelper,
			Config:         config,
			ProductRepository: mongo.ProductRepository,
		}

		group.POST("/find-all", handler.FindAll)
		group.POST("", handler.CreateProduct)
		group.PUT("/:id", handler.UpdateProduct)
		group.GET("/:id", handler.FindById)
	}

}
