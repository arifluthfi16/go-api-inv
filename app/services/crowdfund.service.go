package services

import (
	"fmt"
	"gorm.io/gorm"
	"invest/app/models"
	"invest/app/utils"
)

type CrowdfundService struct {}

func (c *CrowdfundService) CreateCrowdfund (data utils.CreateCrowdfundInput) (models.Crowdfund, error){
	cr := models.Crowdfund{}

	cr = models.Crowdfund{
		Model:       gorm.Model{},
		AdminId:     data.AdminID,
		Title:       data.Title,
		Need:        data.Need,
		Earn:        0,
		Description: data.Description,
		IsActive:    false,
	}

	if err := Manager.DB.Create(&cr).Error; err != nil {
		return cr, err
	}
	return cr, nil
}

func (c *CrowdfundService) FindOneById (crId string) (models.Crowdfund, error){
	var cr models.Crowdfund
	result := Manager.DB.Take(&cr, crId)
	return cr, result.Error
}

func (c *CrowdfundService) FindAll () ([]models.Crowdfund, error){
	var cr []models.Crowdfund
	result := Manager.DB.Find(&cr)
	return cr, result.Error
}

func (c *CrowdfundService) DeleteById (crId string) (int64, error){
	var user models.Crowdfund
	result := Manager.DB.Delete(&user, crId)
	return result.RowsAffected, result.Error
}

func (c *CrowdfundService) AddEarn (crId string, amount int) (models.Crowdfund, error) {
	fmt.Println(crId, amount)
	var cr models.Crowdfund
	Manager.DB.Take(&cr, crId)
	cr.Earn += amount
	if err := Manager.DB.Save(&cr).Error; err != nil {
		return cr, err
	}
	return cr, nil
}

func (c *CrowdfundService) Update (cr *models.Crowdfund, input utils.UpdateCrowdfundInput) (models.Crowdfund, error){
	cr.Description = input.Description
	cr.Need = input.Need
	cr.Title = input.Title
	cr.Earn = input.Earn
	cr.IsActive = input.IsActive

	if err := Manager.DB.Save(&cr).Error; err != nil {
		return *cr, err
	}
	return *cr, nil
}