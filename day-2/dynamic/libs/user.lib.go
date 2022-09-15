package libs

import (
	"dynamic-crud/config"
	"dynamic-crud/models"
)

func CheckExists(id int) (bool, error) {
	var count int64
	err := config.DB.Table("users").Where("id = ?", id).Count(&count)

	return count > 0, err.Error
}

func CreateUser(user *models.User) error {
	return config.DB.Table("users").Create(user).Error
}

func GetUser(id int) (models.User, error) {
	var user models.User
	err := config.DB.Table("users").Where("id = ?", id).First(&user)
	return user, err.Error
}

func GetUsers() ([]models.User, error) {
	var users []models.User
	err := config.DB.Table("users").Find(&users)
	return users, err.Error
}

func UpdateUser(id int, newUser *models.User) (models.User, error) {
	var user models.User
	err := config.DB.Table("users").Where("id = ?", id).First(&user)

	user.FirstName = newUser.FirstName
	user.LastName = newUser.LastName
	config.DB.Table("users").Save(&user)

	return user, err.Error
}

func DeleteUser(id int) error {
	var user models.User
	err := config.DB.Table("users").Where("id = ?", id).First(&user)
	config.DB.Table("users").Delete(&user)

	return err.Error
}
