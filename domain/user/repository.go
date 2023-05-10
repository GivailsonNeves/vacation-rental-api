package user

import (
	"fmt"
	"net/http"

	"github.com/GivailsonNeves/vacation-rental-api/domain"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Filter struct {
	IDs  []int64 `json:"ids"`
	Name *string `json:"name"`
	Type *string `json:"type"`
}

type (
	Repository struct {
		db *gorm.DB
	}
)

func NewRepository(db *gorm.DB) *Repository {
	db.AutoMigrate(&domain.User{})
	return &Repository{
		db: db,
	}
}

func (r Controller) FindAll(c echo.Context) error {
	units, _, _ := r.repo.FindAll(nil, nil)
	return c.JSON(http.StatusOK, units)
}

func (r Repository) Create(user *domain.User) (*domain.User, error) {
	newUser := &domain.User{
		Name:     user.Name,
		Email:    user.Email,
		Photo:    user.Photo,
		Phone:    user.Phone,
		Password: "change@password",
		Type:     user.Type,
	}

	result := r.db.Create((&newUser))

	if result.Error != nil {
		return nil, result.Error
	}

	return newUser, nil
}

func (r Repository) Delete(user *domain.User) (*domain.User, error) {

	result := r.db.Delete(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (r Repository) Update(user *domain.User) (*domain.User, error) {
	updatedUser := &domain.User{
		Name:     user.Name,
		Email:    user.Email,
		Photo:    user.Photo,
		Phone:    user.Phone,
		Password: user.Password,
		Type:     user.Type,
	}

	result := r.db.Model(&user).Updates(&updatedUser)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (r Repository) Find(id uint64) (*domain.User, error) {
	users := []domain.User{}

	if err := r.db.Omit("Password").Find(&users, "id = ? ", id).Error; err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, nil
	}

	return &users[0], nil
}
func (r Repository) FindAllByIds(ids []uint64) ([]domain.User, error) {
	users := []domain.User{}

	if err := r.db.Omit("Password").Find(&users, "id IN ? ", ids).Error; err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, nil
	}

	return users, nil
}

func (e Repository) FindAll(paginationOptions *domain.PaginationInputType, filter *Filter) ([]domain.User, *domain.PaginationResultType, error) {
	users := []domain.User{}

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

	countQuery.Find(&users).Count(&count)

	if count > 0 {
		if err := query.
			Find(&users).Error; err != nil {
			return nil, nil, err
		}
	}

	pageInfo := &domain.PaginationResultType{
		StartIndex:     0,
		EndIndex:       0,
		HasNextPage:    false,
		HasPreviusPage: false,
	}

	if len(users) > 0 {
		endIndex := paginationOptions.StartIndex + len(users)
		pageInfo.StartIndex = paginationOptions.StartIndex
		pageInfo.EndIndex = endIndex
		pageInfo.HasPreviusPage = paginationOptions.StartIndex > 0
		pageInfo.HasNextPage = count > int64(endIndex)
	}

	return users, pageInfo, nil
}
