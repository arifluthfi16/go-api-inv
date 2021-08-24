package utils

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type appError struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
}

func (e appError) Error() string{
	return "Error"
}

type ErrorHandler struct {}

func (erh ErrorHandler) JSONAppErrorReporter() gin.HandlerFunc {
	return erh.jsonAppErrorReporterT(gin.ErrorTypeAny)
}

func (erh ErrorHandler) jsonAppErrorReporterT(errType gin.ErrorType) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		detectedErrors := c.Errors.ByType(errType)

		if len(detectedErrors) > 0 {
			err := detectedErrors[0].Err
			log.Println(err)

			var parsedError *appError
			switch err.(type) {
			case *appError:
				parsedError = err.(*appError)
			case *ValidationError:
				log.Println("Validation Error")
			default:
				parsedError = &appError{
					Code: http.StatusInternalServerError,
					Message: "Internal Server Error",
				}
			}

			c.AbortWithStatusJSON(parsedError.Code, parsedError)
			return
		}
	}
}