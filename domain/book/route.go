package book

import (
	"github.com/GivailsonNeves/vacation-rental-api/storage"
	"github.com/labstack/echo/v4"
)

func InitModule(e *echo.Echo) {
	repo := NewRepository(storage.GetDBInstance())
	controller := NewController(repo)

	g := e.Group("book")
	g.GET("", controller.FindAll)
}
