package controllers

import (
	"github.com/gin-gonic/gin"
	Service "invest/app/services"
	"invest/app/utils"
	"invest/app/utils/formatter"
)

type AdminController struct {}
var AdminService = Service.AdminService{}
var Response = utils.ResponseFormatter{}
var AdminFormatter = formatter.AdminFormatter{}

func (adm AdminController) CreateAdmin(c *gin.Context) {
	var adminInput utils.CreateAdminInput
	if err := c.ShouldBindJSON(&adminInput); err != nil {
		Response.SendFailedResponse(utils.FailedResponseFormat{
			Context: c,
			Msg:     "Validation Error",
			Err:     err,
		}); return
	}

	admin, err := AdminService.CreateAdmin(adminInput)
	if  err != nil {
		Response.SendFailedResponse(utils.FailedResponseFormat{
			Context: c,
			Msg:     "Failed to create user",
			Err:     err,
		}); return
	}

	Response.SendSuccessResponse(utils.SuccessResponseFormat{
		Context: c,
		Msg : "Successfully created new admin",
		Data:    admin,
	})
}

func (adm AdminController) FindById(c *gin.Context){
	adminData, err := AdminService.FindOneById(c.Param("admin_id"))
	if err != nil {
		Response.SendFailedResponse(utils.FailedResponseFormat{
			Context: c,
			Msg:     "Failed to find admin",
			Err:     err,
		}); return
	}
	Response.SendSuccessResponse(utils.SuccessResponseFormat{
		Context: c,
		Data:    adminData,
	})
}