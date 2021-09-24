package formatter

import "invest/app/models"

type RequestKYCFormatter struct {}

type CreateRequestKYCReturn struct {
	ID uint `json:"id"`
	UserId string `json:"user_id"`
	AdminId string `json:"admin_id"`
	Filepath string `json:"file_path"`
	RejectReason string `json:"reject_reason"`
	IsAccepted bool `json:"is_accepted"`
}

func (f RequestKYCFormatter) FormatCreateRequestKYC (kyc models.RequestKYC) CreateRequestKYCReturn {
	fr := CreateRequestKYCReturn{
		ID:           kyc.ID,
		UserId:       kyc.UserId,
		AdminId:      kyc.AdminId.String,
		Filepath:     kyc.Filepath,
		RejectReason: kyc.RejectReason.String,
		IsAccepted:   kyc.IsAccepted,
	}
	return fr
}

func (f RequestKYCFormatter) FormatFindAllRequestKYC (list []models.RequestKYC) []CreateRequestKYCReturn {
	var finalList []CreateRequestKYCReturn
	for _, item := range list {
		t := f.FormatCreateRequestKYC(item)
		finalList = append(finalList, t)
	}
	return finalList
}