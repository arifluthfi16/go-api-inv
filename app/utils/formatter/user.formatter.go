package formatter

import "invest/app/models"

type UserFormatter struct {}

type CreateUserReturn struct {
	ID uint `json:"id"`
	DisplayPict string `json:"display_pict"`
	Name string `json:"name"`
	Email string `json:"email"`
	Phonenumber string `json:"phonenumber"`
	IsVerified bool `json:"is_verified"`
}


func (f UserFormatter) FormatCreateUser (user models.User) CreateUserReturn {
	formatter := CreateUserReturn{
		ID:          user.ID,
		DisplayPict: user.DisplayPict,
		Name:        user.Name,
		Email:       user.Email,
		Phonenumber: user.Phonenumber,
		IsVerified:  user.IsVerified,
	}
	return formatter
}

func (f UserFormatter) FormatFindAllUser (users []models.User) []CreateUserReturn {
	var finalList []CreateUserReturn
	for _, user := range users {
		t := f.FormatCreateUser(user)
		finalList = append(finalList, t)
	}
	return finalList
}