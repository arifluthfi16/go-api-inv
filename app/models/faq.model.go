package models

import "gorm.io/gorm"

type FAQ struct {
	gorm.Model
	AdminId    	string
	Admin		Admin
	Pertanyaan 	string
	Jawaban    	string
	Count      	int
	Kategori   	string
}
