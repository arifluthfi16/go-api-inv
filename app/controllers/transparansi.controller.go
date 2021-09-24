package controllers

import (
	"github.com/gin-gonic/gin"
	"invest/app/services"
	"invest/app/utils"
	"invest/app/utils/formatter"
)

type TransparansiController struct {}
var (
	TransparansiService = services.TransparansiService{}
	TransparansiFormatter = formatter.TransparansiFormatter{}
)


func (controller *TransparansiController) Create(c *gin.Context){
	var createTransparansiInput utils.CreateTransparansiInput
	if err := c.ShouldBindJSON(&createTransparansiInput); err != nil {
		FormattedErr := Response.APIResponseError("Validation Error", err)
		Response.HandleFailedRequest(c, FormattedErr)
		return
	}

	tr, err := TransparansiService.Create(createTransparansiInput)
	if  err != nil {
		FormattedErr := Response.APIResponseError("Failed to create new transparansi", err)
		Response.HandleFailedRequest(c, FormattedErr)
		return
	}

	formatted := TransparansiFormatter.FormatCreateTransparansi(tr)
	payload := Response.APIResponse("Succesfully created a new transparansi", formatted)
	Response.HandleSuccessRequest(c, payload)
	return
}

func (controller *TransparansiController) FindOneById(c *gin.Context) {
	tr, err := TransparansiService.FindOneById(c.Param("id"))
	if err != nil {
		FormattedErr := Response.APIResponseError("Failed to find Transparansi", err)
		Response.HandleFailedRequest(c, FormattedErr)
	}else{
		TrData := TransparansiFormatter.FormatCreateTransparansi(tr)
		Return := Response.APIResponse("Transparansi found", TrData)
		Response.HandleSuccessRequest(c, Return)
	}
	return
}

func (controller *TransparansiController) FindAll(c *gin.Context){
	list, err := TransparansiService.FindAll()
	if err != nil {
		FormattedErr := Response.APIResponseError("Failed to fetch Transparansi data", err)
		Response.HandleFailedRequest(c, FormattedErr)
	}else{
		FormattedData := TransparansiFormatter.FormatFindAllTransparansi(list)
		Return := Response.APIResponse("Succeed finding all transparansi", FormattedData)
		Response.HandleSuccessRequest(c, Return)
	}
	return
}

func (controller *TransparansiController) DeleteById (c *gin.Context){
	TrId := c.Param("id")
	_, err := FAQService.DeleteById(TrId)
	if err != nil {
		FormattedErr := Response.APIResponseError("Failed to delete, or no item deleted", err)
		Response.HandleFailedRequest(c, FormattedErr)
	}else{
		Return := Response.APIResponse("Succeed deleting Transparansi", nil)
		Response.HandleSuccessRequest(c, Return)
	}
	return
}

func (controller *TransparansiController) Update(c *gin.Context){
	tr, err := TransparansiService.FindOneById(c.Param("id"))
	if err != nil {
		FormattedErr := Response.APIResponseError("Failed to find Transparansi", err)
		Response.HandleFailedRequest(c, FormattedErr)
		return
	}

	var input utils.UpdateTransparansiInput
	if err := c.ShouldBindJSON(&input); err != nil {
		FormattedErr := Response.APIResponseError("Validation Error", err)
		Response.HandleFailedRequest(c, FormattedErr)
		return
	}

	result, err := TransparansiService.Update(&tr, input)
	if err != nil {
		FormattedErr := Response.APIResponseError("Failed to update Transparansi", err)
		Response.HandleFailedRequest(c, FormattedErr)
		return
	}

	formatted := TransparansiFormatter.FormatCreateTransparansi(result)
	payload := Response.APIResponse("Succesfully updated Transparansi", formatted)
	Response.HandleSuccessRequest(c, payload)
	return
}