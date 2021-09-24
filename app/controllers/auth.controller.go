package controllers

import (
	"github.com/gin-gonic/gin"
	Service "invest/app/services"
	"invest/app/utils"
	"invest/app/utils/formatter"
)

type AuthController struct{}
var LoginService = Service.AuthService{}
var AuthFormatter = formatter.AuthFormatter{}

func (au AuthController) Login (c *gin.Context){
	var LoginInput utils.LoginInput
	if err := c.ShouldBindJSON(&LoginInput); err != nil {
		Response.SendFailedResponse(utils.FailedResponseFormat{
			Context: c,
			Msg:     "Validation Error",
			Err:     err,
		}); return
	}

	user, err := LoginService.AuthenticateUser(LoginInput)
	if err != nil {
		Response.SendFailedResponse(utils.FailedResponseFormat{
			Context: c,
			Msg:     "Failed to login",
			Err:     err,
		}); return
	}

	token, err := LoginService.CreateToken(user)
	if err != nil {
		Response.SendFailedResponse(utils.FailedResponseFormat{
			Context: c,
			Msg:     "Token generation failed",
			Err:     err,
		}); return
	}

	loginResponse := AuthFormatter.FormatLoginReturn(user, token)
	Response.SendSuccessResponse(utils.SuccessResponseFormat{
		Context: c,
		Msg:     "Login succeed",
		Data:    loginResponse,
	})
}

func (au AuthController) AdminLogin (c *gin.Context){
	var LoginInput utils.AdminLoginInput
	if err := c.ShouldBindJSON(&LoginInput); err != nil {
		Response.SendFailedResponse(utils.FailedResponseFormat{
			Context: c,
			Msg:     "Validation Error",
			Err:     err,
		}); return
	}

	admin, err := LoginService.AuthenticateAdmin(LoginInput)
	if err != nil {
		Response.SendFailedResponse(utils.FailedResponseFormat{
			Context: c,
			Msg:     "Failed to login",
			Err:     err,
		}); return
	}

	token, err := LoginService.CreateAdminToken(admin)
	if err != nil {
		Response.SendFailedResponse(utils.FailedResponseFormat{
			Context: c,
			Msg:     "Token generation failed",
			Err:     err,
		}); return
	}

	loginResponse := AuthFormatter.FormatAdminLogin(admin, token)
	Response.SendSuccessResponse(utils.SuccessResponseFormat{
		Context: c,
		Msg:     "Login succeed",
		Data:    loginResponse,
	})
}

func (au AuthController) VerifyAccess() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := LoginService.ExtractToken(c.Request.Header.Get("Authorization"))
		userData,err := LoginService.VerifyToken(tokenString)
		if err != nil{
			Response.SendFailedResponse(utils.FailedResponseFormat{
				Context: c,
				Msg:     "Failed to verify token",
				Err:     err,
			}); return
		}
		c.Set("login_info", userData)
	}
}