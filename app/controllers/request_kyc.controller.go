package controllers

import (
	"fmt"
	//"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"invest/app/services"
	"invest/app/utils"
	"path/filepath"
)

type RequestKYCController struct {}
var RequestKYCService = services.RequestKycService{}

func (controller *RequestKYCController) Create(c *gin.Context){
	fmt.Println(c.GetHeader("Content-Type"))

	var formDataRequestKYC utils.RequestKYCFormData
	err := c.ShouldBind(&formDataRequestKYC)
	if err != nil {
		Response.SendFailedResponse(utils.FailedResponseFormat{
			Context: c,
			Msg:     "Validation Error",
			Err:     err,
		}); return
	}

	extension := filepath.Ext(formDataRequestKYC.File.Filename)
	newFileName := uuid.New().String()+extension
	if err := c.SaveUploadedFile(formDataRequestKYC.File, "./uploads/" + newFileName); err != nil {
		Response.SendFailedResponse(utils.FailedResponseFormat{
			Context: c,
			Msg:     "Failed to save file",
			Err:     err,
		}); return
	}

	createKYCInput := utils.CreateRequestKYCInput{
		UserId:   formDataRequestKYC.UserId,
		Filepath: newFileName,
	}

	kyc, err := RequestKYCService.Create(createKYCInput)
	if  err != nil {
		Response.SendFailedResponse(utils.FailedResponseFormat{
			Context: c,
			Msg:     "Failed to create new Request KYC",
			Err:     err,
		}); return
	}

	formattedResponse := Formatter.FormatCreateRequestKYC(kyc)
	Response.SendSuccessResponse(utils.SuccessResponseFormat{
		Context: c,
		Msg:     "Successfully created a KYC Validation REQUEST",
		Data:    formattedResponse,
	})
}

func (controller *RequestKYCController) FindOneById(c *gin.Context) {
	requestKycData, err := RequestKYCService.FindOneById(c.Param("id"))
	if err != nil {
		Response.SendFailedResponse(utils.FailedResponseFormat{
			Context: c,
			Msg:     "Failed to find KYC data",
			Err:     err,
		}); return

	}
	Response.SendSuccessResponse(utils.SuccessResponseFormat{
		Context: c,
		Data:    requestKycData,
	})
}

func (controller *RequestKYCController) FindAllKYCRequest(c *gin.Context) {
	list, err := RequestKYCService.FindAll()
	if err != nil {
		Response.SendFailedResponse(utils.FailedResponseFormat{
			Context: c,
			Msg:     "Failed to find any KYC Request",
			Err:     err,
		})
	}else{
		formattedResponse := Formatter.FormatFindAllRequestKYC(list)
		Response.SendSuccessResponse(utils.SuccessResponseFormat{
			Context: c,
			Data:    formattedResponse,
		})
	}
	return
}

func (controller *RequestKYCController) DeleteById (c *gin.Context){
	reqId := c.Param("id")
	_, err := RequestKYCService.DeleteById(reqId)
	if err != nil {
		Response.SendFailedResponse(utils.FailedResponseFormat{
			Context: c,
			Msg:     "Failed to delete, or no item deleted",
			Err:     err,
		})
	}else{
		Response.SendSuccessResponse(utils.SuccessResponseFormat{
			Context: c,
			Msg:     "Successfully delete data",
		})
	}
	return
}

func (controller *RequestKYCController) ApproveValidation (c *gin.Context){
	var formDataRequestKYC utils.ApproveKYCInput
	err := c.ShouldBind(&formDataRequestKYC)
	if err!= nil {
		Response.SendFailedResponse(utils.FailedResponseFormat{
			Context: c,
			Msg:     "Validation Error",
			Err:     err,
		}); return
	}

	result,err := RequestKYCService.ApproveValidation(formDataRequestKYC.RequestId, formDataRequestKYC.AdminId)
	if err!= nil {
		Response.SendFailedResponse(utils.FailedResponseFormat{
			Context: c,
			Msg:     "Failed to approve KYC",
			Err:     err,
		}); return
	}

	formattedResponse := Formatter.FormatCreateRequestKYC(result)
	Response.SendSuccessResponse(utils.SuccessResponseFormat{
		Context: c,
		Msg:     "Successfully approve a KYC",
		Data:    formattedResponse,
	})
}


func (controller *RequestKYCController) Update(c *gin.Context){
	reqKYC, err := RequestKYCService.FindOneById(c.Param("id"))
	if err != nil {
		Response.SendFailedResponse(utils.FailedResponseFormat{
			Context: c,
			Msg:     "Failed to find request kyc data",
			Err:     err,
		}); return
	}

	var input utils.UpdateRequestKYCInput
	if err := c.ShouldBindJSON(&input); err != nil {
		Response.SendFailedResponse(utils.FailedResponseFormat{
			Context: c,
			Msg:     "Validation Error",
			Err:     err,
		}); return
	}

	result, err := RequestKYCService.Update(&reqKYC, input)
	if err != nil {
		Response.SendFailedResponse(utils.FailedResponseFormat{
			Context: c,
			Msg:     "Failed to update KYC",
			Err:     err,
		}); return
	}

	formattedResponse := Formatter.FormatCreateRequestKYC(result)
	Response.SendSuccessResponse(utils.SuccessResponseFormat{
		Context: c,
		Msg:     "Successfully updated KYC data",
		Data:    formattedResponse,
	})
	return
}