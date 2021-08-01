package boot

import (
	"evermos-test/config/env"
	"evermos-test/database/connection/mongo"
	httpHelper "evermos-test/helper"
	_ "evermos-test/http/response"
	"fmt"

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

	_, err := dbMongo.Connect()
	if err != nil {
		fmt.Println("failed connect to mongoDB")
		panic(err.Error())
	}

	// End DB Connection

	// Begin Repository List

	// End repository List

	h.E.GET("/swagger/*", echoSwagger.WrapHandler)

	// Begin EndPoint List

	// Register User Endpoint

	// End EndPoint List

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
