package model

// ActivityLog investor's activities log
type ActivityLog struct {
	Base
	ID       int64
	GroupID  int64
	Action   InvestorAction
	Amount   float64
	Currency string
}
