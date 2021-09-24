package controllers

import (
	"github.com/gin-gonic/gin"
	"invest/app/services"
	"invest/app/utils"
	"invest/app/utils/formatter"
)

type FAQController struct {}

var FAQService = services.FAQService{}
var FAQFormatter = formatter.FAQFormatter{}

func (controller *FAQController) Create(c *gin.Context){
	var createFAQInput utils.CreateFAQInput
	if err := c.ShouldBindJSON(&createFAQInput); err != nil {
		Response.SendFailedResponse(utils.FailedResponseFormat{
			Context: c,
			Msg:     "Validation Error",
			Err:     err,
		}); return
	}
	faq, err := FAQService.Create(createFAQInput)
	if  err != nil {
		Response.SendFailedResponse(utils.FailedResponseFormat{
			Context: c,
			Msg:     "Failed to create new faq",
			Err:     err,
		}); return
	}
	Response.SendSuccessResponse(utils.SuccessResponseFormat{
		Context: c,
		Msg:     "Successfully created a new FAQ",
		Data:    FAQFormatter.FormatCreateFAQ(faq),
	})
}

func (controller *FAQController) FindOneById(c *gin.Context) {
	faq, err := FAQService.FindOneById(c.Param("faq_id"))
	if err != nil {
		Response.SendFailedResponse(utils.FailedResponseFormat{
			Context: c,
			Msg:     "Failed to find faq",
			Err:     err,
		}); return
	}
	Response.SendSuccessResponse(utils.SuccessResponseFormat{
		Context: c,
		Data:    FAQFormatter.FormatCreateFAQ(faq),
	})
}

func (controller *FAQController) FindAll(c *gin.Context){
	list, err := FAQService.FindAll()
	if err != nil {
		Response.SendFailedResponse(utils.FailedResponseFormat{
			Context: c,
			Msg:     "Failed to fetch FAQ data",
			Err:     err,
		}); return
	}
	FormattedList := FAQFormatter.FormatFindAllFAQ(list)
	Response.SendSuccessResponse(utils.SuccessResponseFormat{
		Context: c,
		Data:    FormattedList,
	})
}

func (controller *FAQController) DeleteById (c *gin.Context){
	id := c.Param("faq_id")
	_, err := FAQService.DeleteById(id)
	if err != nil {
		Response.SendFailedResponse(utils.FailedResponseFormat{
			Context: c,
			Msg:     "Failed to delete, or no item deleted",
			Err:     err,
		}); return
	}
	Response.SendSuccessResponse(utils.SuccessResponseFormat{
		Context: c,
		Msg : "Successfully deleted faq",
	})
}

func (controller *FAQController) Update(c *gin.Context){
	faq, err := FAQService.FindOneById(c.Param("faq_id"))
	if err != nil {
		Response.SendFailedResponse(utils.FailedResponseFormat{
			Context: c,
			Msg:     "Failed to find faq",
			Err:     err,
		}); return
	}

	var input utils.UpdateFAQInput
	if err := c.ShouldBindJSON(&input); err != nil {
		Response.SendFailedResponse(utils.FailedResponseFormat{
			Context: c,
			Msg:     "Validation Error",
			Err:     err,
		}); return
	}

	result, err := FAQService.Update(&faq, input)
	if err != nil {
		Response.SendFailedResponse(utils.FailedResponseFormat{
			Context: c,
			Msg:     "Failed to update FAQ",
			Err:     err,
		}); return
	}

	Response.SendSuccessResponse(utils.SuccessResponseFormat{
		Context: c,
		Msg : "Successfully update faq",
		Data : FAQFormatter.FormatCreateFAQ(result),
	})
}