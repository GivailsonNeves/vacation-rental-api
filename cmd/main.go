package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/GivailsonNeves/vacation-rental-api/domain/booking"
	"github.com/GivailsonNeves/vacation-rental-api/domain/guest"
	"github.com/GivailsonNeves/vacation-rental-api/domain/unit"
	"github.com/GivailsonNeves/vacation-rental-api/domain/user"
	"github.com/GivailsonNeves/vacation-rental-api/storage"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()
	storage.NewDB()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// e.Use(echojwt.WithConfig(echojwt.Config{
	// 	SigningKey: []byte("secret"),
	// }))

	e.GET("/", hello)
	booking.InitModule(e)
	guest.InitModule(e)
	unit.InitModule(e)
	user.InitModule(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, Wolrd!")
}
