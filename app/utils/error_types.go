package utils

type ValidationError struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
}

func (validationErr ValidationError) Error() string{
	return "Validation error"
}