// Code generated by SQLBoiler 3.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// CorporationAllianceHistory is an object representing the database table.
type CorporationAllianceHistory struct {
	ID         uint64    `boil:"id" json:"id" toml:"id" yaml:"id"`
	RecordID   uint      `boil:"record_id" json:"record_id" toml:"record_id" yaml:"record_id"`
	AllianceID null.Uint `boil:"alliance_id" json:"alliance_id,omitempty" toml:"alliance_id" yaml:"alliance_id,omitempty"`
	StartDate  time.Time `boil:"start_date" json:"start_date" toml:"start_date" yaml:"start_date"`
	CreatedAt  time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt  time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *corporationAllianceHistoryR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L corporationAllianceHistoryL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CorporationAllianceHistoryColumns = struct {
	ID         string
	RecordID   string
	AllianceID string
	StartDate  string
	CreatedAt  string
	UpdatedAt  string
}{
	ID:         "id",
	RecordID:   "record_id",
	AllianceID: "alliance_id",
	StartDate:  "start_date",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// Generated where

type whereHelpernull_Uint struct{ field string }

func (w whereHelpernull_Uint) EQ(x null.Uint) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Uint) NEQ(x null.Uint) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Uint) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Uint) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Uint) LT(x null.Uint) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Uint) LTE(x null.Uint) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Uint) GT(x null.Uint) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Uint) GTE(x null.Uint) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var CorporationAllianceHistoryWhere = struct {
	ID         whereHelperuint64
	RecordID   whereHelperuint
	AllianceID whereHelpernull_Uint
	StartDate  whereHelpertime_Time
	CreatedAt  whereHelpertime_Time
	UpdatedAt  whereHelpertime_Time
}{
	ID:         whereHelperuint64{field: "`corporation_alliance_history`.`id`"},
	RecordID:   whereHelperuint{field: "`corporation_alliance_history`.`record_id`"},
	AllianceID: whereHelpernull_Uint{field: "`corporation_alliance_history`.`alliance_id`"},
	StartDate:  whereHelpertime_Time{field: "`corporation_alliance_history`.`start_date`"},
	CreatedAt:  whereHelpertime_Time{field: "`corporation_alliance_history`.`created_at`"},
	UpdatedAt:  whereHelpertime_Time{field: "`corporation_alliance_history`.`updated_at`"},
}

// CorporationAllianceHistoryRels is where relationship names are stored.
var CorporationAllianceHistoryRels = struct {
}{}

// corporationAllianceHistoryR is where relationships are stored.
type corporationAllianceHistoryR struct {
}

// NewStruct creates a new relationship struct
func (*corporationAllianceHistoryR) NewStruct() *corporationAllianceHistoryR {
	return &corporationAllianceHistoryR{}
}

// corporationAllianceHistoryL is where Load methods for each relationship are stored.
type corporationAllianceHistoryL struct{}

var (
	corporationAllianceHistoryAllColumns            = []string{"id", "record_id", "alliance_id", "start_date", "created_at", "updated_at"}
	corporationAllianceHistoryColumnsWithoutDefault = []string{"id", "record_id", "alliance_id", "start_date", "created_at", "updated_at"}
	corporationAllianceHistoryColumnsWithDefault    = []string{}
	corporationAllianceHistoryPrimaryKeyColumns     = []string{"id", "record_id"}
)

type (
	// CorporationAllianceHistorySlice is an alias for a slice of pointers to CorporationAllianceHistory.
	// This should generally be used opposed to []CorporationAllianceHistory.
	CorporationAllianceHistorySlice []*CorporationAllianceHistory
	// CorporationAllianceHistoryHook is the signature for custom CorporationAllianceHistory hook methods
	CorporationAllianceHistoryHook func(context.Context, boil.ContextExecutor, *CorporationAllianceHistory) error

	corporationAllianceHistoryQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	corporationAllianceHistoryType                 = reflect.TypeOf(&CorporationAllianceHistory{})
	corporationAllianceHistoryMapping              = queries.MakeStructMapping(corporationAllianceHistoryType)
	corporationAllianceHistoryPrimaryKeyMapping, _ = queries.BindMapping(corporationAllianceHistoryType, corporationAllianceHistoryMapping, corporationAllianceHistoryPrimaryKeyColumns)
	corporationAllianceHistoryInsertCacheMut       sync.RWMutex
	corporationAllianceHistoryInsertCache          = make(map[string]insertCache)
	corporationAllianceHistoryUpdateCacheMut       sync.RWMutex
	corporationAllianceHistoryUpdateCache          = make(map[string]updateCache)
	corporationAllianceHistoryUpsertCacheMut       sync.RWMutex
	corporationAllianceHistoryUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var corporationAllianceHistoryBeforeInsertHooks []CorporationAllianceHistoryHook
var corporationAllianceHistoryBeforeUpdateHooks []CorporationAllianceHistoryHook
var corporationAllianceHistoryBeforeDeleteHooks []CorporationAllianceHistoryHook
var corporationAllianceHistoryBeforeUpsertHooks []CorporationAllianceHistoryHook

var corporationAllianceHistoryAfterInsertHooks []CorporationAllianceHistoryHook
var corporationAllianceHistoryAfterSelectHooks []CorporationAllianceHistoryHook
var corporationAllianceHistoryAfterUpdateHooks []CorporationAllianceHistoryHook
var corporationAllianceHistoryAfterDeleteHooks []CorporationAllianceHistoryHook
var corporationAllianceHistoryAfterUpsertHooks []CorporationAllianceHistoryHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CorporationAllianceHistory) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range corporationAllianceHistoryBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CorporationAllianceHistory) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range corporationAllianceHistoryBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CorporationAllianceHistory) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range corporationAllianceHistoryBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CorporationAllianceHistory) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range corporationAllianceHistoryBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CorporationAllianceHistory) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range corporationAllianceHistoryAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CorporationAllianceHistory) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range corporationAllianceHistoryAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CorporationAllianceHistory) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range corporationAllianceHistoryAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CorporationAllianceHistory) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range corporationAllianceHistoryAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CorporationAllianceHistory) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range corporationAllianceHistoryAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCorporationAllianceHistoryHook registers your hook function for all future operations.
func AddCorporationAllianceHistoryHook(hookPoint boil.HookPoint, corporationAllianceHistoryHook CorporationAllianceHistoryHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		corporationAllianceHistoryBeforeInsertHooks = append(corporationAllianceHistoryBeforeInsertHooks, corporationAllianceHistoryHook)
	case boil.BeforeUpdateHook:
		corporationAllianceHistoryBeforeUpdateHooks = append(corporationAllianceHistoryBeforeUpdateHooks, corporationAllianceHistoryHook)
	case boil.BeforeDeleteHook:
		corporationAllianceHistoryBeforeDeleteHooks = append(corporationAllianceHistoryBeforeDeleteHooks, corporationAllianceHistoryHook)
	case boil.BeforeUpsertHook:
		corporationAllianceHistoryBeforeUpsertHooks = append(corporationAllianceHistoryBeforeUpsertHooks, corporationAllianceHistoryHook)
	case boil.AfterInsertHook:
		corporationAllianceHistoryAfterInsertHooks = append(corporationAllianceHistoryAfterInsertHooks, corporationAllianceHistoryHook)
	case boil.AfterSelectHook:
		corporationAllianceHistoryAfterSelectHooks = append(corporationAllianceHistoryAfterSelectHooks, corporationAllianceHistoryHook)
	case boil.AfterUpdateHook:
		corporationAllianceHistoryAfterUpdateHooks = append(corporationAllianceHistoryAfterUpdateHooks, corporationAllianceHistoryHook)
	case boil.AfterDeleteHook:
		corporationAllianceHistoryAfterDeleteHooks = append(corporationAllianceHistoryAfterDeleteHooks, corporationAllianceHistoryHook)
	case boil.AfterUpsertHook:
		corporationAllianceHistoryAfterUpsertHooks = append(corporationAllianceHistoryAfterUpsertHooks, corporationAllianceHistoryHook)
	}
}

// One returns a single corporationAllianceHistory record from the query.
func (q corporationAllianceHistoryQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CorporationAllianceHistory, error) {
	o := &CorporationAllianceHistory{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for corporation_alliance_history")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CorporationAllianceHistory records from the query.
func (q corporationAllianceHistoryQuery) All(ctx context.Context, exec boil.ContextExecutor) (CorporationAllianceHistorySlice, error) {
	var o []*CorporationAllianceHistory

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CorporationAllianceHistory slice")
	}

	if len(corporationAllianceHistoryAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CorporationAllianceHistory records in the query.
func (q corporationAllianceHistoryQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count corporation_alliance_history rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q corporationAllianceHistoryQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if corporation_alliance_history exists")
	}

	return count > 0, nil
}

// CorporationAllianceHistories retrieves all the records using an executor.
func CorporationAllianceHistories(mods ...qm.QueryMod) corporationAllianceHistoryQuery {
	mods = append(mods, qm.From("`corporation_alliance_history`"))
	return corporationAllianceHistoryQuery{NewQuery(mods...)}
}

// FindCorporationAllianceHistory retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCorporationAllianceHistory(ctx context.Context, exec boil.ContextExecutor, iD uint64, recordID uint, selectCols ...string) (*CorporationAllianceHistory, error) {
	corporationAllianceHistoryObj := &CorporationAllianceHistory{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `corporation_alliance_history` where `id`=? AND `record_id`=?", sel,
	)

	q := queries.Raw(query, iD, recordID)

	err := q.Bind(ctx, exec, corporationAllianceHistoryObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from corporation_alliance_history")
	}

	return corporationAllianceHistoryObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CorporationAllianceHistory) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no corporation_alliance_history provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(corporationAllianceHistoryColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	corporationAllianceHistoryInsertCacheMut.RLock()
	cache, cached := corporationAllianceHistoryInsertCache[key]
	corporationAllianceHistoryInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			corporationAllianceHistoryAllColumns,
			corporationAllianceHistoryColumnsWithDefault,
			corporationAllianceHistoryColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(corporationAllianceHistoryType, corporationAllianceHistoryMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(corporationAllianceHistoryType, corporationAllianceHistoryMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `corporation_alliance_history` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `corporation_alliance_history` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `corporation_alliance_history` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, corporationAllianceHistoryPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into corporation_alliance_history")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
		o.RecordID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for corporation_alliance_history")
	}

CacheNoHooks:
	if !cached {
		corporationAllianceHistoryInsertCacheMut.Lock()
		corporationAllianceHistoryInsertCache[key] = cache
		corporationAllianceHistoryInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CorporationAllianceHistory.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CorporationAllianceHistory) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	corporationAllianceHistoryUpdateCacheMut.RLock()
	cache, cached := corporationAllianceHistoryUpdateCache[key]
	corporationAllianceHistoryUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			corporationAllianceHistoryAllColumns,
			corporationAllianceHistoryPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update corporation_alliance_history, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `corporation_alliance_history` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, corporationAllianceHistoryPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(corporationAllianceHistoryType, corporationAllianceHistoryMapping, append(wl, corporationAllianceHistoryPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update corporation_alliance_history row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for corporation_alliance_history")
	}

	if !cached {
		corporationAllianceHistoryUpdateCacheMut.Lock()
		corporationAllianceHistoryUpdateCache[key] = cache
		corporationAllianceHistoryUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q corporationAllianceHistoryQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for corporation_alliance_history")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for corporation_alliance_history")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CorporationAllianceHistorySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), corporationAllianceHistoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `corporation_alliance_history` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, corporationAllianceHistoryPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in corporationAllianceHistory slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all corporationAllianceHistory")
	}
	return rowsAff, nil
}

var mySQLCorporationAllianceHistoryUniqueColumns = []string{}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CorporationAllianceHistory) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no corporation_alliance_history provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(corporationAllianceHistoryColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCorporationAllianceHistoryUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	corporationAllianceHistoryUpsertCacheMut.RLock()
	cache, cached := corporationAllianceHistoryUpsertCache[key]
	corporationAllianceHistoryUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			corporationAllianceHistoryAllColumns,
			corporationAllianceHistoryColumnsWithDefault,
			corporationAllianceHistoryColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			corporationAllianceHistoryAllColumns,
			corporationAllianceHistoryPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert corporation_alliance_history, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "corporation_alliance_history", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `corporation_alliance_history` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(corporationAllianceHistoryType, corporationAllianceHistoryMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(corporationAllianceHistoryType, corporationAllianceHistoryMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for corporation_alliance_history")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(corporationAllianceHistoryType, corporationAllianceHistoryMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for corporation_alliance_history")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for corporation_alliance_history")
	}

CacheNoHooks:
	if !cached {
		corporationAllianceHistoryUpsertCacheMut.Lock()
		corporationAllianceHistoryUpsertCache[key] = cache
		corporationAllianceHistoryUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CorporationAllianceHistory record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CorporationAllianceHistory) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CorporationAllianceHistory provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), corporationAllianceHistoryPrimaryKeyMapping)
	sql := "DELETE FROM `corporation_alliance_history` WHERE `id`=? AND `record_id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from corporation_alliance_history")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for corporation_alliance_history")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q corporationAllianceHistoryQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no corporationAllianceHistoryQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from corporation_alliance_history")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for corporation_alliance_history")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CorporationAllianceHistorySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(corporationAllianceHistoryBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), corporationAllianceHistoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `corporation_alliance_history` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, corporationAllianceHistoryPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from corporationAllianceHistory slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for corporation_alliance_history")
	}

	if len(corporationAllianceHistoryAfterDeleteHooks) != 0 {
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
func (o *CorporationAllianceHistory) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCorporationAllianceHistory(ctx, exec, o.ID, o.RecordID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CorporationAllianceHistorySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CorporationAllianceHistorySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), corporationAllianceHistoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `corporation_alliance_history`.* FROM `corporation_alliance_history` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, corporationAllianceHistoryPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CorporationAllianceHistorySlice")
	}

	*o = slice

	return nil
}

// CorporationAllianceHistoryExists checks if the CorporationAllianceHistory row exists.
func CorporationAllianceHistoryExists(ctx context.Context, exec boil.ContextExecutor, iD uint64, recordID uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `corporation_alliance_history` where `id`=? AND `record_id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD, recordID)
	}

	row := exec.QueryRowContext(ctx, sql, iD, recordID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if corporation_alliance_history exists")
	}

	return exists, nil
}