package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/GivailsonNeves/vacation-rental-api/domain/book"
	"github.com/GivailsonNeves/vacation-rental-api/storage"
	"github.com/joho/godotenv"
	"github.com/labstack/echo-jwt/v4"
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
	e.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte("secret"),
	}))

	e.GET("/", hello)
	book.InitModule(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, Wolrd!")
}
