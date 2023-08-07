package repository

import (
	"html"
	"strings"

	"github.com/AbdulrahmanDaud10/diveInputCalories/pkg/api"
	"golang.org/x/crypto/bcrypt"
)

// Save user details
func Save(user *api.User) error {
	err := db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

// Generate encrypted password
func BeforeSave(user *api.User) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	return nil
}

// Get all users
func GetUsers(User *[]api.User) (err error) {
	err = db.Find(User).Error
	if err != nil {
		return err
	}
	return nil
}

// Get user by username
func GetUserByUsername(username string) (api.User, error) {
	var user api.User
	err := db.Where("username=?", username).Find(&user).Error
	if err != nil {
		return api.User{}, err
	}
	return user, nil
}

// Validate user password
func ValidateUserPassword(password string) error {
	var user api.User
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

// Get user by id
func GetUserById(id uint) (api.User, error) {
	var user api.User
	err := db.Where("id=?", id).Find(&user).Error
	if err != nil {
		return api.User{}, err
	}
	return user, nil
}

// Get user by id
func GetUser(User *api.User, id int) error {
	err = db.Where("id = ?", id).First(User).Error
	if err != nil {
		return err
	}
	return nil
}

// Update user
func UpdateUser(User *api.User) error {
	err = db.Omit("password").Updates(User).Error
	if err != nil {
		return err
	}
	return nil
}
