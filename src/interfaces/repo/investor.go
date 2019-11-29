package repo

import (
	"context"

	"github.com/hieuphq/tontine/src/interfaces/store"
	"github.com/hieuphq/tontine/src/model"
)

// Investor interface for persistent data
type Investor interface {
	Create(ctx context.Context, store store.Store, g model.Investor) (*model.Investor, error)
	Update(ctx context.Context, store store.Store, g model.Investor) (*model.Investor, error)
	GetByID(ctx context.Context, store store.Store, ID int64) (*model.Investor, error)
	GetList(ctx context.Context, store store.Store) ([]model.Investor, error)
	Delete(ctx context.Context, store store.Store, g model.Investor) error
}
