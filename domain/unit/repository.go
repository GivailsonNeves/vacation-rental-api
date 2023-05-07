package unit

import (
	"github.com/GivailsonNeves/vacation-rental-api/domain"
	"gorm.io/gorm"
)

type Filter struct {
	IDs  []int64 `json:"ids"`
	Type *string `json:"type"`
}

type (
	Repository struct {
		db *gorm.DB
	}
)

func NewRepository(db *gorm.DB) *Repository {
	db.AutoMigrate(&domain.Unit{})
	return &Repository{
		db: db,
	}
}

func (r Repository) Create(unit *domain.Unit) (*domain.Unit, error) {
	newUnit := &domain.Unit{
		Avenue: unit.Avenue,
		Number: unit.Number,
		Type:   unit.Type,
		Photo:  unit.Photo,
	}

	result := r.db.Create((&newUnit))

	if result.Error != nil {
		return nil, result.Error
	}

	return newUnit, nil
}

func (r Repository) Delete(unit *domain.Unit) (*domain.Unit, error) {

	result := r.db.Delete(&unit)

	if result.Error != nil {
		return nil, result.Error
	}

	return unit, nil
}

func (r Repository) Update(unit *domain.Unit) (*domain.Unit, error) {
	updatedUnit := &domain.Unit{
		Avenue: unit.Avenue,
		Number: unit.Number,
		Type:   unit.Type,
		Photo:  unit.Photo,
	}

	result := r.db.Model(&unit).Updates(&updatedUnit)

	if result.Error != nil {
		return nil, result.Error
	}

	return unit, nil
}

func (r Repository) Find(id uint64) (*domain.Unit, error) {
	units := []domain.Unit{}

	if err := r.db.Find(&units, "id = ? ", id).Error; err != nil {
		return nil, err
	}

	if len(units) == 0 {
		return nil, nil
	}

	return &units[0], nil
}
