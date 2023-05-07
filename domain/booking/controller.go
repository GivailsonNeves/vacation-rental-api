package booking

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/GivailsonNeves/vacation-rental-api/domain"
	"github.com/labstack/echo/v4"
)

type (
	BookingRepository interface {
		FindAll(paginationOptions *domain.PaginationInputType, filter *Filter) ([]domain.Booking, *domain.PaginationResultType, error)
		Create(booking *domain.Booking) (*domain.Booking, error)
		Delete(booking *domain.Booking) (*domain.Booking, error)
		Update(booking *domain.Booking) (*domain.Booking, error)
		Find(id uint64) (*domain.Booking, error)
	}

	Controller struct {
		repo BookingRepository
	}
)

func NewController(repo BookingRepository) Controller {
	return Controller{
		repo: repo,
	}
}

func (controller Controller) FindAll(c echo.Context) error {
	bookings, _, _ := controller.repo.FindAll(nil, nil)
	return c.JSON(http.StatusOK, bookings)
}

func (r Controller) Create(c echo.Context) error {

	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)

	if err != nil || json_map["name"] == nil || json_map["startAt"] == nil || json_map["endAt"] == nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}
	startDate, startDateError := time.Parse("2006-01-02", fmt.Sprintf("%s", json_map["startAt"]))
	endtDate, endDateError := time.Parse("2006-01-02", fmt.Sprintf("%s", json_map["endAt"]))

	if startDateError != nil || endDateError != nil {
		return c.JSON(http.StatusBadRequest, "invalid date format")
	}

	booking, _ := r.repo.Create(&domain.Booking{
		ID:      1,
		Name:    fmt.Sprintf("%s", json_map["name"]),
		StartAt: startDate,
		EndAt:   endtDate,
	})
	return c.JSON(http.StatusOK, booking)
}

func (r Controller) Delete(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	booking, _ := r.repo.Delete(&domain.Booking{
		ID: uint(id),
	})
	return c.JSON(http.StatusOK, booking)
}

func (r Controller) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	json_map := make(map[string]interface{})
	bodyErr := json.NewDecoder(c.Request().Body).Decode(&json_map)

	if bodyErr != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	bookingUpdated, _ := r.repo.Update(&domain.Booking{
		ID:   uint(id),
		Name: fmt.Sprintf("%s", json_map["name"]),
	})

	return c.JSON(http.StatusOK, bookingUpdated)
}

func (r Controller) Find(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}
	bookingFound, _ := r.repo.Find(uint64(id))

	return c.JSON(http.StatusOK, bookingFound)
}
