package services

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"invest/app/models"
	"invest/app/utils"
)

type UserService struct {}

func (u *UserService) CreateUser (userData utils.CreateUserInput) (models.User, error){
	user := models.User{}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user = models.User{
		Model:         gorm.Model{},
		DisplayPict:   "",
		Name:          userData.Name,
		Email:         userData.Email,
		Phonenumber:   userData.Phonenumber,
		IdentityImage: "",
		Password:      string(hashedPassword),
		IsVerified:    false,
		GoogleId:      "",
	}

	if err := Manager.DB.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (u *UserService) FindOneById (userId string) (models.User, error){
	var user models.User
	result := Manager.DB.Take(&user, userId)
	return user, result.Error
}

func (u *UserService) FindAll () ([]models.User, error){
	var users []models.User
	result := Manager.DB.Find(&users)
	return users, result.Error
}

func (u *UserService) FindOneByEmail (email string) (models.User, error){
	var user models.User
	result := Manager.DB.Where(&models.User{Email: email}).Take(&user)
	return user, result.Error
}

func (u *UserService) DeleteOneById (userId string) (int64, error){
	var user models.User
	result := Manager.DB.Unscoped().Delete(&user, userId)
	return result.RowsAffected, result.Error
}

func (u *UserService) UpdateUser (user *models.User, input utils.CreateUserInput) (models.User, error){
	user.Email = input.Email
	user.Phonenumber = input.Phonenumber
	user.Name = input.Name

	if err := Manager.DB.Save(&user).Error; err != nil {
		return *user, err
	}
	return *user, nil
}



