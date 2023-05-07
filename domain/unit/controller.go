package unit

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/GivailsonNeves/vacation-rental-api/domain"
	"github.com/labstack/echo/v4"
)

type (
	UnitRepository interface {
		Create(unit *domain.Unit) (*domain.Unit, error)
		Delete(unit *domain.Unit) (*domain.Unit, error)
	}

	Controller struct {
		repo UnitRepository
	}
)

func NewController(repo UnitRepository) Controller {
	return Controller{
		repo: repo,
	}
}

func (r Controller) Create(c echo.Context) error {
	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)

	if err != nil || json_map["avenue"] == nil || json_map["number"] == nil || json_map["type"] == nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	booking, _ := r.repo.Create(&domain.Unit{
		Avenue: fmt.Sprintf("%s", json_map["avenue"]),
		Number: fmt.Sprintf("%s", json_map["number"]),
		Photo:  fmt.Sprintf("%s", json_map["photo"]),
		Type:   fmt.Sprintf("%s", json_map["type"]),
	})
	return c.JSON(http.StatusOK, booking)
}

func (r Controller) Delete(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	booking, _ := r.repo.Delete(&domain.Unit{
		ID: uint(id),
	})
	return c.JSON(http.StatusOK, booking)
}
