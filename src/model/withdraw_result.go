package model

// WithdrawResult withdraw result
type WithdrawResult struct {
	Group    *Group
	Balance  *GroupBalance
	Log      GroupLog
	LogItems []GroupLogItem
}

// GroupBalance group balance info
type GroupBalance struct {
	Group   *Group            `json:"group"`
	Details []InvestorBalance `json:"details"`
}

// InvestorBalance group balance info
type InvestorBalance struct {
	InvestorID int64
	Amount     float64
	Currency   string
}
