package guest

import (
	"github.com/GivailsonNeves/vacation-rental-api/domain"
	"github.com/labstack/echo/v4"
)

type (
	GuestRepository interface {
		Create(guest *domain.Guest) (*domain.Guest, error)
	}

	Controller struct {
		repo GuestRepository
	}
)

func NewController(repo GuestRepository) Controller {
	return Controller{
		repo: repo,
	}
}

func (r Controller) Create(c echo.Context) error {
	return nil
}
