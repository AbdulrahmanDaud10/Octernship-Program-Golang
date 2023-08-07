package api

import "gorm.io/gorm"

// User model
type User struct {
	gorm.Model
	ID       uint   `gorm:"primary_key"`
	RoleID   uint   `gorm:"not null;DEFAULT:3" json:"role_id"`
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Email    string `gorm:"size:255;not null;unique" json:"email"`
	Password string `gorm:"size:255;not null" json:"-"`
	Role     Role   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

// Role model
type Role struct {
	gorm.Model
	ID          uint   `gorm:"primary_key"`
	Name        string `gorm:"size:50;not null;unique" json:"name"`
	Description string `gorm:"size:255;not null" json:"description"`
}

// Register model
type Register struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login model
type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Update model
type Update struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	RoleID   uint   `gorm:"not null" json:"role_id"`
}
