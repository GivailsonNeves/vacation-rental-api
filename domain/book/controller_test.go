package book

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/GivailsonNeves/vacation-rental-api/domain"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Create(book *Book) (*Book, error) {
	args := m.Called()
	return args[0].(*Book), args.Error(1)
}

func (m *MockRepository) FindAll(paginationOptions *domain.PaginationInputType, filter *Filter) ([]Book, *domain.PaginationResultType, error) {
	args := m.Called()
	return args[0].([]Book), &domain.PaginationResultType{}, args.Error(1)
}

func TestCreatBook(t *testing.T) {
	t.Run("should create a book", func(t *testing.T) {
		e := echo.New()
		mcPostBody := map[string]interface{}{
			"name":    "givailson",
			"startAt": time.Now(),
			"endAt":   time.Now(),
		}
		body, _ := json.Marshal(mcPostBody)
		req := httptest.NewRequest(http.MethodGet, "/books", bytes.NewReader(body))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockRepository := new(MockRepository)
		mockRepository.On("Create").Return(&Book{}, nil)

		controller := NewController(mockRepository)
		controller.Create(c)

		assert.Equal(t, http.StatusOK, rec.Code)
	})
	t.Run("should return bad status when body is empty a book", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/books", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockRepository := new(MockRepository)
		mockRepository.On("Create").Return(&Book{}, nil)

		controller := NewController(mockRepository)
		controller.Create(c)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
	t.Run("should return bad status when required paramn is missing", func(t *testing.T) {
		e := echo.New()

		mcPostBody := map[string]interface{}{
			"name":    "givailson",
			"startAt": time.Now(),
			// "endAt":   time.Now(), //missing
		}
		body, _ := json.Marshal(mcPostBody)
		req := httptest.NewRequest(http.MethodGet, "/books", bytes.NewReader(body))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockRepository := new(MockRepository)
		mockRepository.On("Create").Return(&Book{}, nil)

		controller := NewController(mockRepository)
		controller.Create(c)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}
