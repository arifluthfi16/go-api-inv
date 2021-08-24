package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Status bool `json:"status"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}

type PlainResponse struct {
	Status bool `json:"status"`
	Data interface{} `json:"data"`
}

type MessageResponse struct {
	Status bool `json:"status"`
	Msg string `json:"msg"`
}

type ResponseError struct {
	Status bool `json:"status"`
	Msg string `json:"msg"`
	Error interface{} `json:"error"`
}

type FailedResponseFormat struct {
	Context *gin.Context
	Msg string
	Err error
}

type SuccessResponseFormat struct {
	Context *gin.Context
	Msg string
	Data interface{}
}

type ResponseFormatter struct {}

func (rf ResponseFormatter) APIResponse(msg string,  data interface{}) Response{
	r := Response{
		Status: true,
		Msg: msg,
		Data: data,
	}
	return r
}

func (rf ResponseFormatter) APIResponseError(msg string, err error) ResponseError{
	r := ResponseError{
		Status: false,
		Msg:  msg  ,
		Error:   err.Error(),
	}
	return r
}

func (rf ResponseFormatter) HandleFailedRequest(c *gin.Context, payload ResponseError){
	c.AbortWithStatusJSON(http.StatusBadRequest, payload)
}

func (rf ResponseFormatter) HandleSuccessRequest(c *gin.Context, payload Response){
	c.AbortWithStatusJSON(http.StatusOK, payload)
}

func (rf ResponseFormatter) SendFailedResponse (res FailedResponseFormat){
	formattedResponse := rf.APIResponseError(res.Msg, res.Err)
	res.Context.AbortWithStatusJSON(http.StatusOK, formattedResponse)
}

func (rf ResponseFormatter) SendSuccessResponse(res SuccessResponseFormat){
	var formattedResponse interface{}
	if res.Msg == ""{
		formattedResponse = PlainResponse{
			Status: true,
			Data:   res.Data,
		}
	} else if res.Data == nil {
		formattedResponse = MessageResponse{
			Status: true,
			Msg: res.Msg,
		}
	} else {
		formattedResponse = Response{
			Status: true,
			Data:   res.Data,
			Msg: res.Msg,
		}
	}
	res.Context.AbortWithStatusJSON(http.StatusOK, formattedResponse)
}



