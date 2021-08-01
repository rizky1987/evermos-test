package boot

import (
	"evermos-test/config/env"
	"evermos-test/database/connection/mongo"
	"evermos-test/docs"
	httpHelper "evermos-test/helper"
	_ "evermos-test/http/response"
	"fmt"

	route "evermos-test/config/routes"
	ut "github.com/go-playground/universal-translator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"
	"gopkg.in/go-playground/validator.v9"
)

type HTTPHandler struct {
	E               *echo.Echo
	Config          env.Config
	Helper          httpHelper.HTTPHelper
	ValidatorDriver *validator.Validate
	Translator      ut.Translator
}

// RegisterApiHandler ...
func (h *HTTPHandler) RegisterApiHandler() *HTTPHandler {

	h.Helper = httpHelper.HTTPHelper{
		Validate:   h.ValidatorDriver,
		Translator: h.Translator,
	}

	// begin Db Connection
	dbMongo := mongo.Info{
		Hostname: h.Config.GetString("database.mongodb.host"),
		Database: h.Config.GetString("database.mongodb.database"),
		Username: h.Config.GetString("database.mongodb.username"),
		Password: h.Config.GetString("database.mongodb.password"),
	}

	dbMongo.Connect()

	// End DB Connection

	basePathAndVersion := fmt.Sprintf("/%v/%v/",
		h.Config.GetString("app.base_path"),
		h.Config.GetString("app.version"),
	)

	//Begin Global Swagger Configuration
	h.E.GET("/swagger/*", echoSwagger.WrapHandler)
	docs.SwaggerInfo.Title = "API V1 Evermos Test"
	docs.SwaggerInfo.Description = "This is API from Rizky Mochammad Soleh"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host =  h.Config.GetString("app.host")
	docs.SwaggerInfo.BasePath = basePathAndVersion
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	//End Global Swagger Configuration

	// Begin EndPoint List
	baseEndpointGroup := h.E.Group(basePathAndVersion)
	// baseEndpointGroup.Use(middlewareJWT.ValidateToken())

	// Begin Register All End Point
	route.RegisterProductRoutes(baseEndpointGroup, h.Helper, h.Config)
	route.RegisterCartRoutes(baseEndpointGroup, h.Helper, h.Config)
	route.RegisterCustomerRoutes(baseEndpointGroup, h.Helper, h.Config)

	return h
}

func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "evermos test/1.0")
		return next(c)
	}
}

// RegisterMiddleware ...
func (h *HTTPHandler) RegisterMiddleware() {
	h.E.Use(serverHeader)
	h.E.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	h.E.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	if h.Config.GetBool("app.debug") == true {
		h.E.Use(middleware.Logger())
		h.E.HideBanner = true
		h.E.Debug = true
	} else {
		h.E.HideBanner = true
		h.E.Debug = false
		h.E.Use(middleware.Recover())
	}
}
