package repository

import (
	"github.com/AbdulrahmanDaud10/diveInputCalories/pkg/api"
)

// Create a role
func CreateRole(Role *api.Role) error {
	err = db.Create(Role).Error
	if err != nil {
		return err
	}
	return nil
}

// Get all roles
func GetRoles(Role *[]api.Role) error {
	err = db.Find(Role).Error
	if err != nil {
		return err
	}
	return nil
}

// Get role by id
func GetRole(Role *api.Role, id int) error {
	err = db.Where("id = ?", id).First(Role).Error
	if err != nil {
		return err
	}
	return nil
}

// Update role
func UpdateRole(Role *api.Role) error {
	db.Save(Role)
	return nil
}
