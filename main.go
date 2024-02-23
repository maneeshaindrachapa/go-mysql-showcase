package main

import (
	"fmt"

	"github.com/maneeshaindrachapa/go-mysql/configs"
	"github.com/maneeshaindrachapa/go-mysql/database"
	"github.com/maneeshaindrachapa/go-mysql/models"
)

func main() {
	// Initialize environment configurations
	configs.InitEnvConfigs()

	// Set up database configuration using environment variables
	databaseConfig := database.DBConfig{
		DatabaseUser:     configs.EnvConfigs.DatabaseUser,
		DatabasePassword: configs.EnvConfigs.DatabasePassword,
		DatabaseName:     configs.EnvConfigs.DatabaseName,
		DatabaseUrl:      configs.EnvConfigs.DatabaseUrl,
		DatabasePort:     configs.EnvConfigs.DatabasePort,
	}

	// Establish a connection to the MySQL database
	db, _ := database.NewConnectionToMySQL(&databaseConfig)
	// Initialize the database schema
	database.InitializeDatabase(db)

	// Defer closing the database connection until the function exits
	defer database.CloseDB(db)

	// Trying out queries
	// Create new Users
	sampleUser01 := models.User{FirstName: "John", LastName: "Doe", Email: "john.doe@example.com", Age: 20}
	sampleUser02 := models.User{FirstName: "Johnny", LastName: "Depp", Email: "johnny.depp@example.com", Age: 40}
	models.CreateUser(db, &sampleUser01)
	models.CreateUser(db, &sampleUser02)

	// Get All Users
	users := models.GetAllUsers(db)
	for _, user := range users {
		fmt.Println("User FirstName:", user.FirstName)
	}

	// Get Specific User
	user := models.GetUserByID(db, 1)
	fmt.Printf("Queried User %d, %s\n", user.ID, user.FirstName)

	// Update User
	sampleUser03 := models.User{FirstName: "Ally", LastName: "Green", Email: "allygreee@example.com", Age: 27}
	models.UpdateUser(db, 1, &sampleUser03)
	updatedUser := models.GetUserByID(db, 1)
	fmt.Printf("Updated User FirstName %s\n", updatedUser.FirstName)

	// Delete User
	models.DeleteUserByID(db, 1)
	deletedUser := models.GetUserByID(db, 1)
	fmt.Println("Deleted User", deletedUser.FirstName)
}
