package model

// GroupLogItem group log item for group log
type GroupLogItem struct {
	Base
	GroupLogID int64
	Amount     float64
	Currency   string
	InvestorID int64
}
