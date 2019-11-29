package model

// GroupInvestor ..
type GroupInvestor struct {
	Base
	GroupID    int64
	InvestorID int64
	Group      *Group
	Investor   *Investor
}
