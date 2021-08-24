package services

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"invest/app/models"
	"invest/app/utils"
)

type AdminService struct {}

func (a *AdminService) CreateAdmin (adminData utils.CreateAdminInput) (models.Admin, error){
	admin := models.Admin{}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(adminData.Password), bcrypt.MinCost)
	if err != nil {
		return admin, err
	}

	admin = models.Admin{
		Model:    gorm.Model{},
		Name:     adminData.Name,
		Username: adminData.Username,
		Password: string(hashedPassword),
		Email:    adminData.Email,
		HakAkses: "admin",
	}

	if err := Manager.DB.Create(&admin).Error; err != nil {
		return admin, err
	}
	return admin, nil
}

func (a *AdminService) FindOneById (adminId string) (models.Admin, error){
	var admin models.Admin
	result := Manager.DB.Take(&admin, adminId)
	return admin, result.Error
}

func (a *AdminService) FindOneByUsername (username string)(models.Admin, error){
	var admin models.Admin
	result := Manager.DB.Where(&models.Admin{Username: username}).Take(&admin)
	return admin, result.Error
}

func (a *AdminService) FindAll () ([]models.Admin, error){
	var admin []models.Admin
	result := Manager.DB.Find(&admin)
	return admin, result.Error
}

func (a *AdminService) DeleteOneById (adminId string) (int64, error){
	var admin models.Admin
	result := Manager.DB.Delete(&admin, adminId)
	return result.RowsAffected, result.Error
}



