package formatter

import "invest/app/models"

type AdminFormatter struct {}

type CreateAdminReturn struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Username string `json:"username"`
	HakAkses string `json:"hak_akses"`
}

func (f AdminFormatter) FormatCreateAdmin (admin models.Admin) CreateAdminReturn {
	formatter := CreateAdminReturn{
		ID:         admin.ID,
		Name:       admin.Name,
		Email:      admin.Email,
		Username: 	admin.Username,
		HakAkses: 	admin.HakAkses,
	}
	return formatter
}

func (f AdminFormatter) FormatAdminLogin (admin models.Admin, token string) AdminLoginReturn {
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