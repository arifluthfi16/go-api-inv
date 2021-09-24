package formatter

import "invest/app/models"

type AuthFormatter struct {}

type LoginReturn struct {
	ID uint `json:"id"`
	DisplayPict string `json:"display_pict"`
	Name string `json:"name"`
	Email string `json:"email"`
	Phonenumber string `json:"phonenumber"`
	IsVerified bool `json:"is_verified"`
	Token string `json:"token"`
}

type AdminLoginReturn struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Username string `json:"username"`
	HakAkses string `json:"hak_akses"`
	Token string `json:"token"`
}

func (f AuthFormatter) FormatLoginReturn (user models.User, token string) LoginReturn {
	formatter := LoginReturn{
		ID:          user.ID,
		DisplayPict: user.DisplayPict,
		Name:        user.Name,
		Email:       user.Email,
		Phonenumber: user.Phonenumber,
		IsVerified:  user.IsVerified,
		Token: token,
	}
	return formatter
}

func (f AuthFormatter) FormatAdminLogin (admin models.Admin, token string) AdminLoginReturn {
	formatter := AdminLoginReturn{
		ID:          	admin.ID,
		Name:        	admin.Name,
		Email:       	admin.Email,
		Username: 	 	admin.Username,
		HakAkses: 		admin.HakAkses,
		Token: 			token,
	}
	return formatter
}