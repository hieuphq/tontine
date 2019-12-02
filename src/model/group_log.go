package model

// GroupLog group's balance log
type GroupLog struct {
	Base
	Name     string  `json:"name,omitempty"`
	GroupID  int64   `json:"group_id,omitempty"`
	Amount   float64 `json:"amount,omitempty"`
	Currency string  `json:"currency,omitempty"`
}
