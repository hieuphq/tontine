package model

// Group tontine group
type Group struct {
	Base
	Name            string
	StrategyPercent float64
	Investors       []Investor
}
