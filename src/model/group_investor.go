package model

// GroupInvestor ..
type GroupInvestor struct {
	Base
	GroupID    int64
	InvestorID int64
	Amount     float64
	Currency   string
	Group      *Group
	Investor   *Investor
}
