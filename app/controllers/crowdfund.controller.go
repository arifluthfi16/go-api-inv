package controllers

import (
	"github.com/gin-gonic/gin"
	"invest/app/services"
	"invest/app/utils"
	"invest/app/utils/formatter"
)

type CrowdfundController struct {}
var  CrowdfundService = services.CrowdfundService{}
var CrowdfundFormatter = formatter.CrowdfundFormatter{}

func (controller *CrowdfundController) Create(c *gin.Context){
	var createCrowdfundInput utils.CreateCrowdfundInput
	if err := c.ShouldBindJSON(&createCrowdfundInput); err != nil {
		Response.SendFailedResponse(utils.FailedResponseFormat{
			Context: c,
			Msg:     "Validation Error",
			Err:     err,
		}); return
	}

	crowdfundData, err := CrowdfundService.CreateCrowdfund(createCrowdfundInput)
	if  err != nil {
		Response.SendFailedResponse(utils.FailedResponseFormat{
			Context: c,
			Msg:     "Failed to create a crowdfunding",
			Err:     err,
		}); return
	}

	Response.SendSuccessResponse(utils.SuccessResponseFormat{
		Context: c,
		Msg:     "Successfully created a new crowdfunding",
		Data:    CrowdfundFormatter.FormatCreateCrowdfund(crowdfundData),
	})
}

func (controller *CrowdfundController) FindOneById(c *gin.Context) {
	crowdfundData, err := CrowdfundService.FindOneById(c.Param("id"))
	if err != nil {
		c.Error(utils.ValidationError{})
		return
	}

	Response.SendSuccessResponse(utils.SuccessResponseFormat{
		Context: c,
		Data:    CrowdfundFormatter.FormatCreateCrowdfund(crowdfundData),
	})
}

func (controller *CrowdfundController) FindAll(c *gin.Context){
	crowdfundingList, err := CrowdfundService.FindAll()
	if err != nil {
		Response.SendFailedResponse(utils.FailedResponseFormat{
			Context: c,
			Msg:     "Failed to get crowdfunding data",
			Err:     err,
		}); return
	}

	FormattedList := CrowdfundFormatter.FormatFindAllCrowdfunding(crowdfundingList)
	Response.SendSuccessResponse(utils.SuccessResponseFormat{
		Context: c,
		Data:    FormattedList,
	})
}

func (controller *CrowdfundController) DeleteById (c *gin.Context){
	_, err := CrowdfundService.DeleteById(c.Param("id"))
	if err != nil {
		Response.SendFailedResponse(utils.FailedResponseFormat{
			Context: c,
			Msg:     "Failed to delete, or no item deleted",
			Err:     err,
		}); return
	}
	Response.SendSuccessResponse(utils.SuccessResponseFormat{
		Context: c,
		Msg:     "Deletion successful",
	})
}

func (controller *CrowdfundController) Update(c *gin.Context){
	crowdfundData, err := CrowdfundService.FindOneById(c.Param("id"))
	if err != nil {
		Response.SendFailedResponse(utils.FailedResponseFormat{
			Context: c,
			Msg:     "Failed to find crowdfund data",
			Err:     err,
		}); return
	}

	var input utils.UpdateCrowdfundInput
	if err := c.ShouldBindJSON(&input); err != nil {
		Response.SendFailedResponse(utils.FailedResponseFormat{
			Context: c,
			Msg: "Validation Error",
			Err:     err,
		}); return
	}

	updateResult, err := CrowdfundService.Update(&crowdfundData, input)
	if err != nil {
		Response.SendFailedResponse(utils.FailedResponseFormat{
			Context: c,
			Msg: "Failed to update crowdfunding",
			Err:     err,
		}); return
	}

	Response.SendSuccessResponse(utils.SuccessResponseFormat{
		Context: c,
		Msg:     "Successfully updated a crowdfunding",
		Data:    CrowdfundFormatter.FormatCreateCrowdfund(updateResult),
	})
}