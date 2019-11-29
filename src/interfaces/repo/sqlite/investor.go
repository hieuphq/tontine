package sqlite

import (
	"context"
	"time"

	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/hieuphq/tontine/src/interfaces/repo"
	"github.com/hieuphq/tontine/src/interfaces/store"
	"github.com/hieuphq/tontine/src/model"
	"github.com/hieuphq/tontine/src/model/dbmodel"
)

// investorRepo interface work with persistent data
type investorRepo struct{}

// NewInvestorRepo implementation for investor repo
func newInvestorRepo() repo.Investor {
	return &investorRepo{}
}

func (r *investorRepo) Create(ctx context.Context, store store.Store, g model.Investor) (*model.Investor, error) {
	db := store.DB()
	dt := dbmodel.Investor{
		Name: g.Name,
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

func (r *investorRepo) Update(ctx context.Context, store store.Store, g model.Investor) (*model.Investor, error) {
	db := store.DB()
	dt, err := dbmodel.Investors(qm.Where("id = ? AND deleted_at IS NULL", g.ID)).One(ctx, db)
	if err != nil {
		return nil, err
	}

	dt.Name = g.Name

	if g.DeletedAt != nil {
		dt.DeletedAt = null.NewTime(*g.DeletedAt, true)
	}

	_, err = dt.Update(ctx, db, boil.Infer())
	if err != nil {
		return nil, err
	}

	return &g, nil
}

func (r *investorRepo) GetByID(ctx context.Context, store store.Store, ID int64) (*model.Investor, error) {
	db := store.DB()
	dt, err := dbmodel.Investors(qm.Where("id = ? AND deleted_at IS NULL", ID)).One(ctx, db)
	if err != nil {
		return nil, err
	}

	rs := toInvestorModel(dt)
	return &rs, nil
}

func (r *investorRepo) GetList(ctx context.Context, store store.Store) ([]model.Investor, error) {
	db := store.DB()
	dt, err := dbmodel.Investors(qm.Where("deleted_at IS NULL")).All(ctx, db)
	if err != nil {
		return nil, err
	}

	var gs []model.Investor
	for idx := range dt {
		itm := dt[idx]
		gs = append(gs, toInvestorModel(itm))
	}
	return gs, nil
}

func toInvestorModel(dt *dbmodel.Investor) model.Investor {
	base := model.Base{
		ID:        dt.ID,
		CreatedAt: dt.CreatedAt,
		UpdatedAt: dt.UpdatedAt,
	}
	if dt.DeletedAt.Valid {
		base.DeletedAt = &dt.DeletedAt.Time
	}

	return model.Investor{
		Base: base,
		Name: dt.Name,
	}
}

func (r *investorRepo) Delete(ctx context.Context, store store.Store, g model.Investor) error {
	now := time.Now().UTC()
	g.DeletedAt = &now
	_, err := r.Update(ctx, store, g)
	return err
}
