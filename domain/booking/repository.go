package booking

import (
	"fmt"

	"github.com/GivailsonNeves/vacation-rental-api/domain"
	"gorm.io/gorm"
)

type Filter struct {
	IDs  []int64 `json:"ids"`
	Name *string `json:"name"`
	Date *string `json:"date"`
}

type (
	Repository struct {
		db *gorm.DB
	}
)

func NewRepository(db *gorm.DB) *Repository {
	db.AutoMigrate(&domain.Booking{})
	return &Repository{
		db: db,
	}
}

func (e Repository) FindAll(paginationOptions *domain.PaginationInputType, filter *Filter) ([]domain.Booking, *domain.PaginationResultType, error) {
	bookings := []domain.Booking{}

	if paginationOptions == nil {
		paginationOptions = &domain.PaginationInputType{
			StartIndex:    0,
			First:         10,
			OrderBy:       "id",
			SortDirection: domain.SortDirectionEnumDesc,
		}
	}

	countQuery := e.db.Limit(paginationOptions.First)
	query := e.db.Limit(paginationOptions.First).
		Offset(paginationOptions.StartIndex).
		Order(fmt.Sprintf("%s %s", paginationOptions.OrderBy, paginationOptions.SortDirection))

	if filter != nil && filter.IDs != nil {
		countQuery = countQuery.Where("id IN ?", filter.IDs)
		query = query.Where("id IN ?", filter.IDs)
	}

	if filter != nil && filter.Name != nil {
		countQuery = countQuery.Where("name LIKE ?", fmt.Sprintf("%%%s%%", *filter.Name))
		query = query.Where("name LIKE ?", fmt.Sprintf("%%%s%%", *filter.Name))
	}

	var count int64

	countQuery.Find(&bookings).Count(&count)

	if count > 0 {
		if err := query.
			Find(&bookings).Error; err != nil {
			return nil, nil, err
		}
	}

	pageInfo := &domain.PaginationResultType{
		StartIndex:     0,
		EndIndex:       0,
		HasNextPage:    false,
		HasPreviusPage: false,
	}

	if len(bookings) > 0 {
		endIndex := paginationOptions.StartIndex + len(bookings)
		pageInfo.StartIndex = paginationOptions.StartIndex
		pageInfo.EndIndex = endIndex
		pageInfo.HasPreviusPage = paginationOptions.StartIndex > 0
		pageInfo.HasNextPage = count > int64(endIndex)
	}

	return bookings, pageInfo, nil
}

func (r Repository) Create(booking *domain.Booking) (*domain.Booking, error) {

	newBooking := &domain.Booking{
		Name:    booking.Name,
		StartAt: booking.StartAt,
		EndAt:   booking.EndAt,
	}
	result := r.db.Create(&newBooking)

	if result.Error != nil {
		return nil, result.Error
	}

	return newBooking, nil
}

func (r Repository) Update(booking *domain.Booking) (*domain.Booking, error) {
	updatedBooking := &domain.Booking{
		Name:    booking.Name,
		StartAt: booking.StartAt,
		EndAt:   booking.EndAt,
	}

	result := r.db.Model(&booking).Updates(&updatedBooking)

	if result.Error != nil {
		return nil, result.Error
	}

	return booking, nil
}

func (r Repository) Delete(booking *domain.Booking) (*domain.Booking, error) {
	result := r.db.Delete(&booking)

	if result.Error != nil {
		return nil, result.Error
	}

	return booking, nil
}

func (r Repository) Find(id uint64) (*domain.Booking, error) {
	bookings := []domain.Booking{}

	if err := r.db.Find(&bookings, "id = ? ", id).Error; err != nil {
		return nil, err
	}

	if len(bookings) == 0 {
		return nil, nil
	}

	return &bookings[0], nil
}
