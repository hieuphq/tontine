package sqlite

import (
	"context"

	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/hieuphq/tontine/src/interfaces/repo"
	"github.com/hieuphq/tontine/src/interfaces/store"
	"github.com/hieuphq/tontine/src/model"
	"github.com/hieuphq/tontine/src/model/dbmodel"
)

type activityLogRepo struct {
}

func newActivityLogRepo() repo.ActivityLog {
	return &activityLogRepo{}
}

func (r *activityLogRepo) LogGroup(ctx context.Context, store store.Store, g model.GroupLog) (*model.GroupLog, error) {
	db := store.DB()
	dt := dbmodel.GroupLog{
		GroupID:  g.GroupID,
		Amount:   g.Amount,
		Currency: g.Currency,
		Name:     g.Name,
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

func (r *activityLogRepo) LogActivity(ctx context.Context, store store.Store, g model.ActivityLog) (*model.ActivityLog, error) {
	db := store.DB()
	dt := dbmodel.ActivityLog{
		InvestorID: g.InvestorID,
		GroupID:    null.NewInt64(g.GroupID, g.GroupID > 0),
		Action:     string(g.Action),
		Amount:     null.NewFloat64(g.Amount, true),
		Currency:   null.NewString(g.Currency, true),
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

func (r *activityLogRepo) LogGroupItem(ctx context.Context, store store.Store, g model.GroupLogItem) (*model.GroupLogItem, error) {
	db := store.DB()
	dt := dbmodel.GroupLogItem{
		GroupLogID: g.GroupLogID,
		Amount:     g.Amount,
		Currency:   g.Currency,
		InvestorID: null.NewInt64(g.InvestorID, g.InvestorID > 0),
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

func (r *activityLogRepo) GetGroupLogs(ctx context.Context, store store.Store, g model.Group) ([]model.GroupLog, error) {
	db := store.DB()

	dt, err := dbmodel.GroupLogs(qm.Where("deleted_at IS NULL")).All(ctx, db)
	if err != nil {
		return nil, err
	}

	var gls []model.GroupLog
	for idx := range dt {
		itm := dt[idx]
		gls = append(gls, toGroupLogModel(itm))
	}
	return gls, nil
}

func toGroupLogModel(dt *dbmodel.GroupLog) model.GroupLog {
	base := model.Base{
		ID:        dt.ID,
		CreatedAt: dt.CreatedAt,
		UpdatedAt: dt.UpdatedAt,
	}
	if dt.DeletedAt.Valid {
		base.DeletedAt = &dt.DeletedAt.Time
	}

	return model.GroupLog{
		Base:     base,
		GroupID:  dt.GroupID,
		Name:     dt.Name,
		Amount:   dt.Amount,
		Currency: dt.Currency,
	}
}

func (r *activityLogRepo) GetGroupLogItems(ctx context.Context, store store.Store, g *model.GroupLog) ([]model.GroupLogItem, error) {
	db := store.DB()

	query := []qm.QueryMod{
		qm.Where("deleted_at IS NULL AND group_log_id = ?", g.ID),
	}

	dt, err := dbmodel.GroupLogItems(query...).All(ctx, db)
	if err != nil {
		return nil, err
	}

	var gls []model.GroupLogItem
	for idx := range dt {
		itm := dt[idx]
		gls = append(gls, toGroupLogItemModel(itm))
	}
	return gls, nil
}

func (r *activityLogRepo) GetActivityLogs(ctx context.Context, store store.Store, invt model.Investor, g *model.Group) ([]model.ActivityLog, error) {
	db := store.DB()

	query := []qm.QueryMod{
		qm.Where("deleted_at IS NULL AND investor_id = ?", invt.ID),
	}

	if g != nil {
		query = append(query, qm.Where("group_id = ?", g.ID))
	}
	dt, err := dbmodel.ActivityLogs(query...).All(ctx, db)
	if err != nil {
		return nil, err
	}

	var gls []model.ActivityLog
	for idx := range dt {
		itm := dt[idx]
		gls = append(gls, toActivityLogModel(itm))
	}
	return gls, nil
}

func toGroupLogItemModel(dt *dbmodel.GroupLogItem) model.GroupLogItem {
	base := model.Base{
		ID:        dt.ID,
		CreatedAt: dt.CreatedAt,
		UpdatedAt: dt.UpdatedAt,
	}
	if dt.DeletedAt.Valid {
		base.DeletedAt = &dt.DeletedAt.Time
	}

	return model.GroupLogItem{
		Base:       base,
		GroupLogID: dt.GroupLogID,
		InvestorID: dt.InvestorID.Int64,
		Amount:     dt.Amount,
		Currency:   dt.Currency,
	}
}

func toActivityLogModel(dt *dbmodel.ActivityLog) model.ActivityLog {
	base := model.Base{
		ID:        dt.ID,
		CreatedAt: dt.CreatedAt,
		UpdatedAt: dt.UpdatedAt,
	}
	if dt.DeletedAt.Valid {
		base.DeletedAt = &dt.DeletedAt.Time
	}

	return model.ActivityLog{
		Base:       base,
		InvestorID: dt.InvestorID,
		GroupID:    dt.GroupID.Int64,
		Action:     model.InvestorAction(dt.Action),
		Amount:     dt.Amount.Float64,
		Currency:   dt.Currency.String,
	}
}
