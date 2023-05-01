package storage

import (
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	config "github.com/GivailsonNeves/vacation-rental-api/config"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDB(params ...string) *gorm.DB {
	var err error
	conString := config.GetPostgresConnectionString()
	log.Print(conString)

	DB, err = gorm.Open(postgres.Open(config.GetPostgresConnectionString()), &gorm.Config{})

	if err != nil {
		log.Panic(err)
	}

	return DB
}

func GetDBInstance() *gorm.DB {
	return DB
}

func GetFakeDB() (sqlmock.Sqlmock, *gorm.DB) {

	conn, mock, _ := sqlmock.New()

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 conn,
		PreferSimpleProtocol: true,
	})

	DB, _ := gorm.Open(dialector, &gorm.Config{})

	return mock, DB
}
