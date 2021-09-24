package formatter

import "invest/app/models"

type CrowdfundFormatter struct {}

type CreateCrowdfundReturn struct {
	ID uint `json:"id"`
	AdminID string `json:"admin_id"`
	Title string `json:"title"`
	Need int `json:"need"`
	Earn int `json:"earn"`
	Description string `json:"description"`
	IsActive bool `json:"is_active"`
}

func (f CrowdfundFormatter) FormatCreateCrowdfund (crowdfund models.Crowdfund) CreateCrowdfundReturn {
	formatter := CreateCrowdfundReturn{
		ID:          crowdfund.ID,
		AdminID:     crowdfund.AdminId,
		Title:       crowdfund.Title,
		Need:        crowdfund.Need,
		Earn:        crowdfund.Earn,
		Description: crowdfund.Description,
		IsActive:    crowdfund.IsActive,
	}
	return formatter
}

func (f CrowdfundFormatter) FormatFindAllCrowdfunding (list []models.Crowdfund) []CreateCrowdfundReturn {
	var finalList []CreateCrowdfundReturn
	for _, item := range list {
		t := f.FormatCreateCrowdfund(item)
		finalList = append(finalList, t)
	}
	return finalList
}
