package user

import (
	"github.com/GivailsonNeves/vacation-rental-api/domain"
	"github.com/labstack/echo/v4"
)

type (
	UserRepository interface {
		Create(user *domain.User) (*domain.User, error)
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
	return nil
}
