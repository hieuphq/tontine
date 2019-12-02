package model

// Investor contains user info
type Investor struct {
	Base
	Name string `json:"name,omitempty"`
}
