package book

import (
	"fmt"
	"time"

	"github.com/GivailsonNeves/vacation-rental-api/domain"
	"gorm.io/gorm"
)

type Book struct {
	ID        uint           `gorm:"primaryKey"`
	Name      string         `gorm:"size:255" json:"name"`
	StartAt   time.Time      `json:"startAt"`
	EndAt     time.Time      `json:"endAt"`
	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

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
	db.AutoMigrate(&Book{})
	return &Repository{
		db: db,
	}
}

func (e Repository) FindAll(paginationOptions *domain.PaginationInputType, filter *Filter) ([]Book, *domain.PaginationResultType, error) {
	books := []Book{}

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

	countQuery.Find(&books).Count(&count)

	if count > 0 {
		if err := query.
			Find(&books).Error; err != nil {
			return nil, nil, err
		}
	}

	pageInfo := &domain.PaginationResultType{
		StartIndex:     0,
		EndIndex:       0,
		HasNextPage:    false,
		HasPreviusPage: false,
	}

	if len(books) > 0 {
		endIndex := paginationOptions.StartIndex + len(books)
		pageInfo.StartIndex = paginationOptions.StartIndex
		pageInfo.EndIndex = endIndex
		pageInfo.HasPreviusPage = paginationOptions.StartIndex > 0
		pageInfo.HasNextPage = count > int64(endIndex)
	}

	return books, pageInfo, nil
}

func (r Repository) Create(book *Book) (*Book, error) {

	newBook := &Book{
		Name:    book.Name,
		StartAt: book.StartAt,
		EndAt:   book.EndAt,
	}
	result := r.db.Create(&newBook)

	if result.Error != nil {
		return nil, result.Error
	}

	return newBook, nil
}

func (r Repository) Update(book *Book) (*Book, error) {
	updatedBook := &Book{
		Name:    book.Name,
		StartAt: book.StartAt,
		EndAt:   book.EndAt,
	}

	result := r.db.Model(&book).Updates(&updatedBook)

	if result.Error != nil {
		return nil, result.Error
	}

	return book, nil
}

func (r Repository) Delete(book *Book) (*Book, error) {
	result := r.db.Delete(&book)

	if result.Error != nil {
		return nil, result.Error
	}

	return book, nil
}

func (r Repository) Find(id uint64) (*Book, error) {
	books := []Book{}

	if err := r.db.Find(&books, "id = ? ", id).Error; err != nil {
		return nil, err
	}

	if len(books) == 0 {
		return nil, nil
	}

	return &books[0], nil
}
