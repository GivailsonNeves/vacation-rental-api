package user

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

	user := &domain.User{
		Name:     "Givailson",
		Phone:    "+5561984896181",
		Email:    "givailson@gmail.com",
		Photo:    "http://foto.jpg",
		Password: "change@password",
		Type:     "staff",
	}

	t.Run("should create a staff user", func(t *testing.T) {
		mock, db := storage.GetFakeDB()
		repo := NewTestRepository(db)

		mock.ExpectBegin()
		mock.
			ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users" ("name","phone","email","photo","password","type","created_at","updated_at","deleted_at") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING "id"`)).
			WithArgs(user.Name, user.Phone, user.Email, user.Photo, user.Password, user.Type, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()

		guestCreated, err := repo.Create(user)
		assert.NoError(t, err)
		assert.Equal(t, guestCreated.ID, uint(1))
	})
}

func TestDelete(t *testing.T) {

	user := &domain.User{
		ID: 1,
	}

	t.Run("should delete a user", func(t *testing.T) {
		mock, db := storage.GetFakeDB()
		repo := NewTestRepository(db)

		mock.ExpectBegin()
		mock.
			ExpectExec(regexp.QuoteMeta(`UPDATE "users" SET "deleted_at"=$1 WHERE "users"."id" = $2 AND "users"."deleted_at" IS NULL`)).
			WithArgs(sqlmock.AnyArg(), user.ID).
			WillReturnResult(driver.RowsAffected(1))

		mock.ExpectCommit()
		userDeleted, err := repo.Delete(user)
		assert.NoError(t, err)
		assert.NotNil(t, userDeleted.DeletedAt)

	})
}

func TestUpdate(t *testing.T) {
	user := &domain.User{
		ID:       1,
		Name:     "Givailson",
		Phone:    "+5561984896181",
		Email:    "givailson@gmail.com",
		Password: "change@password",
		Photo:    "http://photo.jpg",
		Type:     "staff",
	}

	t.Run("Should update a user", func(t *testing.T) {
		mock, db := storage.GetFakeDB()
		repo := NewTestRepository(db)

		mock.ExpectBegin()
		mock.
			ExpectExec(regexp.QuoteMeta(`UPDATE "users" SET "name"=$1,"phone"=$2,"email"=$3,"photo"=$4,"password"=$5,"type"=$6,"updated_at"=$7 WHERE "users"."deleted_at" IS NULL AND "id" = $8`)).
			WithArgs(user.Name, user.Phone, user.Email, user.Photo, user.Password, user.Type, sqlmock.AnyArg(), user.ID).
			WillReturnResult(driver.RowsAffected(1))
		mock.ExpectCommit()
		updatedUser, err := repo.Update(user)
		assert.NoError(t, err)
		assert.Equal(t, user.ID, updatedUser.ID)
	})
}

func TestFind(t *testing.T) {

	user := &domain.User{
		ID:       1,
		Name:     "Givailson",
		Phone:    "+5561984896181",
		Email:    "givailson@gmail.com",
		Password: "change@password",
		Photo:    "http://photo.jpg",
		Type:     "staff",
	}

	t.Run("should find a user", func(t *testing.T) {

		mock, db := storage.GetFakeDB()
		repo := NewTestRepository(db)

		mock.
			ExpectQuery(regexp.QuoteMeta(`SELECT "users"."id","users"."name","users"."phone","users"."email","users"."photo","users"."type","users"."created_at","users"."updated_at","users"."deleted_at" FROM "users" WHERE id = $1 AND "users"."deleted_at" IS NULL`)).
			WithArgs(user.ID).
			WillReturnRows(
				sqlmock.NewRows([]string{"id", "name", "phone", "email", "photo", "type", "created_at", "updated_at", "deleted_at"}).
					AddRow(user.ID, user.Name, user.Phone, user.Email, user.Photo, user.Type, user.CreatedAt, user.UpdatedAt, user.DeletedAt))
		guest, err := repo.Find(uint64(user.ID))
		assert.NoError(t, err)
		assert.NotNil(t, guest.Name)
	})
}
