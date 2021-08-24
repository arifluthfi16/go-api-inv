package services

import (
	"invest/app/models"
	"invest/app/utils"
)

type InvestorService struct {}

func (i *InvestorService) Create (data utils.CreateInvestorInput) (models.Investor, error){
	inv := models.Investor{}

	inv = models.Investor{
		UserId:     	data.UserId,
		CrowdfundId: 	data.CrowdfundId,
		Fund: 			data.Fund,
	}

	if err := Manager.DB.Create(&inv).Error; err != nil {
		return inv, err
	}
	return inv, nil
}

func (i *InvestorService) FindOneById (invId string) (models.Investor, error){
	var inv models.Investor
	result := Manager.DB.Take(&inv, invId)
	return inv, result.Error
}

func (i *InvestorService) FindAll () ([]models.Investor, error){
	var inv []models.Investor
	result := Manager.DB.Find(&inv)
	return inv, result.Error
}

func (i *InvestorService) FindAllByCrowdfundId (crId string) ([]models.Investor, error){
	var inv []models.Investor
	result := Manager.DB.Find(&models.Investor{CrowdfundId: crId})
	return inv, result.Error
}

func (i *InvestorService) DeleteById (invId string) (int64, error){
	var inv models.Investor
	result := Manager.DB.Delete(&inv, invId)
	return result.RowsAffected, result.Error
}

func (i *InvestorService) Update (inv *models.Investor, input utils.UpdateInvestorInput) (models.Investor, error){
	inv.Fund = input.Fund

	if err := Manager.DB.Save(&inv).Error; err != nil {
		return *inv, err
	}
	return *inv, nil
}