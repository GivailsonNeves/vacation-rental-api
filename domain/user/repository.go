package user

import (
	"github.com/GivailsonNeves/vacation-rental-api/domain"
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

	if err := r.db.Find(&users, "id = ? ", id).Error; err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, nil
	}

	return &users[0], nil
}
