package guest

import (
	"database/sql/driver"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/GivailsonNeves/vacation-rental-api/domain"
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

	guest := &domain.Guest{
		Name:      "Holliday",
		DocNumber: "123456",
		UnitID:    uint(1),
		DocType:   "RG",
		Photo:     "asdfas",
	}

	t.Run("should create a guest", func(t *testing.T) {
		mock, db := storage.GetFakeDB()
		repo := NewTestRepository(db)

		mock.ExpectBegin()
		mock.
			ExpectQuery(regexp.QuoteMeta(`INSERT INTO "guests" ("name","doc_number","doc_type","photo","unit_id","phone","email","created_at","updated_at","deleted_at") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING "id"`)).
			WithArgs(guest.Name, guest.DocNumber, guest.DocType, guest.Photo, guest.UnitID, guest.Phone, guest.Email, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()

		guestCreated, err := repo.Create(guest)
		assert.NoError(t, err)
		assert.Equal(t, guestCreated.ID, uint(1))
	})
}

func TestDelete(t *testing.T) {

	guest := &domain.Guest{
		ID: 1,
	}

	t.Run("should delete a guest", func(t *testing.T) {
		mock, db := storage.GetFakeDB()
		repo := NewTestRepository(db)

		mock.ExpectBegin()
		mock.
			ExpectExec(regexp.QuoteMeta(`UPDATE "guests" SET "deleted_at"=$1 WHERE "guests"."id" = $2 AND "guests"."deleted_at" IS NULL`)).
			WithArgs(sqlmock.AnyArg(), guest.ID).
			WillReturnResult(driver.RowsAffected(1))

		mock.ExpectCommit()
		guestDeleted, err := repo.Delete(guest)
		assert.NoError(t, err)
		assert.NotNil(t, guestDeleted.DeletedAt)

	})
}

func TestUpdate(t *testing.T) {
	guest := &domain.Guest{
		ID:        1,
		Name:      "Holliday",
		DocNumber: "123456",
		DocType:   "RG",
		Photo:     "asdfas",
	}

	t.Run("Should update a guest", func(t *testing.T) {
		mock, db := storage.GetFakeDB()
		repo := NewTestRepository(db)

		mock.ExpectBegin()
		mock.
			ExpectExec(regexp.QuoteMeta(`UPDATE "guests" SET "name"=$1,"doc_number"=$2,"doc_type"=$3,"photo"=$4,"updated_at"=$5 WHERE "guests"."deleted_at" IS NULL AND "id" = $6`)).
			WithArgs(guest.Name, guest.DocNumber, guest.DocType, guest.Photo, sqlmock.AnyArg(), guest.ID).
			WillReturnResult(driver.RowsAffected(1))
		mock.ExpectCommit()
		updatedGuest, err := repo.Update(guest)
		assert.NoError(t, err)
		assert.Equal(t, guest.ID, updatedGuest.ID)
	})
}

func TestFind(t *testing.T) {

	guest := &domain.Guest{
		ID:        1,
		Name:      "Holliday",
		DocNumber: "123456",
		DocType:   "RG",
		Photo:     "asdfas",
	}

	t.Run("should find a guest", func(t *testing.T) {

		mock, db := storage.GetFakeDB()
		repo := NewTestRepository(db)

		mock.
			ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "guests" WHERE id = $1  AND "guests"."deleted_at" IS NULL`)).
			WithArgs(guest.ID).
			WillReturnRows(
				sqlmock.NewRows([]string{"id", "name", "doc_number", "doc_type", "photo", "created_at", "updated_at", "deleted_at"}).
					AddRow(guest.ID, guest.Name, guest.DocNumber, guest.DocType, guest.Photo, guest.CreatedAt, guest.UpdatedAt, guest.DeletedAt))
		guest, err := repo.Find(uint64(guest.ID))
		assert.NoError(t, err)
		assert.NotNil(t, guest.Name)
	})
}
