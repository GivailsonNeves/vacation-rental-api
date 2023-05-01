package book

import (
	"database/sql/driver"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/GivailsonNeves/vacation-rental-api/storage"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func NewTestRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func TestCreate(t *testing.T) {

	book := &Book{
		Name:    "Holliday",
		StartAt: time.Now(),
		EndAt:   time.Now(),
	}

	t.Run("should create a book", func(t *testing.T) {
		mock, db := storage.GetFakeDB()
		repo := NewTestRepository(db)

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

func TestUpdate(t *testing.T) {
	book := &Book{
		ID:      1,
		Name:    "Christimas Holliday",
		StartAt: time.Now(),
		EndAt:   time.Now(),
	}

	t.Run("Should update a book name", func(t *testing.T) {
		mock, db := storage.GetFakeDB()
		repo := NewTestRepository(db)

		mock.ExpectBegin()
		mock.
			ExpectExec(regexp.QuoteMeta(`UPDATE "books" SET "name"=$1,"start_at"=$2,"end_at"=$3,"updated_at"=$4 WHERE "books"."deleted_at" IS NULL AND "id" = $5`)).
			WithArgs(book.Name, book.StartAt, book.EndAt, sqlmock.AnyArg(), book.ID).
			WillReturnResult(driver.RowsAffected(1))
		mock.ExpectCommit()
		updatedBook, err := repo.Update(book)
		assert.NoError(t, err)
		assert.Equal(t, book.ID, updatedBook.ID)
	})
}

func TestDelete(t *testing.T) {

	book := &Book{
		ID: 1,
	}
	t.Run("Should delete a given book", func(t *testing.T) {
		mock, db := storage.GetFakeDB()
		repo := NewTestRepository(db)

		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(`UPDATE "books" SET "deleted_at"=$1 WHERE "books"."id" = $2 AND "books"."deleted_at" IS NULL`)).
			WithArgs(sqlmock.AnyArg(), book.ID).
			WillReturnResult(driver.RowsAffected(1))
		mock.ExpectCommit()

		book, err := repo.Delete(book)
		assert.NoError(t, err)
		assert.NotNil(t, book.DeletedAt)
	})
}

func TestFind(t *testing.T) {

	book := &Book{
		ID:        1,
		Name:      "Christimas Holliday",
		CreatedAt: time.Now(),
		StartAt:   time.Now(),
		EndAt:     time.Now(),
	}

	t.Run("should find a book", func(t *testing.T) {

		mock, db := storage.GetFakeDB()
		repo := NewTestRepository(db)

		mock.
			ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "books" WHERE id = $1  AND "books"."deleted_at" IS NULL`)).
			WithArgs(book.ID).
			WillReturnRows(
				sqlmock.NewRows([]string{"id", "name", "start_at", "end_at", "created_at", "updated_at", "deleted_at"}).
					AddRow(book.ID, book.Name, book.StartAt, book.EndAt, book.CreatedAt, book.UpdatedAt, book.DeletedAt))
		recipe, err := repo.Find(uint64(book.ID))
		assert.NoError(t, err)
		assert.NotNil(t, recipe.Name)
	})
}
