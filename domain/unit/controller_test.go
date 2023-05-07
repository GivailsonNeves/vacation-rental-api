package unit

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

func (m *MockRepository) Create(unit *domain.Unit) (*domain.Unit, error) {
	args := m.Called()
	return args[0].(*domain.Unit), args.Error(1)
}

func (m *MockRepository) Delete(unit *domain.Unit) (*domain.Unit, error) {
	args := m.Called()
	return args[0].(*domain.Unit), args.Error(1)
}

// func (m *MockRepository) FindAll(paginationOptions *domain.PaginationInputType, filter *Filter) ([]domain.Booking, *domain.PaginationResultType, error) {
// 	args := m.Called()
// 	return args[0].([]domain.Booking), &domain.PaginationResultType{}, args.Error(1)
// }

// func (m *MockRepository) Delete(booking *domain.Booking) (*domain.Booking, error) {
// 	args := m.Called()
// 	return args[0].(*domain.Booking), args.Error(1)
// }
// func (m *MockRepository) Update(booking *domain.Booking) (*domain.Booking, error) {
// 	args := m.Called()
// 	return args[0].(*domain.Booking), args.Error(1)
// }
// func (m *MockRepository) Find(id uint64) (*domain.Booking, error) {
// 	args := m.Called()
// 	return args[0].(*domain.Booking), args.Error(1)
// }

func TestCreatUnit(t *testing.T) {
	t.Run("should create a unit", func(t *testing.T) {
		e := echo.New()
		mcPostBody := map[string]interface{}{
			"avenue": "quadra 02",
			"number": "26",
			"type":   "casa da pscina",
			"photo":  "http://foto.jpg",
		}
		body, _ := json.Marshal(mcPostBody)
		req := httptest.NewRequest(http.MethodPost, "/units", bytes.NewReader(body))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockRepository := new(MockRepository)
		mockRepository.On("Create").Return(&domain.Unit{}, nil)

		controller := NewController(mockRepository)
		controller.Create(c)

		assert.Equal(t, http.StatusOK, rec.Code)
	})
	// t.Run("should return bad status when body is empty a booking", func(t *testing.T) {
	// 	e := echo.New()

	// 	req := httptest.NewRequest(http.MethodGet, "/bookings", nil)
	// 	rec := httptest.NewRecorder()
	// 	c := e.NewContext(req, rec)

	// 	mockRepository := new(MockRepository)
	// 	mockRepository.On("Create").Return(&domain.Booking{}, nil)

	// 	controller := NewController(mockRepository)
	// 	controller.Create(c)

	// 	assert.Equal(t, http.StatusBadRequest, rec.Code)
	// })
	// t.Run("should return bad status when required paramn is missing", func(t *testing.T) {
	// 	e := echo.New()

	// 	mcPostBody := map[string]interface{}{
	// 		"name":    "givailson",
	// 		"startAt": time.Now(),
	// 		// "endAt":   time.Now(), //missing
	// 	}
	// 	body, _ := json.Marshal(mcPostBody)
	// 	req := httptest.NewRequest(http.MethodGet, "/bookings", bytes.NewReader(body))
	// 	rec := httptest.NewRecorder()
	// 	c := e.NewContext(req, rec)

	// 	mockRepository := new(MockRepository)
	// 	mockRepository.On("Create").Return(&domain.Booking{}, nil)

	// 	controller := NewController(mockRepository)
	// 	controller.Create(c)

	// 	assert.Equal(t, http.StatusBadRequest, rec.Code)
	// })
}

func TestDeleteUnit(t *testing.T) {
	t.Run("should delete a unit", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodDelete, "/units/:id", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("2")

		mockRepository := new(MockRepository)
		mockRepository.On("Delete").Return(&domain.Unit{}, nil)

		controller := NewController(mockRepository)
		controller.Delete(c)

		assert.Equal(t, http.StatusOK, rec.Code)
	})
}

// func TestUpdateBooking(t *testing.T) {
// 	t.Run("should update a booking", func(t *testing.T) {
// 		e := echo.New()

// 		mcPostBody := map[string]interface{}{
// 			"name":    "givailson",
// 			"startAt": time.Now(),
// 			"endAt":   time.Now(),
// 		}
// 		body, _ := json.Marshal(mcPostBody)

// 		req := httptest.NewRequest(http.MethodPost, "/bookings/:id", bytes.NewReader(body))
// 		rec := httptest.NewRecorder()
// 		c := e.NewContext(req, rec)
// 		c.SetParamNames("id")
// 		c.SetParamValues("2")

// 		mockRepository := new(MockRepository)
// 		mockRepository.On("Update").Return(&domain.Booking{}, nil)

// 		controller := NewController(mockRepository)
// 		controller.Update(c)

// 		assert.Equal(t, http.StatusOK, rec.Code)
// 	})
// }

// func TestFindBooking(t *testing.T) {
// 	t.Run("should find a booking", func(t *testing.T) {
// 		e := echo.New()

// 		req := httptest.NewRequest(http.MethodGet, "/bookings/:id", nil)
// 		rec := httptest.NewRecorder()
// 		c := e.NewContext(req, rec)
// 		c.SetParamNames("id")
// 		c.SetParamValues("2")

// 		mockRepository := new(MockRepository)
// 		mockRepository.On("Find").Return(&domain.Booking{}, nil)

// 		controller := NewController(mockRepository)
// 		controller.Find(c)

// 		assert.Equal(t, http.StatusOK, rec.Code)
// 	})
// }

// func TestFindAllBooking(t *testing.T) {
// 	t.Run("should find all bookings", func(t *testing.T) {
// 		e := echo.New()

// 		req := httptest.NewRequest(http.MethodGet, "/bookings", nil)
// 		rec := httptest.NewRecorder()
// 		c := e.NewContext(req, rec)

// 		mockRepository := new(MockRepository)
// 		mockRepository.On("FindAll").Return([]domain.Booking{}, nil)

// 		controller := NewController(mockRepository)
// 		controller.FindAll(c)

// 		assert.Equal(t, http.StatusOK, rec.Code)
// 	})
// }
