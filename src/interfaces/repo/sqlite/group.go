package sqlite

import (
	"context"

	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"

	"github.com/hieuphq/tontine/src/interfaces/repo"
	"github.com/hieuphq/tontine/src/interfaces/store"
	"github.com/hieuphq/tontine/src/model"
	"github.com/hieuphq/tontine/src/model/dbmodel"
)

// groupRepo interface work with persistent data
type groupRepo struct{}

// NewGroupRepo implementation for group repo
func newGroupRepo() repo.Group {
	return &groupRepo{}
}

func (r *groupRepo) Create(ctx context.Context, store store.Store, g model.Group) (*model.Group, error) {
	db := store.DB()
	dt := dbmodel.Group{
		Name:            g.Name,
		StrategyPercent: null.Float64{Float64: g.StrategyPercent, Valid: true},
	}
	err := dt.Insert(ctx, db, boil.Infer())
	if err != nil {
		return nil, err
	}

	g.ID = dt.ID
	g.CreatedAt = dt.CreatedAt
	g.UpdatedAt = dt.UpdatedAt
	if dt.DeletedAt.Valid {
		v := dt.DeletedAt.Time
		g.DeletedAt = &v
	}
	return &g, nil
}

// func (r *groupRepo) Update(ctx context.Context, store store.Store, g model.Group) (*model.Group, error) {

// }

// func (r *groupRepo) GetByID(ctx context.Context, store store.Store, ID int64) (*model.Group, error) {

// }

func (r *groupRepo) GetList(ctx context.Context, store store.Store) ([]model.Group, error) {
	dt, err := dbmodel.Groups().All(ctx, store.DB())
	if err != nil {
		return nil, err
	}

	var gs []model.Group
	for idx := range dt {
		itm := dt[idx]
		gs = append(gs, toModel(itm))
	}
	return gs, nil
}

func toModel(dt *dbmodel.Group) model.Group {
	base := model.Base{
		ID:        dt.ID,
		CreatedAt: dt.CreatedAt,
		UpdatedAt: dt.UpdatedAt,
	}
	if dt.DeletedAt.Valid {
		base.DeletedAt = &dt.DeletedAt.Time
	}

	return model.Group{
		Base:            base,
		Name:            dt.Name,
		StrategyPercent: dt.StrategyPercent.Float64,
	}
}

// func (r *groupRepo) Delete(ctx context.Context, store store.Store, g model.Group) error {

// }
