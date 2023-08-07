package main

import (
	"log"
	"os"

	"github.com/AbdulrahmanDaud10/diveInputCalories/pkg/api"
	"github.com/AbdulrahmanDaud10/diveInputCalories/pkg/app"
	"github.com/AbdulrahmanDaud10/diveInputCalories/pkg/repository"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func LoadEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Println(".env file loaded successfully")
}

func main() {
	LoadEnv()

	app.ServeApplicationRoute()
}

// run database migrations and add seed data
func LoadDatabase() {
	repository.InitDb()
	db.AutoMigrate(&api.Role{})
	db.AutoMigrate(&api.User{})
	// Db.AutoMigrate(&api.Room{})
	// Db.AutoMigrate(&api.Booking{})
	seedData()
}

// load seed data into the database
func seedData() {
	var roles = []api.Role{{Name: "admin", Description: "Administrator role"}, {Name: "customer", Description: "Authenticated customer role"}, {Name: "visitor", Description: "Unauthenticated customer role"}}
	var user = []api.User{{Username: os.Getenv("ADMIN_USERNAME"), Email: os.Getenv("ADMIN_EMAIL"), Password: os.Getenv("ADMIN_PASSWORD"), RoleID: 1}}
	db.Save(&roles)
	db.Save(&user)
}
