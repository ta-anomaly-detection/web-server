package route

import (
	"github.com/labstack/echo/v4"
	"github.com/ta-anomaly-detection/web-server/internal/delivery/http"
)

type RouteConfig struct {
	App               *echo.Echo
	UserController    *http.UserController
	ContactController *http.ContactController
	AddressController *http.AddressController
	AuthMiddleware    echo.MiddlewareFunc
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.POST("/api/users", c.UserController.Register)
	c.App.POST("/api/users/_login", c.UserController.Login)
}

func (c *RouteConfig) SetupAuthRoute() {
	c.App.Use(c.AuthMiddleware)
	c.App.DELETE("/api/users", c.UserController.Logout)
	c.App.PATCH("/api/users/_current", c.UserController.Update)
	c.App.GET("/api/users/_current", c.UserController.Current)

	c.App.GET("/api/contacts", c.ContactController.List)
	c.App.POST("/api/contacts", c.ContactController.Create)
	c.App.PUT("/api/contacts/:contactId", c.ContactController.Update)
	c.App.GET("/api/contacts/:contactId", c.ContactController.Get)
	c.App.DELETE("/api/contacts/:contactId", c.ContactController.Delete)

	c.App.GET("/api/contacts/:contactId/addresses", c.AddressController.List)
	c.App.POST("/api/contacts/:contactId/addresses", c.AddressController.Create)
	c.App.PUT("/api/contacts/:contactId/addresses/:addressId", c.AddressController.Update)
	c.App.GET("/api/contacts/:contactId/addresses/:addressId", c.AddressController.Get)
	c.App.DELETE("/api/contacts/:contactId/addresses/:addressId", c.AddressController.Delete)
}