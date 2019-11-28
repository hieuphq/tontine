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

// GroupLog is an object representing the database table.
type GroupLog struct {
	ID        int64        `boil:"id" json:"id" toml:"id" yaml:"id"`
	GroupID   null.Int64   `boil:"group_id" json:"group_id,omitempty" toml:"group_id" yaml:"group_id,omitempty"`
	Amount    null.Float64 `boil:"amount" json:"amount,omitempty" toml:"amount" yaml:"amount,omitempty"`
	Currency  null.String  `boil:"currency" json:"currency,omitempty" toml:"currency" yaml:"currency,omitempty"`
	CreatedAt time.Time    `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time    `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	DeletedAt null.Time    `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`

	R *groupLogR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L groupLogL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var GroupLogColumns = struct {
	ID        string
	GroupID   string
	Amount    string
	Currency  string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}{
	ID:        "id",
	GroupID:   "group_id",
	Amount:    "amount",
	Currency:  "currency",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// Generated where

var GroupLogWhere = struct {
	ID        whereHelperint64
	GroupID   whereHelpernull_Int64
	Amount    whereHelpernull_Float64
	Currency  whereHelpernull_String
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
	DeletedAt whereHelpernull_Time
}{
	ID:        whereHelperint64{field: "\"group_logs\".\"id\""},
	GroupID:   whereHelpernull_Int64{field: "\"group_logs\".\"group_id\""},
	Amount:    whereHelpernull_Float64{field: "\"group_logs\".\"amount\""},
	Currency:  whereHelpernull_String{field: "\"group_logs\".\"currency\""},
	CreatedAt: whereHelpertime_Time{field: "\"group_logs\".\"created_at\""},
	UpdatedAt: whereHelpertime_Time{field: "\"group_logs\".\"updated_at\""},
	DeletedAt: whereHelpernull_Time{field: "\"group_logs\".\"deleted_at\""},
}

// GroupLogRels is where relationship names are stored.
var GroupLogRels = struct {
	Group string
}{
	Group: "Group",
}

// groupLogR is where relationships are stored.
type groupLogR struct {
	Group *Group
}

// NewStruct creates a new relationship struct
func (*groupLogR) NewStruct() *groupLogR {
	return &groupLogR{}
}

// groupLogL is where Load methods for each relationship are stored.
type groupLogL struct{}

var (
	groupLogAllColumns            = []string{"id", "group_id", "amount", "currency", "created_at", "updated_at", "deleted_at"}
	groupLogColumnsWithoutDefault = []string{"group_id", "currency", "deleted_at"}
	groupLogColumnsWithDefault    = []string{"id", "amount", "created_at", "updated_at"}
	groupLogPrimaryKeyColumns     = []string{"id"}
)

type (
	// GroupLogSlice is an alias for a slice of pointers to GroupLog.
	// This should generally be used opposed to []GroupLog.
	GroupLogSlice []*GroupLog
	// GroupLogHook is the signature for custom GroupLog hook methods
	GroupLogHook func(context.Context, boil.ContextExecutor, *GroupLog) error

	groupLogQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	groupLogType                 = reflect.TypeOf(&GroupLog{})
	groupLogMapping              = queries.MakeStructMapping(groupLogType)
	groupLogPrimaryKeyMapping, _ = queries.BindMapping(groupLogType, groupLogMapping, groupLogPrimaryKeyColumns)
	groupLogInsertCacheMut       sync.RWMutex
	groupLogInsertCache          = make(map[string]insertCache)
	groupLogUpdateCacheMut       sync.RWMutex
	groupLogUpdateCache          = make(map[string]updateCache)
	groupLogUpsertCacheMut       sync.RWMutex
	groupLogUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var groupLogBeforeInsertHooks []GroupLogHook
var groupLogBeforeUpdateHooks []GroupLogHook
var groupLogBeforeDeleteHooks []GroupLogHook
var groupLogBeforeUpsertHooks []GroupLogHook

var groupLogAfterInsertHooks []GroupLogHook
var groupLogAfterSelectHooks []GroupLogHook
var groupLogAfterUpdateHooks []GroupLogHook
var groupLogAfterDeleteHooks []GroupLogHook
var groupLogAfterUpsertHooks []GroupLogHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *GroupLog) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range groupLogBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *GroupLog) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range groupLogBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *GroupLog) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range groupLogBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *GroupLog) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range groupLogBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *GroupLog) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range groupLogAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *GroupLog) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range groupLogAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *GroupLog) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range groupLogAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *GroupLog) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range groupLogAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *GroupLog) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range groupLogAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddGroupLogHook registers your hook function for all future operations.
func AddGroupLogHook(hookPoint boil.HookPoint, groupLogHook GroupLogHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		groupLogBeforeInsertHooks = append(groupLogBeforeInsertHooks, groupLogHook)
	case boil.BeforeUpdateHook:
		groupLogBeforeUpdateHooks = append(groupLogBeforeUpdateHooks, groupLogHook)
	case boil.BeforeDeleteHook:
		groupLogBeforeDeleteHooks = append(groupLogBeforeDeleteHooks, groupLogHook)
	case boil.BeforeUpsertHook:
		groupLogBeforeUpsertHooks = append(groupLogBeforeUpsertHooks, groupLogHook)
	case boil.AfterInsertHook:
		groupLogAfterInsertHooks = append(groupLogAfterInsertHooks, groupLogHook)
	case boil.AfterSelectHook:
		groupLogAfterSelectHooks = append(groupLogAfterSelectHooks, groupLogHook)
	case boil.AfterUpdateHook:
		groupLogAfterUpdateHooks = append(groupLogAfterUpdateHooks, groupLogHook)
	case boil.AfterDeleteHook:
		groupLogAfterDeleteHooks = append(groupLogAfterDeleteHooks, groupLogHook)
	case boil.AfterUpsertHook:
		groupLogAfterUpsertHooks = append(groupLogAfterUpsertHooks, groupLogHook)
	}
}

// One returns a single groupLog record from the query.
func (q groupLogQuery) One(ctx context.Context, exec boil.ContextExecutor) (*GroupLog, error) {
	o := &GroupLog{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dbmodel: failed to execute a one query for group_logs")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all GroupLog records from the query.
func (q groupLogQuery) All(ctx context.Context, exec boil.ContextExecutor) (GroupLogSlice, error) {
	var o []*GroupLog

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "dbmodel: failed to assign all query results to GroupLog slice")
	}

	if len(groupLogAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all GroupLog records in the query.
func (q groupLogQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: failed to count group_logs rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q groupLogQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "dbmodel: failed to check if group_logs exists")
	}

	return count > 0, nil
}

// Group pointed to by the foreign key.
func (o *GroupLog) Group(mods ...qm.QueryMod) groupQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.GroupID),
	}

	queryMods = append(queryMods, mods...)

	query := Groups(queryMods...)
	queries.SetFrom(query.Query, "\"groups\"")

	return query
}

// LoadGroup allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (groupLogL) LoadGroup(ctx context.Context, e boil.ContextExecutor, singular bool, maybeGroupLog interface{}, mods queries.Applicator) error {
	var slice []*GroupLog
	var object *GroupLog

	if singular {
		object = maybeGroupLog.(*GroupLog)
	} else {
		slice = *maybeGroupLog.(*[]*GroupLog)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &groupLogR{}
		}
		if !queries.IsNil(object.GroupID) {
			args = append(args, object.GroupID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &groupLogR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.GroupID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.GroupID) {
				args = append(args, obj.GroupID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`groups`), qm.WhereIn(`groups.id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Group")
	}

	var resultSlice []*Group
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Group")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for groups")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for groups")
	}

	if len(groupLogAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Group = foreign
		if foreign.R == nil {
			foreign.R = &groupR{}
		}
		foreign.R.GroupLogs = append(foreign.R.GroupLogs, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.GroupID, foreign.ID) {
				local.R.Group = foreign
				if foreign.R == nil {
					foreign.R = &groupR{}
				}
				foreign.R.GroupLogs = append(foreign.R.GroupLogs, local)
				break
			}
		}
	}

	return nil
}

// SetGroup of the groupLog to the related item.
// Sets o.R.Group to related.
// Adds o to related.R.GroupLogs.
func (o *GroupLog) SetGroup(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Group) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"group_logs\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, []string{"group_id"}),
		strmangle.WhereClause("\"", "\"", 0, groupLogPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.GroupID, related.ID)
	if o.R == nil {
		o.R = &groupLogR{
			Group: related,
		}
	} else {
		o.R.Group = related
	}

	if related.R == nil {
		related.R = &groupR{
			GroupLogs: GroupLogSlice{o},
		}
	} else {
		related.R.GroupLogs = append(related.R.GroupLogs, o)
	}

	return nil
}

// RemoveGroup relationship.
// Sets o.R.Group to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *GroupLog) RemoveGroup(ctx context.Context, exec boil.ContextExecutor, related *Group) error {
	var err error

	queries.SetScanner(&o.GroupID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("group_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.Group = nil
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.GroupLogs {
		if queries.Equal(o.GroupID, ri.GroupID) {
			continue
		}

		ln := len(related.R.GroupLogs)
		if ln > 1 && i < ln-1 {
			related.R.GroupLogs[i] = related.R.GroupLogs[ln-1]
		}
		related.R.GroupLogs = related.R.GroupLogs[:ln-1]
		break
	}
	return nil
}

// GroupLogs retrieves all the records using an executor.
func GroupLogs(mods ...qm.QueryMod) groupLogQuery {
	mods = append(mods, qm.From("\"group_logs\""))
	return groupLogQuery{NewQuery(mods...)}
}

// FindGroupLog retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindGroupLog(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*GroupLog, error) {
	groupLogObj := &GroupLog{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"group_logs\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, groupLogObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dbmodel: unable to select from group_logs")
	}

	return groupLogObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *GroupLog) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("dbmodel: no group_logs provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(groupLogColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	groupLogInsertCacheMut.RLock()
	cache, cached := groupLogInsertCache[key]
	groupLogInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			groupLogAllColumns,
			groupLogColumnsWithDefault,
			groupLogColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(groupLogType, groupLogMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(groupLogType, groupLogMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"group_logs\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"group_logs\" () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT \"%s\" FROM \"group_logs\" WHERE %s", strings.Join(returnColumns, "\",\""), strmangle.WhereClause("\"", "\"", 0, groupLogPrimaryKeyColumns))
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
		return errors.Wrap(err, "dbmodel: unable to insert into group_logs")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == groupLogMapping["id"] {
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
		return errors.Wrap(err, "dbmodel: unable to populate default values for group_logs")
	}

CacheNoHooks:
	if !cached {
		groupLogInsertCacheMut.Lock()
		groupLogInsertCache[key] = cache
		groupLogInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the GroupLog.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *GroupLog) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	groupLogUpdateCacheMut.RLock()
	cache, cached := groupLogUpdateCache[key]
	groupLogUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			groupLogAllColumns,
			groupLogPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("dbmodel: unable to update group_logs, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"group_logs\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, groupLogPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(groupLogType, groupLogMapping, append(wl, groupLogPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "dbmodel: unable to update group_logs row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: failed to get rows affected by update for group_logs")
	}

	if !cached {
		groupLogUpdateCacheMut.Lock()
		groupLogUpdateCache[key] = cache
		groupLogUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q groupLogQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: unable to update all for group_logs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: unable to retrieve rows affected for group_logs")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o GroupLogSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), groupLogPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"group_logs\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, groupLogPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: unable to update all in groupLog slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: unable to retrieve rows affected all in update all groupLog")
	}
	return rowsAff, nil
}

// Delete deletes a single GroupLog record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *GroupLog) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("dbmodel: no GroupLog provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), groupLogPrimaryKeyMapping)
	sql := "DELETE FROM \"group_logs\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: unable to delete from group_logs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: failed to get rows affected by delete for group_logs")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q groupLogQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("dbmodel: no groupLogQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: unable to delete all from group_logs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: failed to get rows affected by deleteall for group_logs")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o GroupLogSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(groupLogBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), groupLogPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"group_logs\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, groupLogPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: unable to delete all from groupLog slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: failed to get rows affected by deleteall for group_logs")
	}

	if len(groupLogAfterDeleteHooks) != 0 {
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
func (o *GroupLog) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindGroupLog(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *GroupLogSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := GroupLogSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), groupLogPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"group_logs\".* FROM \"group_logs\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, groupLogPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "dbmodel: unable to reload all in GroupLogSlice")
	}

	*o = slice

	return nil
}

// GroupLogExists checks if the GroupLog row exists.
func GroupLogExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"group_logs\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "dbmodel: unable to check if group_logs exists")
	}

	return exists, nil
}