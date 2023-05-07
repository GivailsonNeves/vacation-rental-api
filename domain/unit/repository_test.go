package unit

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

	unit := &domain.Unit{
		Avenue: "Quadra 02",
		Number: "26",
		Photo:  "http://photo.jpg",
		Type:   "casa",
	}

	t.Run("should create a unit", func(t *testing.T) {
		mock, db := storage.GetFakeDB()
		repo := NewTestRepository(db)

		mock.ExpectBegin()
		mock.
			ExpectQuery(regexp.QuoteMeta(`INSERT INTO "units" ("avenue","number","type","photo","created_at","updated_at","deleted_at") VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING "id"`)).
			WithArgs(unit.Avenue, unit.Number, unit.Type, unit.Photo, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()

		unitCreated, err := repo.Create(unit)
		assert.NoError(t, err)
		assert.Equal(t, unitCreated.ID, uint(1))
	})
}

func TestDelete(t *testing.T) {

	unit := &domain.Unit{
		ID: 1,
	}

	t.Run("should delete a unit", func(t *testing.T) {
		mock, db := storage.GetFakeDB()
		repo := NewTestRepository(db)

		mock.ExpectBegin()
		mock.
			ExpectExec(regexp.QuoteMeta(`UPDATE "units" SET "deleted_at"=$1 WHERE "units"."id" = $2 AND "units"."deleted_at" IS NULL`)).
			WithArgs(sqlmock.AnyArg(), unit.ID).
			WillReturnResult(driver.RowsAffected(1))

		mock.ExpectCommit()
		userDeleted, err := repo.Delete(unit)
		assert.NoError(t, err)
		assert.NotNil(t, userDeleted.DeletedAt)

	})
}

func TestUpdate(t *testing.T) {
	unit := &domain.Unit{
		ID:     1,
		Avenue: "Quadra 02",
		Number: "26",
		Photo:  "http://photo.jpg",
		Type:   "casa",
	}

	t.Run("Should update a unit", func(t *testing.T) {
		mock, db := storage.GetFakeDB()
		repo := NewTestRepository(db)

		mock.ExpectBegin()
		mock.
			ExpectExec(regexp.QuoteMeta(`UPDATE "units" SET "avenue"=$1,"number"=$2,"type"=$3,"photo"=$4,"updated_at"=$5 WHERE "units"."deleted_at" IS NULL AND "id" = $6`)).
			WithArgs(unit.Avenue, unit.Number, unit.Type, unit.Photo, sqlmock.AnyArg(), unit.ID).
			WillReturnResult(driver.RowsAffected(1))
		mock.ExpectCommit()
		updatedUnit, err := repo.Update(unit)
		assert.NoError(t, err)
		assert.Equal(t, unit.ID, updatedUnit.ID)
	})
}

func TestFind(t *testing.T) {

	unit := &domain.Unit{
		ID: 1,
	}

	t.Run("should find a user", func(t *testing.T) {

		mock, db := storage.GetFakeDB()
		repo := NewTestRepository(db)

		mock.
			ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "units" WHERE id = $1  AND "units"."deleted_at" IS NULL`)).
			WithArgs(unit.ID).
			WillReturnRows(
				sqlmock.NewRows([]string{"id", "avenue", "number", "type", "photo", "created_at", "updated_at", "deleted_at"}).
					AddRow(unit.ID, unit.Avenue, unit.Number, unit.Type, unit.Photo, unit.CreatedAt, unit.UpdatedAt, unit.DeletedAt))
		unitFound, err := repo.Find(uint64(unit.ID))
		assert.NoError(t, err)
		assert.NotNil(t, unitFound.Avenue)
	})
}
