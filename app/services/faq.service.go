package services

import (
	"invest/app/models"
	"invest/app/utils"
)

type FAQService struct {}

func (f *FAQService) Create (data utils.CreateFAQInput) (models.FAQ, error){
	faq := models.FAQ{}

	faq = models.FAQ{
		AdminId:    data.AdminID,
		Pertanyaan: data.Pertanyaan,
		Jawaban:    data.Jawaban,
		Count:      0,
		Kategori:   data.Kategori,
	}

	if err := Manager.DB.Create(&faq).Error; err != nil {
		return faq, err
	}
	return faq, nil
}

func (f *FAQService) FindOneById (faqId string) (models.FAQ, error){
	var faq models.FAQ
	result := Manager.DB.Take(&faq, faqId)
	return faq, result.Error
}

func(f *FAQService) FindAll () ([]models.FAQ, error){
	var faqs []models.FAQ
	result := Manager.DB.Find(&faqs)
	return faqs, result.Error
}

func (f *FAQService) DeleteById (faqId string) (int64, error){
	var faq models.FAQ
	result := Manager.DB.Delete(&faq, faqId)
	return result.RowsAffected, result.Error
}

func (f *FAQService) Update (faq *models.FAQ, input utils.UpdateFAQInput) (models.FAQ, error){
	faq.Kategori = input.Kategori
	faq.Jawaban = input.Jawaban
	faq.Pertanyaan = input.Pertanyaan
	faq.Count = input.Count

	if err := Manager.DB.Save(&faq).Error; err != nil {
		return *faq, err
	}
	return *faq, nil
}