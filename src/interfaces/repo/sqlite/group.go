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
		Amount:          g.Amount,
		Currency:        g.Currency,
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
	dt.Amount = g.Amount
	dt.Currency = g.Currency

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
		Amount:          dt.Amount,
		Currency:        dt.Currency,
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

func (r *groupRepo) AddInvestor(ctx context.Context, store store.Store, gi model.GroupInvestor) (*model.GroupInvestor, error) {
	db := store.DB()
	dt := dbmodel.GroupsInvestor{
		GroupID:    gi.GroupID,
		InvestorID: gi.InvestorID,
		Amount:     gi.Amount,
		Currency:   gi.Currency,
	}
	err := dt.Insert(ctx, db, boil.Infer())
	if err != nil {
		return nil, err
	}

	gi.ID = dt.ID
	gi.CreatedAt = dt.CreatedAt
	gi.UpdatedAt = dt.UpdatedAt
	if dt.DeletedAt.Valid {
		v := dt.DeletedAt.Time
		gi.DeletedAt = &v
	}
	return &gi, nil
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
		Amount:     dt.Amount,
		Currency:   dt.Currency,
	}
}

func toGroupBalanceModel(dt *dbmodel.Group) model.GroupBalance {
	g := toModel(dt)

	dts := []model.InvestorBalance{}

	if len(dt.R.GroupsInvestors) > 0 {
		for idx := range dt.R.GroupsInvestors {
			itm := dt.R.GroupsInvestors[idx]

			dts = append(dts, model.InvestorBalance{
				InvestorID: itm.InvestorID,
				Amount:     itm.Amount,
				Currency:   itm.Currency,
			})
		}
	}

	return model.GroupBalance{
		Group:   &g,
		Details: dts,
	}
}

func (r *groupRepo) GetDetailByID(ctx context.Context, store store.Store, ID int64) (*model.GroupBalance, error) {
	db := store.DB()
	dt, err := dbmodel.
		Groups(
			qm.Where("id = ? AND deleted_at IS NULL", ID),
			qm.Load(dbmodel.GroupRels.GroupsInvestors, qm.Where("deleted_at IS NULL")),
		).
		One(ctx, db)
	if err != nil {
		return nil, err
	}

	rs := toGroupBalanceModel(dt)
	return &rs, nil
}

func (r *groupRepo) ExistedInvestor(ctx context.Context, store store.Store, groupID, invtID int64) (*model.GroupInvestor, error) {
	db := store.DB()
	dt, err := dbmodel.GroupsInvestors(
		qm.Where("group_id = ? AND investor_id = ? AND deleted_at IS NULL", groupID, invtID),
	).One(ctx, db)
	if err != nil {
		return nil, err
	}
	rs := toGroupInvestorModel(dt)
	return &rs, nil
}

func (r *groupRepo) UpdateInvestor(ctx context.Context, store store.Store, gi model.GroupInvestor) (*model.GroupInvestor, error) {
	db := store.DB()
	dt, err := dbmodel.
		GroupsInvestors(qm.Where("id = ? AND deleted_at IS NULL", gi.ID)).
		One(ctx, db)

	if err != nil {
		return nil, err
	}

	dt.Amount = gi.Amount
	dt.Currency = gi.Currency

	_, err = dt.Update(ctx, db, boil.Infer())
	if err != nil {
		return nil, err
	}

	return &gi, nil
}

func (r *groupRepo) InvestorList(ctx context.Context, store store.Store, groupID int64) ([]model.GroupInvestor, error) {
	db := store.DB()
	dt, err := dbmodel.GroupsInvestors(
		qm.Where("group_id = ? AND deleted_at IS NULL", groupID),
	).All(ctx, db)
	if err != nil {
		return nil, err
	}
	rs := []model.GroupInvestor{}
	for idx := range dt {
		itm := dt[idx]
		rs = append(rs, toGroupInvestorModel(itm))
	}
	return rs, nil
}
