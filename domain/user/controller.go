package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GivailsonNeves/vacation-rental-api/domain"
	"github.com/labstack/echo/v4"
)

type (
	UserRepository interface {
		Create(user *domain.User) (*domain.User, error)
		FindAll(paginationOptions *domain.PaginationInputType, filter *Filter) ([]domain.User, *domain.PaginationResultType, error)
	}

	Controller struct {
		repo UserRepository
	}
)

func NewController(repo UserRepository) Controller {
	return Controller{
		repo: repo,
	}
}

func (r Controller) Create(c echo.Context) error {
	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)

	if err != nil ||
		json_map["name"] == nil ||
		json_map["email"] == nil ||
		json_map["phone"] == nil ||
		json_map["password"] == nil ||
		json_map["type"] == nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	user, _ := r.repo.Create(&domain.User{
		Name:     fmt.Sprintf("%s", json_map["name"]),
		Email:    fmt.Sprintf("%s", json_map["email"]),
		Phone:    fmt.Sprintf("%s", json_map["phone"]),
		Photo:    fmt.Sprintf("%s", json_map["photo"]),
		Password: fmt.Sprintf("%s", json_map["password"]),
		Type:     fmt.Sprintf("%s", json_map["type"]),
	})
	return c.JSON(http.StatusOK, user)
}
