package unit

import (
	"fmt"

	"github.com/GivailsonNeves/vacation-rental-api/domain"
	"github.com/GivailsonNeves/vacation-rental-api/domain/user"
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

	var userIds []uint64

	for _, owner := range unit.Owners {
		userIds = append(userIds, uint64(owner.ID))
	}

	usersRepo := user.NewRepository(r.db)

	users, error := usersRepo.FindAllByIds(userIds)

	if error != nil {
		return nil, error
	}

	if len(unit.Owners) != len(users) {
		return nil, fmt.Errorf("unexpected user id")
	}

	newUnit := &domain.Unit{
		Avenue: unit.Avenue,
		Number: unit.Number,
		Type:   unit.Type,
		Photo:  unit.Photo,
		Owners: users,
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

func (e Repository) FindAll(paginationOptions *domain.PaginationInputType, filter *Filter) ([]domain.Unit, *domain.PaginationResultType, error) {
	units := []domain.Unit{}

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

	var count int64

	countQuery.Find(&units).Count(&count)

	if count > 0 {
		if err := query.
			Find(&units).Error; err != nil {
			return nil, nil, err
		}
	}

	pageInfo := &domain.PaginationResultType{
		StartIndex:     0,
		EndIndex:       0,
		HasNextPage:    false,
		HasPreviusPage: false,
	}

	if len(units) > 0 {
		endIndex := paginationOptions.StartIndex + len(units)
		pageInfo.StartIndex = paginationOptions.StartIndex
		pageInfo.EndIndex = endIndex
		pageInfo.HasPreviusPage = paginationOptions.StartIndex > 0
		pageInfo.HasNextPage = count > int64(endIndex)
	}

	return units, pageInfo, nil
}
