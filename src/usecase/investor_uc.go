package usecase

import (
	"github.com/hieuphq/tontine/src/model"
)

// InvestorUC group usecase
type InvestorUC interface {
	CreateInvestor(g model.Investor) (*model.Investor, error)
	UpdateInvestor(g model.Investor) (*model.Investor, error)
	GetInvestorByID(invtID int64) (*model.Investor, error)
	GetInvestorList() ([]model.Investor, error)
}
