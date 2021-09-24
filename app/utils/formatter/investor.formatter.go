package formatter

import "invest/app/models"

type InvestorFormatter struct {}

type CreateInvestorReturn struct {
	ID uint `json:"id"`
	UserId string `json:"user_id"`
	CrowdfundId string `json:"crowdfund_id"`
	Fund int `json:"fund"`
}

func (f InvestorFormatter) FormatCreateInvestor (inv models.Investor) CreateInvestorReturn {
	fr := CreateInvestorReturn{
		ID:          inv.ID,
		UserId:      inv.UserId,
		CrowdfundId: inv.CrowdfundId,
		Fund:        inv.Fund,
	}
	return fr
}

func (f InvestorFormatter) FormatFindAllInvestor (list []models.Investor) []CreateInvestorReturn {
	var finalList []CreateInvestorReturn
	for _, item := range list {
		t := f.FormatCreateInvestor(item)
		finalList = append(finalList, t)
	}
	return finalList
}
