package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GivailsonNeves/vacation-rental-api/domain"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Create(user *domain.User) (*domain.User, error) {
	args := m.Called()
	return args[0].(*domain.User), args.Error(1)
}

func (m *MockRepository) FindAll(paginationOptions *domain.PaginationInputType, filter *Filter) ([]domain.User, *domain.PaginationResultType, error) {
	args := m.Called()
	return args[0].([]domain.User), &domain.PaginationResultType{}, args.Error(1)
}

func TestCreatUser(t *testing.T) {
	t.Run("should create a owner user", func(t *testing.T) {
		e := echo.New()
		mcPostBody := map[string]interface{}{
			"name":     "givailson",
			"email":    "givailson@gmail.com",
			"phone":    "+5561984896181",
			"photo":    "http://foto.jpg",
			"password": "123456",
			"type":     "owner",
		}
		body, _ := json.Marshal(mcPostBody)
		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(body))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockRepository := new(MockRepository)
		mockRepository.On("Create").Return(&domain.User{}, nil)

		controller := NewController(mockRepository)
		controller.Create(c)

		assert.Equal(t, http.StatusOK, rec.Code)
	})
}
