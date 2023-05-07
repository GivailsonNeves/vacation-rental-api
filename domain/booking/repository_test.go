package booking

import (
	"database/sql/driver"
	"regexp"
	"testing"
	"time"

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

	booking := &domain.Booking{
		Name:    "Holliday",
		StartAt: time.Now(),
		EndAt:   time.Now(),
	}

	t.Run("should create a booking", func(t *testing.T) {
		mock, db := storage.GetFakeDB()
		repo := NewTestRepository(db)

		mock.ExpectBegin()
		mock.
			ExpectQuery(regexp.QuoteMeta(`INSERT INTO "bookings" ("name","start_at","end_at","created_at","updated_at","deleted_at") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id`)).
			WithArgs(booking.Name, booking.StartAt, booking.EndAt, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()

		booking, err := repo.Create(booking)
		assert.NoError(t, err)
		assert.Equal(t, booking.ID, uint(1))
	})
}

func TestUpdate(t *testing.T) {
	booking := &domain.Booking{
		ID:      1,
		Name:    "Christimas Holliday",
		StartAt: time.Now(),
		EndAt:   time.Now(),
	}

	t.Run("Should update a booking name", func(t *testing.T) {
		mock, db := storage.GetFakeDB()
		repo := NewTestRepository(db)

		mock.ExpectBegin()
		mock.
			ExpectExec(regexp.QuoteMeta(`UPDATE "bookings" SET "name"=$1,"start_at"=$2,"end_at"=$3,"updated_at"=$4 WHERE "bookings"."deleted_at" IS NULL AND "id" = $5`)).
			WithArgs(booking.Name, booking.StartAt, booking.EndAt, sqlmock.AnyArg(), booking.ID).
			WillReturnResult(driver.RowsAffected(1))
		mock.ExpectCommit()
		updatedBooking, err := repo.Update(booking)
		assert.NoError(t, err)
		assert.Equal(t, booking.ID, updatedBooking.ID)
	})
}

func TestDelete(t *testing.T) {

	booking := &domain.Booking{
		ID: 1,
	}
	t.Run("Should delete a given booking", func(t *testing.T) {
		mock, db := storage.GetFakeDB()
		repo := NewTestRepository(db)

		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(`UPDATE "bookings" SET "deleted_at"=$1 WHERE "bookings"."id" = $2 AND "bookings"."deleted_at" IS NULL`)).
			WithArgs(sqlmock.AnyArg(), booking.ID).
			WillReturnResult(driver.RowsAffected(1))
		mock.ExpectCommit()

		booking, err := repo.Delete(booking)
		assert.NoError(t, err)
		assert.NotNil(t, booking.DeletedAt)
	})
}

func TestFind(t *testing.T) {

	booking := &domain.Booking{
		ID:        1,
		Name:      "Christimas Holliday",
		CreatedAt: time.Now(),
		StartAt:   time.Now(),
		EndAt:     time.Now(),
	}

	t.Run("should find a booking", func(t *testing.T) {

		mock, db := storage.GetFakeDB()
		repo := NewTestRepository(db)

		mock.
			ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "bookings" WHERE id = $1  AND "bookings"."deleted_at" IS NULL`)).
			WithArgs(booking.ID).
			WillReturnRows(
				sqlmock.NewRows([]string{"id", "name", "start_at", "end_at", "created_at", "updated_at", "deleted_at"}).
					AddRow(booking.ID, booking.Name, booking.StartAt, booking.EndAt, booking.CreatedAt, booking.UpdatedAt, booking.DeletedAt))
		bookingFound, err := repo.Find(uint64(booking.ID))
		assert.NoError(t, err)
		assert.NotNil(t, bookingFound.Name)
	})
}
