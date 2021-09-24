package controllers

import (
	"github.com/gin-gonic/gin"
	"invest/app/services"
	"invest/app/utils"
	"invest/app/utils/formatter"
)

type InvestorController struct {}

var InvestorService = services.InvestorService{}
var CrowdfundingService = services.CrowdfundService{}
var InvestorFormatter = formatter.InvestorFormatter{}

func (controller *InvestorController) Create(c *gin.Context){
	var createInvestorInput utils.CreateInvestorInput
	if err := c.ShouldBindJSON(&createInvestorInput); err != nil {
		Response.SendFailedResponse(utils.FailedResponseFormat{
			Context: c,
			Msg:     "Validation Error",
			Err:     err,
		}); return
	}

	inv, err := InvestorService.Create(createInvestorInput)
	if  err != nil {
		Response.SendFailedResponse(utils.FailedResponseFormat{
			Context: c,
			Msg:     "Failed to create new investor",
			Err:     err,
		}); return
	}

	_, err = CrowdfundingService.AddEarn(createInvestorInput.CrowdfundId, createInvestorInput.Fund)
	if err != nil {
		Response.SendFailedResponse(utils.FailedResponseFormat{
			Context: c,
			Msg:     "Failed to add fund to crowdfunding",
			Err:     err,
		}); return
	}

	formattedResponse := InvestorFormatter.FormatCreateInvestor(inv)
	Response.SendSuccessResponse(utils.SuccessResponseFormat{
		Context: c,
		Msg:     "Successfully created a new investor data",
		Data:    formattedResponse,
	})
}

func (controller *InvestorController) FindOneById(c *gin.Context) {
	inv, err := InvestorService.FindOneById(c.Param("id"))
	if err != nil {
		Response.SendFailedResponse(utils.FailedResponseFormat{
			Context: c,
			Msg:     "Failed to find investor  data",
			Err:     err,
		})
	}else{
		formattedResponse := InvestorFormatter.FormatCreateInvestor(inv)
		Response.SendSuccessResponse(utils.SuccessResponseFormat{
			Context: c,
			Data:    formattedResponse,
		})
	}
	return
}

func (controller *InvestorController) FindAllByCrowdfundId(c *gin.Context) {
	list, err := InvestorService.FindAllByCrowdfundId(c.Param("crowdfund_id"))
	if err != nil {
		Response.SendFailedResponse(utils.FailedResponseFormat{
			Context: c,
			Msg:     "Failed to find investor  data",
			Err:     err,
		})
	}else{
		formattedResponse := InvestorFormatter.FormatFindAllInvestor(list)
		Response.SendSuccessResponse(utils.SuccessResponseFormat{
			Context: c,
			Data:    formattedResponse,
		})
	}
	return
}

func (controller *InvestorController) FindAll(c *gin.Context){
	list, err := InvestorService.FindAll()
	if err != nil {
		Response.SendFailedResponse(utils.FailedResponseFormat{
			Context: c,
			Msg:     "Failed to find investor  data",
			Err:     err,
		})
	}else{
		formattedResponse := InvestorFormatter.FormatFindAllInvestor(list)
		Response.SendSuccessResponse(utils.SuccessResponseFormat{
			Context: c,
			Data:    formattedResponse,
		})
	}
	return
}

func (controller *InvestorController) DeleteById (c *gin.Context){
	InvId := c.Param("id")
	_, err := InvestorService.DeleteById(InvId)
	if err != nil {
		Response.SendFailedResponse(utils.FailedResponseFormat{
			Context: c,
			Msg:     "Failed to delete, or no item deleted",
			Err:     err,
		})
	}else{
		Response.SendSuccessResponse(utils.SuccessResponseFormat{
			Context: c,
			Msg:     "Successfully deleted data",
		})
	}
	return
}

func (controller *InvestorController) Update(c *gin.Context){
	inv, err := InvestorService.FindOneById(c.Param("id"))
	if err != nil {
		Response.SendFailedResponse(utils.FailedResponseFormat{
			Context: c,
			Msg:     "Failed to find investor data",
			Err:     err,
		}); return
	}

	var input utils.UpdateInvestorInput
	if err := c.ShouldBindJSON(&input); err != nil {
		Response.SendFailedResponse(utils.FailedResponseFormat{
			Context: c,
			Msg:     "Validation Error",
			Err:     err,
		}); return
	}

	result, err := InvestorService.Update(&inv, input)
	if err != nil {
		Response.SendFailedResponse(utils.FailedResponseFormat{
			Context: c,
			Msg:     "Failed to update Investor",
			Err:     err,
		}); return
	}

	formattedResponse := InvestorFormatter.FormatCreateInvestor(result)
	Response.SendSuccessResponse(utils.SuccessResponseFormat{
		Context: c,
		Msg:     "Successfully updated investor data",
		Data:    formattedResponse,
	})
	return
}