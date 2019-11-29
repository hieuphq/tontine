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

func (r *groupRepo) Update(ctx context.Context, store store.Store, g model.Group) (*model.Group, error) {
	db := store.DB()
	dt, err := dbmodel.Groups(qm.Where("id = ? AND deleted_at IS NULL", g.ID)).One(ctx, db)
	if err != nil {
		return nil, err
	}

	dt.Name = g.Name
	dt.StrategyPercent = null.Float64{
		Float64: g.StrategyPercent,
		Valid:   true,
	}

	if g.DeletedAt != nil {
		dt.DeletedAt = null.NewTime(*g.DeletedAt, true)
	}

	_, err = dt.Update(ctx, db, boil.Infer())
	if err != nil {
		return nil, err
	}

	return &g, nil
}

func (r *groupRepo) GetByID(ctx context.Context, store store.Store, ID int64) (*model.Group, error) {
	db := store.DB()
	dt, err := dbmodel.Groups(qm.Where("id = ? AND deleted_at IS NULL", ID)).One(ctx, db)
	if err != nil {
		return nil, err
	}

	rs := toModel(dt)
	return &rs, nil
}

func (r *groupRepo) GetList(ctx context.Context, store store.Store) ([]model.Group, error) {
	db := store.DB()
	dt, err := dbmodel.Groups(qm.Where("deleted_at IS NULL")).All(ctx, db)
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

func (r *groupRepo) Delete(ctx context.Context, store store.Store, g model.Group) error {
	now := time.Now().UTC()
	g.DeletedAt = &now
	_, err := r.Update(ctx, store, g)
	return err
}

func (r *groupRepo) GetGroupInvestorByGroupAndInvestorID(ctx context.Context, store store.Store, invtID int64, gID int64) (*model.GroupInvestor, error) {
	db := store.DB()
	dt, err := dbmodel.
		GroupsInvestors(
			qm.Where("investor_id = ? AND group_id = ? AND deleted_at IS NULL", invtID, gID),
		).
		One(ctx, db)

	if err != nil {
		return nil, err
	}

	rs := toGroupInvestorModel(dt)
	return &rs, nil
}
func (r *groupRepo) AddInvestor(ctx context.Context, store store.Store, invtID int64, gID int64) error {
	db := store.DB()
	dt := dbmodel.GroupsInvestor{
		GroupID:    gID,
		InvestorID: invtID,
	}
	err := dt.Insert(ctx, db, boil.Infer())
	if err != nil {
		return err
	}

	return nil
}

func (r *groupRepo) FarawellInvestor(ctx context.Context, store store.Store, invtID int64, gID int64) error {
	db := store.DB()
	dt, err := dbmodel.
		GroupsInvestors(qm.Where("investor_id = ? AND group_id = ? AND deleted_at IS NULL", invtID, gID)).
		One(ctx, db)

	if err != nil {
		return err
	}

	dt.DeletedAt = null.NewTime(time.Now().UTC(), true)
	_, err = dt.Update(ctx, db, boil.Infer())
	if err != nil {
		return err
	}

	return nil
}

func toGroupInvestorModel(dt *dbmodel.GroupsInvestor) model.GroupInvestor {
	base := model.Base{
		ID:        dt.ID,
		CreatedAt: dt.CreatedAt,
		UpdatedAt: dt.UpdatedAt,
	}
	if dt.DeletedAt.Valid {
		base.DeletedAt = &dt.DeletedAt.Time
	}

	return model.GroupInvestor{
		Base:       base,
		GroupID:    dt.GroupID,
		InvestorID: dt.InvestorID,
	}
}
