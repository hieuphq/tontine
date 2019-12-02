package model

// Group tontine group
type Group struct {
	Base
	Name            string  `json:"name,omitempty"`
	StrategyPercent float64 `json:"strategy_percent,omitempty"`
	Amount          float64 `json:"total,omitempty"`
	Currency        string  `json:"currency,omitempty"`
}
