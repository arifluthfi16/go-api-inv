package models

type Model struct {
	Model interface{}
}

func RegisterModels() []Model {
	return []Model{
		{Model: Admin{}},
		{Model: User{}},
		{Model: Crowdfund{}},
		{Model: FAQ{}},
		{Model: Investor{}},
		{Model: Transparansi{}},
		{Model: RequestKYC{}},
	}
}

