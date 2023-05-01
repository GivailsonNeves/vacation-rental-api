package book

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/GivailsonNeves/vacation-rental-api/domain"
	"github.com/labstack/echo/v4"
)

type (
	BookRepository interface {
		FindAll(paginationOptions *domain.PaginationInputType, filter *Filter) ([]Book, *domain.PaginationResultType, error)
		Create(book *Book) (*Book, error)
	}

	Controller struct {
		repo BookRepository
	}
)

func NewController(repo BookRepository) Controller {
	return Controller{
		repo: repo,
	}
}

func (controller Controller) FindAll(c echo.Context) error {
	books, _, _ := controller.repo.FindAll(nil, nil)
	return c.JSON(http.StatusOK, books)
}

func (r Controller) Create(c echo.Context) error {

	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)

	if err != nil || json_map["name"] == nil || json_map["startAt"] == nil || json_map["endAt"] == nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	book, _ := r.repo.Create(&Book{
		ID:      1,
		Name:    fmt.Sprintf("%s", json_map["name"]),
		StartAt: time.Now(),
		EndAt:   time.Now(),
	})
	return c.JSON(http.StatusOK, book)
}
