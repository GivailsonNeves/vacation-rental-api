package guest

import (
	"github.com/GivailsonNeves/vacation-rental-api/domain"
	"gorm.io/gorm"
)

type Filter struct {
	IDs       []int64 `json:"ids"`
	Name      *string `json:"name"`
	DocNumber *string `json:"docNumber"`
}

type (
	Repository struct {
		db *gorm.DB
	}
)

func NewRepository(db *gorm.DB) *Repository {
	db.AutoMigrate(&domain.Guest{})
	return &Repository{
		db: db,
	}
}

func (r Repository) Create(guest *domain.Guest) (*domain.Guest, error) {

	newGuest := &domain.Guest{
		Name:      guest.Name,
		DocNumber: guest.DocNumber,
		DocType:   guest.DocType,
		Photo:     guest.Photo,
		UnitID:    guest.UnitID,
		Email:     guest.Email,
		Phone:     guest.Phone,
	}

	result := r.db.Create((&newGuest))

	if result.Error != nil {
		return nil, result.Error
	}

	return newGuest, nil
}

func (r Repository) Delete(guest *domain.Guest) (*domain.Guest, error) {

	result := r.db.Delete(&guest)

	if result.Error != nil {
		return nil, result.Error
	}

	return guest, nil
}

func (r Repository) Update(guest *domain.Guest) (*domain.Guest, error) {
	updatedGuest := &domain.Guest{
		Name:      guest.Name,
		DocNumber: guest.DocNumber,
		DocType:   guest.DocType,
		Photo:     guest.Photo,
	}

	result := r.db.Model(&guest).Updates(&updatedGuest)

	if result.Error != nil {
		return nil, result.Error
	}

	return guest, nil
}

func (r Repository) Find(id uint64) (*domain.Guest, error) {
	guests := []domain.Guest{}

	if err := r.db.Find(&guests, "id = ? ", id).Error; err != nil {
		return nil, err
	}

	if len(guests) == 0 {
		return nil, nil
	}

	return &guests[0], nil
}
