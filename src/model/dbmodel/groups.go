// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package dbmodel

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Group is an object representing the database table.
type Group struct {
	ID              int64        `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name            string       `boil:"name" json:"name" toml:"name" yaml:"name"`
	StrategyPercent null.Float64 `boil:"strategy_percent" json:"strategy_percent,omitempty" toml:"strategy_percent" yaml:"strategy_percent,omitempty"`
	Amount          float64      `boil:"amount" json:"amount" toml:"amount" yaml:"amount"`
	Currency        string       `boil:"currency" json:"currency" toml:"currency" yaml:"currency"`
	CreatedAt       time.Time    `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt       time.Time    `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	DeletedAt       null.Time    `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`

	R *groupR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L groupL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var GroupColumns = struct {
	ID              string
	Name            string
	StrategyPercent string
	Amount          string
	Currency        string
	CreatedAt       string
	UpdatedAt       string
	DeletedAt       string
}{
	ID:              "id",
	Name:            "name",
	StrategyPercent: "strategy_percent",
	Amount:          "amount",
	Currency:        "currency",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
	DeletedAt:       "deleted_at",
}

// Generated where

var GroupWhere = struct {
	ID              whereHelperint64
	Name            whereHelperstring
	StrategyPercent whereHelpernull_Float64
	Amount          whereHelperfloat64
	Currency        whereHelperstring
	CreatedAt       whereHelpertime_Time
	UpdatedAt       whereHelpertime_Time
	DeletedAt       whereHelpernull_Time
}{
	ID:              whereHelperint64{field: "\"groups\".\"id\""},
	Name:            whereHelperstring{field: "\"groups\".\"name\""},
	StrategyPercent: whereHelpernull_Float64{field: "\"groups\".\"strategy_percent\""},
	Amount:          whereHelperfloat64{field: "\"groups\".\"amount\""},
	Currency:        whereHelperstring{field: "\"groups\".\"currency\""},
	CreatedAt:       whereHelpertime_Time{field: "\"groups\".\"created_at\""},
	UpdatedAt:       whereHelpertime_Time{field: "\"groups\".\"updated_at\""},
	DeletedAt:       whereHelpernull_Time{field: "\"groups\".\"deleted_at\""},
}

// GroupRels is where relationship names are stored.
var GroupRels = struct {
	ActivityLogs    string
	GroupLogs       string
	GroupsInvestors string
}{
	ActivityLogs:    "ActivityLogs",
	GroupLogs:       "GroupLogs",
	GroupsInvestors: "GroupsInvestors",
}

// groupR is where relationships are stored.
type groupR struct {
	ActivityLogs    ActivityLogSlice
	GroupLogs       GroupLogSlice
	GroupsInvestors GroupsInvestorSlice
}

// NewStruct creates a new relationship struct
func (*groupR) NewStruct() *groupR {
	return &groupR{}
}

// groupL is where Load methods for each relationship are stored.
type groupL struct{}

var (
	groupAllColumns            = []string{"id", "name", "strategy_percent", "amount", "currency", "created_at", "updated_at", "deleted_at"}
	groupColumnsWithoutDefault = []string{}
	groupColumnsWithDefault    = []string{"id", "name", "strategy_percent", "amount", "currency", "created_at", "updated_at", "deleted_at"}
	groupPrimaryKeyColumns     = []string{"id"}
)

type (
	// GroupSlice is an alias for a slice of pointers to Group.
	// This should generally be used opposed to []Group.
	GroupSlice []*Group
	// GroupHook is the signature for custom Group hook methods
	GroupHook func(context.Context, boil.ContextExecutor, *Group) error

	groupQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	groupType                 = reflect.TypeOf(&Group{})
	groupMapping              = queries.MakeStructMapping(groupType)
	groupPrimaryKeyMapping, _ = queries.BindMapping(groupType, groupMapping, groupPrimaryKeyColumns)
	groupInsertCacheMut       sync.RWMutex
	groupInsertCache          = make(map[string]insertCache)
	groupUpdateCacheMut       sync.RWMutex
	groupUpdateCache          = make(map[string]updateCache)
	groupUpsertCacheMut       sync.RWMutex
	groupUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var groupBeforeInsertHooks []GroupHook
var groupBeforeUpdateHooks []GroupHook
var groupBeforeDeleteHooks []GroupHook
var groupBeforeUpsertHooks []GroupHook

var groupAfterInsertHooks []GroupHook
var groupAfterSelectHooks []GroupHook
var groupAfterUpdateHooks []GroupHook
var groupAfterDeleteHooks []GroupHook
var groupAfterUpsertHooks []GroupHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Group) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range groupBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Group) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range groupBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Group) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range groupBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Group) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range groupBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Group) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range groupAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Group) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range groupAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Group) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range groupAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Group) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range groupAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Group) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range groupAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddGroupHook registers your hook function for all future operations.
func AddGroupHook(hookPoint boil.HookPoint, groupHook GroupHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		groupBeforeInsertHooks = append(groupBeforeInsertHooks, groupHook)
	case boil.BeforeUpdateHook:
		groupBeforeUpdateHooks = append(groupBeforeUpdateHooks, groupHook)
	case boil.BeforeDeleteHook:
		groupBeforeDeleteHooks = append(groupBeforeDeleteHooks, groupHook)
	case boil.BeforeUpsertHook:
		groupBeforeUpsertHooks = append(groupBeforeUpsertHooks, groupHook)
	case boil.AfterInsertHook:
		groupAfterInsertHooks = append(groupAfterInsertHooks, groupHook)
	case boil.AfterSelectHook:
		groupAfterSelectHooks = append(groupAfterSelectHooks, groupHook)
	case boil.AfterUpdateHook:
		groupAfterUpdateHooks = append(groupAfterUpdateHooks, groupHook)
	case boil.AfterDeleteHook:
		groupAfterDeleteHooks = append(groupAfterDeleteHooks, groupHook)
	case boil.AfterUpsertHook:
		groupAfterUpsertHooks = append(groupAfterUpsertHooks, groupHook)
	}
}

// One returns a single group record from the query.
func (q groupQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Group, error) {
	o := &Group{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dbmodel: failed to execute a one query for groups")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Group records from the query.
func (q groupQuery) All(ctx context.Context, exec boil.ContextExecutor) (GroupSlice, error) {
	var o []*Group

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "dbmodel: failed to assign all query results to Group slice")
	}

	if len(groupAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Group records in the query.
func (q groupQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: failed to count groups rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q groupQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "dbmodel: failed to check if groups exists")
	}

	return count > 0, nil
}

// ActivityLogs retrieves all the activity_log's ActivityLogs with an executor.
func (o *Group) ActivityLogs(mods ...qm.QueryMod) activityLogQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"activity_logs\".\"group_id\"=?", o.ID),
	)

	query := ActivityLogs(queryMods...)
	queries.SetFrom(query.Query, "\"activity_logs\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"activity_logs\".*"})
	}

	return query
}

// GroupLogs retrieves all the group_log's GroupLogs with an executor.
func (o *Group) GroupLogs(mods ...qm.QueryMod) groupLogQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"group_logs\".\"group_id\"=?", o.ID),
	)

	query := GroupLogs(queryMods...)
	queries.SetFrom(query.Query, "\"group_logs\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"group_logs\".*"})
	}

	return query
}

// GroupsInvestors retrieves all the groups_investor's GroupsInvestors with an executor.
func (o *Group) GroupsInvestors(mods ...qm.QueryMod) groupsInvestorQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"groups_investors\".\"group_id\"=?", o.ID),
	)

	query := GroupsInvestors(queryMods...)
	queries.SetFrom(query.Query, "\"groups_investors\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"groups_investors\".*"})
	}

	return query
}

// LoadActivityLogs allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (groupL) LoadActivityLogs(ctx context.Context, e boil.ContextExecutor, singular bool, maybeGroup interface{}, mods queries.Applicator) error {
	var slice []*Group
	var object *Group

	if singular {
		object = maybeGroup.(*Group)
	} else {
		slice = *maybeGroup.(*[]*Group)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &groupR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &groupR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ID) {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`activity_logs`), qm.WhereIn(`activity_logs.group_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load activity_logs")
	}

	var resultSlice []*ActivityLog
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice activity_logs")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on activity_logs")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for activity_logs")
	}

	if len(activityLogAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.ActivityLogs = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &activityLogR{}
			}
			foreign.R.Group = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.ID, foreign.GroupID) {
				local.R.ActivityLogs = append(local.R.ActivityLogs, foreign)
				if foreign.R == nil {
					foreign.R = &activityLogR{}
				}
				foreign.R.Group = local
				break
			}
		}
	}

	return nil
}

// LoadGroupLogs allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (groupL) LoadGroupLogs(ctx context.Context, e boil.ContextExecutor, singular bool, maybeGroup interface{}, mods queries.Applicator) error {
	var slice []*Group
	var object *Group

	if singular {
		object = maybeGroup.(*Group)
	} else {
		slice = *maybeGroup.(*[]*Group)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &groupR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &groupR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`group_logs`), qm.WhereIn(`group_logs.group_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load group_logs")
	}

	var resultSlice []*GroupLog
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice group_logs")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on group_logs")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for group_logs")
	}

	if len(groupLogAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.GroupLogs = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &groupLogR{}
			}
			foreign.R.Group = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.GroupID {
				local.R.GroupLogs = append(local.R.GroupLogs, foreign)
				if foreign.R == nil {
					foreign.R = &groupLogR{}
				}
				foreign.R.Group = local
				break
			}
		}
	}

	return nil
}

// LoadGroupsInvestors allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (groupL) LoadGroupsInvestors(ctx context.Context, e boil.ContextExecutor, singular bool, maybeGroup interface{}, mods queries.Applicator) error {
	var slice []*Group
	var object *Group

	if singular {
		object = maybeGroup.(*Group)
	} else {
		slice = *maybeGroup.(*[]*Group)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &groupR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &groupR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`groups_investors`), qm.WhereIn(`groups_investors.group_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load groups_investors")
	}

	var resultSlice []*GroupsInvestor
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice groups_investors")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on groups_investors")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for groups_investors")
	}

	if len(groupsInvestorAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.GroupsInvestors = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &groupsInvestorR{}
			}
			foreign.R.Group = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.GroupID {
				local.R.GroupsInvestors = append(local.R.GroupsInvestors, foreign)
				if foreign.R == nil {
					foreign.R = &groupsInvestorR{}
				}
				foreign.R.Group = local
				break
			}
		}
	}

	return nil
}

// AddActivityLogs adds the given related objects to the existing relationships
// of the group, optionally inserting them as new records.
// Appends related to o.R.ActivityLogs.
// Sets related.R.Group appropriately.
func (o *Group) AddActivityLogs(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*ActivityLog) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.GroupID, o.ID)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"activity_logs\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 0, []string{"group_id"}),
				strmangle.WhereClause("\"", "\"", 0, activityLogPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.GroupID, o.ID)
		}
	}

	if o.R == nil {
		o.R = &groupR{
			ActivityLogs: related,
		}
	} else {
		o.R.ActivityLogs = append(o.R.ActivityLogs, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &activityLogR{
				Group: o,
			}
		} else {
			rel.R.Group = o
		}
	}
	return nil
}

// SetActivityLogs removes all previously related items of the
// group replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Group's ActivityLogs accordingly.
// Replaces o.R.ActivityLogs with related.
// Sets related.R.Group's ActivityLogs accordingly.
func (o *Group) SetActivityLogs(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*ActivityLog) error {
	query := "update \"activity_logs\" set \"group_id\" = null where \"group_id\" = ?"
	values := []interface{}{o.ID}
	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, query)
		fmt.Fprintln(writer, values)
	}
	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.ActivityLogs {
			queries.SetScanner(&rel.GroupID, nil)
			if rel.R == nil {
				continue
			}

			rel.R.Group = nil
		}

		o.R.ActivityLogs = nil
	}
	return o.AddActivityLogs(ctx, exec, insert, related...)
}

// RemoveActivityLogs relationships from objects passed in.
// Removes related items from R.ActivityLogs (uses pointer comparison, removal does not keep order)
// Sets related.R.Group.
func (o *Group) RemoveActivityLogs(ctx context.Context, exec boil.ContextExecutor, related ...*ActivityLog) error {
	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.GroupID, nil)
		if rel.R != nil {
			rel.R.Group = nil
		}
		if _, err = rel.Update(ctx, exec, boil.Whitelist("group_id")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.ActivityLogs {
			if rel != ri {
				continue
			}

			ln := len(o.R.ActivityLogs)
			if ln > 1 && i < ln-1 {
				o.R.ActivityLogs[i] = o.R.ActivityLogs[ln-1]
			}
			o.R.ActivityLogs = o.R.ActivityLogs[:ln-1]
			break
		}
	}

	return nil
}

// AddGroupLogs adds the given related objects to the existing relationships
// of the group, optionally inserting them as new records.
// Appends related to o.R.GroupLogs.
// Sets related.R.Group appropriately.
func (o *Group) AddGroupLogs(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*GroupLog) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.GroupID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"group_logs\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 0, []string{"group_id"}),
				strmangle.WhereClause("\"", "\"", 0, groupLogPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.GroupID = o.ID
		}
	}

	if o.R == nil {
		o.R = &groupR{
			GroupLogs: related,
		}
	} else {
		o.R.GroupLogs = append(o.R.GroupLogs, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &groupLogR{
				Group: o,
			}
		} else {
			rel.R.Group = o
		}
	}
	return nil
}

// AddGroupsInvestors adds the given related objects to the existing relationships
// of the group, optionally inserting them as new records.
// Appends related to o.R.GroupsInvestors.
// Sets related.R.Group appropriately.
func (o *Group) AddGroupsInvestors(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*GroupsInvestor) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.GroupID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"groups_investors\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 0, []string{"group_id"}),
				strmangle.WhereClause("\"", "\"", 0, groupsInvestorPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.GroupID = o.ID
		}
	}

	if o.R == nil {
		o.R = &groupR{
			GroupsInvestors: related,
		}
	} else {
		o.R.GroupsInvestors = append(o.R.GroupsInvestors, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &groupsInvestorR{
				Group: o,
			}
		} else {
			rel.R.Group = o
		}
	}
	return nil
}

// Groups retrieves all the records using an executor.
func Groups(mods ...qm.QueryMod) groupQuery {
	mods = append(mods, qm.From("\"groups\""))
	return groupQuery{NewQuery(mods...)}
}

// FindGroup retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindGroup(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Group, error) {
	groupObj := &Group{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"groups\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, groupObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dbmodel: unable to select from groups")
	}

	return groupObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Group) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("dbmodel: no groups provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(groupColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	groupInsertCacheMut.RLock()
	cache, cached := groupInsertCache[key]
	groupInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			groupAllColumns,
			groupColumnsWithDefault,
			groupColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(groupType, groupMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(groupType, groupMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"groups\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"groups\" () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT \"%s\" FROM \"groups\" WHERE %s", strings.Join(returnColumns, "\",\""), strmangle.WhereClause("\"", "\"", 0, groupPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "dbmodel: unable to insert into groups")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == groupMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "dbmodel: unable to populate default values for groups")
	}

CacheNoHooks:
	if !cached {
		groupInsertCacheMut.Lock()
		groupInsertCache[key] = cache
		groupInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Group.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Group) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	groupUpdateCacheMut.RLock()
	cache, cached := groupUpdateCache[key]
	groupUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			groupAllColumns,
			groupPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("dbmodel: unable to update groups, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"groups\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, groupPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(groupType, groupMapping, append(wl, groupPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: unable to update groups row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: failed to get rows affected by update for groups")
	}

	if !cached {
		groupUpdateCacheMut.Lock()
		groupUpdateCache[key] = cache
		groupUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q groupQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: unable to update all for groups")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: unable to retrieve rows affected for groups")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o GroupSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("dbmodel: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), groupPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"groups\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, groupPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: unable to update all in group slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: unable to retrieve rows affected all in update all group")
	}
	return rowsAff, nil
}

// Delete deletes a single Group record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Group) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("dbmodel: no Group provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), groupPrimaryKeyMapping)
	sql := "DELETE FROM \"groups\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: unable to delete from groups")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: failed to get rows affected by delete for groups")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q groupQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("dbmodel: no groupQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: unable to delete all from groups")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: failed to get rows affected by deleteall for groups")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o GroupSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(groupBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), groupPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"groups\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, groupPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: unable to delete all from group slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: failed to get rows affected by deleteall for groups")
	}

	if len(groupAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Group) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindGroup(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *GroupSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := GroupSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), groupPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"groups\".* FROM \"groups\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, groupPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "dbmodel: unable to reload all in GroupSlice")
	}

	*o = slice

	return nil
}

// GroupExists checks if the Group row exists.
func GroupExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"groups\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "dbmodel: unable to check if groups exists")
	}

	return exists, nil
}
