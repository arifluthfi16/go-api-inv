package formatter

import "invest/app/models"

type TransparansiFormatter struct {

}

type CreateTransparansiReturn struct {
	ID uint `json:"id"`
	AdminId string `json:"admin_id"`
	CrowdfundId string `json:"crowdfund_id"`
	Title string `json:"title"`
	ShortDesc string `json:"short_desc"`
	Content	string `json:"content"`
	Thumbnail string `json:"thumbnail"`
	IsPublished bool `json:"is_published"`
}

func (f TransparansiFormatter) FormatCreateTransparansi (tr models.Transparansi) CreateTransparansiReturn {
	fr := CreateTransparansiReturn{
		ID:          tr.ID,
		AdminId:     tr.AdminId,
		CrowdfundId: tr.CrowdfundId,
		Title:       tr.Title,
		ShortDesc:   tr.ShortDesc,
		Content:     tr.Content,
		Thumbnail:   tr.Thumbnail,
		IsPublished: tr.IsPublished,
	}
	return fr
}

func (f TransparansiFormatter) FormatFindAllTransparansi (list []models.Transparansi) []CreateTransparansiReturn {
	var finalList []CreateTransparansiReturn
	for _, item := range list {
		t := f.FormatCreateTransparansi(item)
		finalList = append(finalList, t)
	}
	return finalList
}