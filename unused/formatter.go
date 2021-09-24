package unused

import "invest/app/models"

type Formatter struct {}

type CreateUserReturn struct {
	ID uint `json:"id"`
	DisplayPict string `json:"display_pict"`
	Name string `json:"name"`
	Email string `json:"email"`
	Phonenumber string `json:"phonenumber"`
	IsVerified bool `json:"is_verified"`
}

type CreateAdminReturn struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Username string `json:"username"`
	HakAkses string `json:"hak_akses"`
}

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

type CreateCrowdfundReturn struct {
	ID uint `json:"id"`
	AdminID string `json:"admin_id"`
	Title string `json:"title"`
	Need int `json:"need"`
	Earn int `json:"earn"`
	Description string `json:"description"`
	IsActive bool `json:"is_active"`
}

type CreateFAQReturn struct {
	ID uint `json:"id"`
	AdminID string `json:"admin_id"`
	Pertanyaan string `json:"pertanyaan"`
	Jawaban string `json:"jawaban"`
	Kategori string `json:"kategori"`
	Count int `json:"count"`
}

type CreateInvestorReturn struct {
	ID uint `json:"id"`
	UserId string `json:"user_id"`
	CrowdfundId string `json:"crowdfund_id"`
	Fund int `json:"fund"`
}

type CreateTransparansiReturn struct {
	ID uint `json:"id"`
	AdminId string `json:"admin_id"`
	CrowdfundId string `json:"crowdfund_id"`
	Title string `json:"title"`
	ShortDesc string `json:"short_desc"`
	Content	string `json:"content"`
	Thumbnail string `json:"thumbnail"`
	IsPublished bool `json:"is_published"`
}

type CreateRequestKYCReturn struct {
	ID uint `json:"id"`
	UserId string `json:"user_id"`
	AdminId string `json:"admin_id"`
	Filepath string `json:"file_path"`
	RejectReason string `json:"reject_reason"`
	IsAccepted bool `json:"is_accepted"`
}

func (f Formatter) FormatCreateUser (user models.User) CreateUserReturn {
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

func (f Formatter) FormatFindAllUser (users []models.User) []CreateUserReturn {
	var finalList []CreateUserReturn
	for _, user := range users {
		t := f.FormatCreateUser(user)
		finalList = append(finalList, t)
	}
	return finalList
}

func (f Formatter) FormatCreateAdmin (admin models.Admin) CreateAdminReturn {
	formatter := CreateAdminReturn{
		ID:         admin.ID,
		Name:       admin.Name,
		Email:      admin.Email,
		Username: 	admin.Username,
		HakAkses: 	admin.HakAkses,
	}
	return formatter
}

func (f Formatter) FormatLoginReturn (user models.User, token string) LoginReturn {
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

func (f Formatter) FormatAdminLogin (admin models.Admin, token string) AdminLoginReturn {
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

func (f Formatter) FormatCreateCrowdfund (crowdfund models.Crowdfund) CreateCrowdfundReturn {
	formatter := CreateCrowdfundReturn{
		ID:          crowdfund.ID,
		AdminID:     crowdfund.AdminId,
		Title:       crowdfund.Title,
		Need:        crowdfund.Need,
		Earn:        crowdfund.Earn,
		Description: crowdfund.Description,
		IsActive:    crowdfund.IsActive,
	}
	return formatter
}

func (f Formatter) FormatFindAllCrowdfunding (list []models.Crowdfund) []CreateCrowdfundReturn {
	var finalList []CreateCrowdfundReturn
	for _, item := range list {
		t := f.FormatCreateCrowdfund(item)
		finalList = append(finalList, t)
	}
	return finalList
}

func (f Formatter) FormatCreateFAQ (faq models.FAQ) CreateFAQReturn {
	fr := CreateFAQReturn{
		ID:         faq.ID,
		AdminID:    faq.AdminId,
		Pertanyaan: faq.Pertanyaan,
		Jawaban:    faq.Jawaban,
		Kategori:   faq.Jawaban,
		Count:      faq.Count,
	}
	return fr
}

func (f Formatter) FormatFindAllFAQ (list []models.FAQ) []CreateFAQReturn {
	var finalList []CreateFAQReturn
	for _, item := range list {
		t := f.FormatCreateFAQ(item)
		finalList = append(finalList, t)
	}
	return finalList
}

func (f Formatter) FormatCreateTransparansi (tr models.Transparansi) CreateTransparansiReturn {
	fr := CreateTransparansiReturn{
		ID:          tr.ID,
		AdminId:     tr.AdminId,
		CrowdfundId: tr.CrowdfundId,
		Title:       tr.Title,
		ShortDesc:   tr.ShortDesc,
		Content:     tr.Content,
		Thumbnail:   tr.Thumbnail,
		IsPublished: tr.IsPublished,
	}
	return fr
}

func (f Formatter) FormatFindAllTransparansi (list []models.Transparansi) []CreateTransparansiReturn {
	var finalList []CreateTransparansiReturn
	for _, item := range list {
		t := f.FormatCreateTransparansi(item)
		finalList = append(finalList, t)
	}
	return finalList
}

func (f Formatter) FormatCreateInvestor (inv models.Investor) CreateInvestorReturn {
	fr := CreateInvestorReturn{
		ID:          inv.ID,
		UserId:      inv.UserId,
		CrowdfundId: inv.CrowdfundId,
		Fund:        inv.Fund,
	}
	return fr
}

func (f Formatter) FormatFindAllInvestor (list []models.Investor) []CreateInvestorReturn {
	var finalList []CreateInvestorReturn
	for _, item := range list {
		t := f.FormatCreateInvestor(item)
		finalList = append(finalList, t)
	}
	return finalList
}

func (f Formatter) FormatCreateRequestKYC (kyc models.RequestKYC) CreateRequestKYCReturn {
	fr := CreateRequestKYCReturn{
		ID:           kyc.ID,
		UserId:       kyc.UserId,
		AdminId:      kyc.AdminId.String,
		Filepath:     kyc.Filepath,
		RejectReason: kyc.RejectReason.String,
		IsAccepted:   kyc.IsAccepted,
	}
	return fr
}

func (f Formatter) FormatFindAllRequestKYC (list []models.RequestKYC) []CreateRequestKYCReturn {
	var finalList []CreateRequestKYCReturn
	for _, item := range list {
		t := f.FormatCreateRequestKYC(item)
		finalList = append(finalList, t)
	}
	return finalList
}