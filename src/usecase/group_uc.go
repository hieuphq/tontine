package usecase

import (
	"github.com/hieuphq/tontine/src/model"
)

// GroupUC group usecase
type GroupUC interface {
	CreateGroup(g model.Group) (*model.Group, error)
	UpdateGroup(g model.Group) (*model.Group, error)
	GetGroupList(g model.Group) (*model.Group, error)
	GetGroupDetail(g model.Group) (*model.Group, error)
	AddInvestorIntoGroup(gID int64, invt model.Investor) (*model.GroupBalance, error)
	RemoveInvestorFromGroup(gID int64, invt model.Investor) (*model.GroupBalance, error)
	WithdrawProfit(gID int64, amount float64) (*model.WithdrawResult, error)
	UpdateBalance(gID int64, amount float64) (*model.GroupBalance, error)
	Close(g model.Group) (*model.Group, error)
}

// type implGroup struct{}

// func NewGroupUC()
