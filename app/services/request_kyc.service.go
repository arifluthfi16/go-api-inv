package services

import (
	"database/sql"
	"invest/app/models"
	"invest/app/utils"
)

type RequestKycService struct {}

func (req *RequestKycService) Create (data utils.CreateRequestKYCInput) (models.RequestKYC, error){
	kyc := models.RequestKYC{}

	kyc = models.RequestKYC{
		UserId:     "1",
		Filepath: 	data.Filepath,
		IsAccepted: false,
	}

	if err := Manager.DB.Create(&kyc).Error; err != nil {
		return kyc, err
	}
	return kyc, nil
}

func (req *RequestKycService) FindOneById (kycId string) (models.RequestKYC, error){
	var kyc models.RequestKYC
	result := Manager.DB.Take(&kyc, kycId)
	return kyc, result.Error
}

func(req *RequestKycService) FindAll () ([]models.RequestKYC, error){
	var kyc []models.RequestKYC
	result := Manager.DB.Find(&kyc)
	return kyc, result.Error
}

func (req *RequestKycService) DeleteById (trId string) (int64, error){
	var kyc models.RequestKYC
	result := Manager.DB.Delete(&kyc, trId)
	return result.RowsAffected, result.Error
}

func (req *RequestKycService) ApproveValidation (requestId string, adminId string) (models.RequestKYC, error){
	var kyc models.RequestKYC
	result := Manager.DB.Take(&kyc, requestId)
	if result.Error != nil {
		return kyc, result.Error
	}

	kyc.AdminId = sql.NullString{String: adminId}
	kyc.IsAccepted = true
	if err := Manager.DB.Save(&kyc).Error; err != nil {
		return kyc, err
	}
	return kyc, nil
}

func (req *RequestKycService) Update (kyc *models.RequestKYC, input utils.UpdateRequestKYCInput) (models.RequestKYC, error){
	kyc.AdminId = sql.NullString{String: input.AdminID}
	kyc.RejectReason = sql.NullString{String: input.RejectReason}
	kyc.IsAccepted = input.IsAccepted

	if err := Manager.DB.Save(&kyc).Error; err != nil {
		return *kyc, err
	}
	return *kyc, nil
}