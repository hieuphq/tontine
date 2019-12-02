package repo

import (
	"context"

	"github.com/hieuphq/tontine/src/interfaces/store"
	"github.com/hieuphq/tontine/src/model"
)

// Group interface work with persistent data
type Group interface {
	Create(ctx context.Context, store store.Store, g model.Group) (*model.Group, error)
	Update(ctx context.Context, store store.Store, g model.Group) (*model.Group, error)
	GetByID(ctx context.Context, store store.Store, ID int64) (*model.Group, error)
	GetDetailByID(ctx context.Context, store store.Store, ID int64) (*model.GroupBalance, error)
	GetList(ctx context.Context, store store.Store) ([]model.Group, error)
	Delete(ctx context.Context, store store.Store, g model.Group) error
	AddInvestor(ctx context.Context, store store.Store, gi model.GroupInvestor) (*model.GroupInvestor, error)
	ExistedInvestor(ctx context.Context, store store.Store, groupID, invtID int64) (*model.GroupInvestor, error)
	InvestorList(ctx context.Context, store store.Store, groupID int64) ([]model.GroupInvestor, error)
	UpdateInvestor(ctx context.Context, store store.Store, gi model.GroupInvestor) (*model.GroupInvestor, error)
	FarawellInvestor(ctx context.Context, store store.Store, invtID int64, gID int64) error
}
