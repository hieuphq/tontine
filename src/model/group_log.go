package model

// GroupLog group's balance log
type GroupLog struct {
	Base
	Name     string
	GroupID  int64
	Amount   float64
	Currency string
}
