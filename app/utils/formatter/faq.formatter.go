package formatter

import "invest/app/models"

type FAQFormatter struct {}

type CreateFAQReturn struct {
	ID uint `json:"id"`
	AdminID string `json:"admin_id"`
	Pertanyaan string `json:"pertanyaan"`
	Jawaban string `json:"jawaban"`
	Kategori string `json:"kategori"`
	Count int `json:"count"`
}

func (f FAQFormatter) FormatCreateFAQ (faq models.FAQ) CreateFAQReturn {
	fr := CreateFAQReturn{
		ID:         faq.ID,
		AdminID:    faq.AdminId,
		Pertanyaan: faq.Pertanyaan,
		Jawaban:    faq.Jawaban,
		Kategori:   faq.Jawaban,
		Count:      faq.Count,
	}
	return fr
}

func (f FAQFormatter) FormatFindAllFAQ (list []models.FAQ) []CreateFAQReturn {
	var finalList []CreateFAQReturn
	for _, item := range list {
		t := f.FormatCreateFAQ(item)
		finalList = append(finalList, t)
	}
	return finalList
}