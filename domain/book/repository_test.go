package book

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/GivailsonNeves/vacation-rental-api/storage"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {

	book := &Book{
		Name:    "Givailson",
		StartAt: time.Now(),
		EndAt:   time.Now(),
	}

	t.Run("should create a book", func(t *testing.T) {
		mock, db := storage.GetFakeDB()
		repo := NewRepository(db)

		mock.ExpectBegin()
		mock.
			ExpectQuery(regexp.QuoteMeta(`INSERT INTO "books" ("name","start_at","end_at","created_at","updated_at","deleted_at") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id`)).
			WithArgs(book.Name, book.StartAt, book.EndAt, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()

		book, err := repo.Create(book)
		assert.NoError(t, err)
		assert.Equal(t, book.ID, uint(1))
	})
}
