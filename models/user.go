package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// User represents the user entity with its attributes
type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string `gorm:"unique;not null" json:"email"`
	Age       int
}

// GetUserByID retrieves a user by their ID from the database
func GetUserByID(db *gorm.DB, id uint) User {
	var user User
	db.First(&user, id)
	return user
}

// GetAllUsers retrieves all users from the database
func GetAllUsers(db *gorm.DB) []User {
	var users []User
	db.Find(&users)
	return users
}

// CreateUser creates a new user in the database
func CreateUser(db *gorm.DB, user *User) {
	db.Create(&user)
	fmt.Println("User Created Successfully:", user.FirstName)
}

// UpdateUser updates a user in the database by their ID
func UpdateUser(db *gorm.DB, id int, user *User) {
	result := db.Model(&User{}).Where("id = ?", id).Updates(user)
	if result.Error != nil {
		fmt.Println("Error updating user:", result.Error)
		return
	}
}

// DeleteUserByID deletes a user from the database by their ID
func DeleteUserByID(db *gorm.DB, id int) {
	err := db.Delete(&User{}, id).Error
	if err != nil {
		fmt.Println("Error deleting user:", err)
		return
	}
}
