package repository

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func InitDb() (*gorm.DB, error) {
	db, err = SetUpDatabaseConnection()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func SetUpDatabaseConnection() (*gorm.DB, error) {

	var (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		username = os.Getenv("DB_USER")
		database = os.Getenv("DB_NAME")
		password = os.Getenv("DB_PASSWORD")
	)
	dsn := fmt.Sprintf("host=%s port=%s user=%s database=%s  password=%s sslmode=disable",
		host,
		port,
		username,
		database,
		password,
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("db err: ", err)
		return nil, err
	}

	fmt.Println("Database connection successful...")

	return db, nil
}
