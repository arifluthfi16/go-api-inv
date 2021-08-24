package utils

import "mime/multipart"

type CreateUserInput struct {
	Name string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Phonenumber string `json:"phonenumber" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateAdminInput struct {
	Name string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email string `json:"email" binding:"required"`
	HakAkses string `json:"hak_akses"`
}

type LoginInput struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AdminLoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateCrowdfundInput struct {
	AdminID string `json:"admin_id"`
	Title string `json:"title"`
	Need int `json:"need"`
	Description string `json:"description"`
}

type UpdateCrowdfundInput struct{
	Title string `json:"title"`
	Need int `json:"need"`
	Earn int `json:"earn"`
	Description string `json:"description"`
	IsActive bool `json:"is_active"`
}

type CreateFAQInput struct {
	AdminID string `json:"admin_id"`
	Pertanyaan string `json:"pertanyaan"`
	Jawaban string `json:"jawaban"`
	Kategori string `json:"kategori"`
}

type UpdateFAQInput struct {
	Pertanyaan string `json:"pertanyaan"`
	Jawaban string `json:"jawaban"`
	Kategori string `json:"kategori"`
	Count int `json:"count"`
}

type CreateTransparansiInput struct {
	AdminId string `json:"admin_id"`
	CrowdfundId string `json:"crowdfund_id"`
	Title string `json:"title"`
	ShortDesc string `json:"short_desc"`
	Content	string `json:"content"`
	Thumbnail string `json:"thumbnail"`
	IsPublished bool `json:"is_published"`
}

type UpdateTransparansiInput struct {
	Title string `json:"title"`
	ShortDesc string `json:"short_desc"`
	Content	string `json:"content"`
	Thumbnail string `json:"thumbnail"`
	IsPublished bool `json:"is_published"`
}

type CreateInvestorInput struct {
	UserId string `json:"user_id"`
	CrowdfundId string `json:"crowdfund_id"`
	Fund int `json:"fund"`
}

type UpdateInvestorInput struct {
	Fund int `json:"fund"`
}

type CreateRequestKYCInput struct {
	UserId 		string `form:"user_id" json:"user_id"`
	Filepath	string `form:"file_path" json:"file_path"`
}

type RequestKYCFormData struct {
	UserId 		string `form:"user_id" json:"user_id"`
	File 		*multipart.FileHeader `form:"file" json:"file" binding:"required"`
}

type ApproveKYCInput struct {
	AdminId 	string `json:"admin_id"`
	RequestId	string `json:"request_id"`
}


type UpdateRequestKYCInput struct {
	AdminID string `json:"admin_id"`
	RejectReason	string `json:"reject_reason"`
	IsAccepted bool `json:"is_accepted"`
}
