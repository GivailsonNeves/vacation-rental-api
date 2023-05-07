package booking

import (
	"github.com/GivailsonNeves/vacation-rental-api/storage"
	"github.com/labstack/echo/v4"
)

func InitModule(e *echo.Echo) {
	repo := NewRepository(storage.GetDBInstance())
	controller := NewController(repo)

	g := e.Group("booking")
	g.GET("", controller.FindAll)
	g.DELETE("/:id", controller.Delete)
	g.POST("", controller.Create)
	g.POST("/:id", controller.Update)
	g.GET("/:id", controller.Find)
}
