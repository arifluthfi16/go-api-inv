package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	Service "invest/app/services"
	"invest/app/utils"
	"invest/app/utils/formatter"
)

var UserService = Service.UserService{}
var UserFormatter = formatter.UserFormatter{}

func CreateUser(c *gin.Context) {
	fmt.Println("Creating User")
	var createUserInputFormat utils.CreateUserInput
	if err := c.ShouldBindJSON(&createUserInputFormat); err != nil {
		FormattedErr := Response.APIResponseError("Validation Error", err)
		Response.HandleFailedRequest(c, FormattedErr)
		return
	}

	user, err := UserService.CreateUser(createUserInputFormat)
	if  err != nil {
		FormattedErr := Response.APIResponseError("Failed to create user", err)
		Response.HandleFailedRequest(c, FormattedErr)
		return
	}

	formatted := UserFormatter.FormatCreateUser(user)
	payload := Response.APIResponse("Succesfully created a new user", formatted)
	Response.HandleSuccessRequest(c, payload)
	return
}

func FindUserById(c *gin.Context){
	user, err := UserService.FindOneById(c.Param("userid"))
	if err != nil {
		FormattedErr := Response.APIResponseError("Failed to find user", err)
		Response.HandleFailedRequest(c, FormattedErr)
	}else{
		UserData := UserFormatter.FormatCreateUser(user)
		Return := Response.APIResponse("Succeed finding all user", UserData)
		Response.HandleSuccessRequest(c, Return)
	}
	return
}

func FindAllUser(c *gin.Context){
	users, err := UserService.FindAll()
	if err != nil {
		FormattedErr := Response.APIResponseError("Failed to find user", err)
		Response.HandleFailedRequest(c, FormattedErr)
	}else{
		UsersData := UserFormatter.FormatFindAllUser(users)
		Return := Response.APIResponse("Succeed finding all user", UsersData)
		Response.HandleSuccessRequest(c, Return)
	}
	return
}

func DeleteUserById(c *gin.Context){
	userId := c.Param("userid")
	_, err := UserService.DeleteOneById(userId)
	if err != nil {
		FormattedErr := Response.APIResponseError("Failed to delete user, or no user deleted", err)
		Response.HandleFailedRequest(c, FormattedErr)
	}else{
		Return := Response.APIResponse("Succeed deleting user", nil)
		Response.HandleSuccessRequest(c, Return)
	}
	return
}

func UpdateUser(c *gin.Context){
	user, err := UserService.FindOneById(c.Param("userid"))
	if err != nil {
		FormattedErr := Response.APIResponseError("Failed to find user", err)
		Response.HandleFailedRequest(c, FormattedErr)
		return
	}

	var input utils.CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		FormattedErr := Response.APIResponseError("Validation Error", err)
		Response.HandleFailedRequest(c, FormattedErr)
		return
	}

	result, err := UserService.UpdateUser(&user, input)
	if err != nil {
		FormattedErr := Response.APIResponseError("Failed to update user Error", err)
		Response.HandleFailedRequest(c, FormattedErr)
		return
	}

	formatted := UserFormatter.FormatCreateUser(result)
	payload := Response.APIResponse("Succesfully updated user", formatted)
	Response.HandleSuccessRequest(c, payload)
	return
}