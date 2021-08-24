package routers

type IndexRouter struct{}

func (idx IndexRouter) LoadRoutes() []RouteGroup {
	user := new (User)
	crowdfund := new (Crowdfund)
	admin := new (Admin)
	transparansi := new (Transparansi)
	faq := new (FAQ)
	investor := new(Investor)
	kyc := new (RequestKYC)
	return []RouteGroup{
		user,
		crowdfund,
		admin,
		transparansi,
		faq,
		investor,
		kyc,
	}
}