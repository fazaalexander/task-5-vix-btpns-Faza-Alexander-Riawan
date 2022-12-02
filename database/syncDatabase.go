package database

import "final-task/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
