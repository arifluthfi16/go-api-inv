package services

import (
	"invest/app/models"
	"invest/app/utils"
)

type TransparansiService struct {}

func (t *TransparansiService) Create (data utils.CreateTransparansiInput) (models.Transparansi, error){
	tr := models.Transparansi{}

	tr = models.Transparansi{
		AdminId:     data.AdminId,
		CrowdfundId: data.CrowdfundId,
		Title:       data.Title,
		ShortDesc:   data.ShortDesc,
		Content:     data.Content,
		Thumbnail:   data.Thumbnail,
		IsPublished: false,
	}

	if err := Manager.DB.Create(&tr).Error; err != nil {
		return tr, err
	}
	return tr, nil
}

func (t *TransparansiService) FindOneById (trId string) (models.Transparansi, error){
	var tr models.Transparansi
	result := Manager.DB.Take(&tr, trId)
	return tr, result.Error
}

func(t *TransparansiService) FindAll () ([]models.Transparansi, error){
	var tr []models.Transparansi
	result := Manager.DB.Find(&tr)
	return tr, result.Error
}

func (t *TransparansiService) DeleteById (trId string) (int64, error){
	var tr models.FAQ
	result := Manager.DB.Delete(&tr, trId)
	return result.RowsAffected, result.Error
}

func (t *TransparansiService) Update (tr *models.Transparansi, input utils.UpdateTransparansiInput) (models.Transparansi, error){
	tr.Title = input.Title
	tr.ShortDesc = input.ShortDesc
	tr.Content = input.Content
	tr.Thumbnail = input.Thumbnail
	tr.IsPublished = input.IsPublished

	if err := Manager.DB.Save(&tr).Error; err != nil {
		return *tr, err
	}
	return *tr, nil
}