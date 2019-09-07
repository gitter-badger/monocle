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
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// CharacterCorporationHistory is an object representing the database table.
type CharacterCorporationHistory struct {
	ID            uint64    `boil:"id" json:"id" toml:"id" yaml:"id"`
	RecordID      uint      `boil:"record_id" json:"record_id" toml:"record_id" yaml:"record_id"`
	CorporationID uint      `boil:"corporation_id" json:"corporation_id" toml:"corporation_id" yaml:"corporation_id"`
	StartDate     time.Time `boil:"start_date" json:"start_date" toml:"start_date" yaml:"start_date"`
	CreatedAt     time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt     time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *characterCorporationHistoryR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L characterCorporationHistoryL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CharacterCorporationHistoryColumns = struct {
	ID            string
	RecordID      string
	CorporationID string
	StartDate     string
	CreatedAt     string
	UpdatedAt     string
}{
	ID:            "id",
	RecordID:      "record_id",
	CorporationID: "corporation_id",
	StartDate:     "start_date",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

// Generated where

var CharacterCorporationHistoryWhere = struct {
	ID            whereHelperuint64
	RecordID      whereHelperuint
	CorporationID whereHelperuint
	StartDate     whereHelpertime_Time
	CreatedAt     whereHelpertime_Time
	UpdatedAt     whereHelpertime_Time
}{
	ID:            whereHelperuint64{field: "`character_corporation_history`.`id`"},
	RecordID:      whereHelperuint{field: "`character_corporation_history`.`record_id`"},
	CorporationID: whereHelperuint{field: "`character_corporation_history`.`corporation_id`"},
	StartDate:     whereHelpertime_Time{field: "`character_corporation_history`.`start_date`"},
	CreatedAt:     whereHelpertime_Time{field: "`character_corporation_history`.`created_at`"},
	UpdatedAt:     whereHelpertime_Time{field: "`character_corporation_history`.`updated_at`"},
}

// CharacterCorporationHistoryRels is where relationship names are stored.
var CharacterCorporationHistoryRels = struct {
}{}

// characterCorporationHistoryR is where relationships are stored.
type characterCorporationHistoryR struct {
}

// NewStruct creates a new relationship struct
func (*characterCorporationHistoryR) NewStruct() *characterCorporationHistoryR {
	return &characterCorporationHistoryR{}
}

// characterCorporationHistoryL is where Load methods for each relationship are stored.
type characterCorporationHistoryL struct{}

var (
	characterCorporationHistoryAllColumns            = []string{"id", "record_id", "corporation_id", "start_date", "created_at", "updated_at"}
	characterCorporationHistoryColumnsWithoutDefault = []string{"id", "record_id", "corporation_id", "start_date", "created_at", "updated_at"}
	characterCorporationHistoryColumnsWithDefault    = []string{}
	characterCorporationHistoryPrimaryKeyColumns     = []string{"id", "record_id"}
)

type (
	// CharacterCorporationHistorySlice is an alias for a slice of pointers to CharacterCorporationHistory.
	// This should generally be used opposed to []CharacterCorporationHistory.
	CharacterCorporationHistorySlice []*CharacterCorporationHistory
	// CharacterCorporationHistoryHook is the signature for custom CharacterCorporationHistory hook methods
	CharacterCorporationHistoryHook func(context.Context, boil.ContextExecutor, *CharacterCorporationHistory) error

	characterCorporationHistoryQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	characterCorporationHistoryType                 = reflect.TypeOf(&CharacterCorporationHistory{})
	characterCorporationHistoryMapping              = queries.MakeStructMapping(characterCorporationHistoryType)
	characterCorporationHistoryPrimaryKeyMapping, _ = queries.BindMapping(characterCorporationHistoryType, characterCorporationHistoryMapping, characterCorporationHistoryPrimaryKeyColumns)
	characterCorporationHistoryInsertCacheMut       sync.RWMutex
	characterCorporationHistoryInsertCache          = make(map[string]insertCache)
	characterCorporationHistoryUpdateCacheMut       sync.RWMutex
	characterCorporationHistoryUpdateCache          = make(map[string]updateCache)
	characterCorporationHistoryUpsertCacheMut       sync.RWMutex
	characterCorporationHistoryUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var characterCorporationHistoryBeforeInsertHooks []CharacterCorporationHistoryHook
var characterCorporationHistoryBeforeUpdateHooks []CharacterCorporationHistoryHook
var characterCorporationHistoryBeforeDeleteHooks []CharacterCorporationHistoryHook
var characterCorporationHistoryBeforeUpsertHooks []CharacterCorporationHistoryHook

var characterCorporationHistoryAfterInsertHooks []CharacterCorporationHistoryHook
var characterCorporationHistoryAfterSelectHooks []CharacterCorporationHistoryHook
var characterCorporationHistoryAfterUpdateHooks []CharacterCorporationHistoryHook
var characterCorporationHistoryAfterDeleteHooks []CharacterCorporationHistoryHook
var characterCorporationHistoryAfterUpsertHooks []CharacterCorporationHistoryHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CharacterCorporationHistory) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range characterCorporationHistoryBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CharacterCorporationHistory) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range characterCorporationHistoryBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CharacterCorporationHistory) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range characterCorporationHistoryBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CharacterCorporationHistory) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range characterCorporationHistoryBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CharacterCorporationHistory) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range characterCorporationHistoryAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CharacterCorporationHistory) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range characterCorporationHistoryAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CharacterCorporationHistory) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range characterCorporationHistoryAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CharacterCorporationHistory) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range characterCorporationHistoryAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CharacterCorporationHistory) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range characterCorporationHistoryAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCharacterCorporationHistoryHook registers your hook function for all future operations.
func AddCharacterCorporationHistoryHook(hookPoint boil.HookPoint, characterCorporationHistoryHook CharacterCorporationHistoryHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		characterCorporationHistoryBeforeInsertHooks = append(characterCorporationHistoryBeforeInsertHooks, characterCorporationHistoryHook)
	case boil.BeforeUpdateHook:
		characterCorporationHistoryBeforeUpdateHooks = append(characterCorporationHistoryBeforeUpdateHooks, characterCorporationHistoryHook)
	case boil.BeforeDeleteHook:
		characterCorporationHistoryBeforeDeleteHooks = append(characterCorporationHistoryBeforeDeleteHooks, characterCorporationHistoryHook)
	case boil.BeforeUpsertHook:
		characterCorporationHistoryBeforeUpsertHooks = append(characterCorporationHistoryBeforeUpsertHooks, characterCorporationHistoryHook)
	case boil.AfterInsertHook:
		characterCorporationHistoryAfterInsertHooks = append(characterCorporationHistoryAfterInsertHooks, characterCorporationHistoryHook)
	case boil.AfterSelectHook:
		characterCorporationHistoryAfterSelectHooks = append(characterCorporationHistoryAfterSelectHooks, characterCorporationHistoryHook)
	case boil.AfterUpdateHook:
		characterCorporationHistoryAfterUpdateHooks = append(characterCorporationHistoryAfterUpdateHooks, characterCorporationHistoryHook)
	case boil.AfterDeleteHook:
		characterCorporationHistoryAfterDeleteHooks = append(characterCorporationHistoryAfterDeleteHooks, characterCorporationHistoryHook)
	case boil.AfterUpsertHook:
		characterCorporationHistoryAfterUpsertHooks = append(characterCorporationHistoryAfterUpsertHooks, characterCorporationHistoryHook)
	}
}

// One returns a single characterCorporationHistory record from the query.
func (q characterCorporationHistoryQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CharacterCorporationHistory, error) {
	o := &CharacterCorporationHistory{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for character_corporation_history")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CharacterCorporationHistory records from the query.
func (q characterCorporationHistoryQuery) All(ctx context.Context, exec boil.ContextExecutor) (CharacterCorporationHistorySlice, error) {
	var o []*CharacterCorporationHistory

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CharacterCorporationHistory slice")
	}

	if len(characterCorporationHistoryAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CharacterCorporationHistory records in the query.
func (q characterCorporationHistoryQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count character_corporation_history rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q characterCorporationHistoryQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if character_corporation_history exists")
	}

	return count > 0, nil
}

// CharacterCorporationHistories retrieves all the records using an executor.
func CharacterCorporationHistories(mods ...qm.QueryMod) characterCorporationHistoryQuery {
	mods = append(mods, qm.From("`character_corporation_history`"))
	return characterCorporationHistoryQuery{NewQuery(mods...)}
}

// FindCharacterCorporationHistory retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCharacterCorporationHistory(ctx context.Context, exec boil.ContextExecutor, iD uint64, recordID uint, selectCols ...string) (*CharacterCorporationHistory, error) {
	characterCorporationHistoryObj := &CharacterCorporationHistory{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `character_corporation_history` where `id`=? AND `record_id`=?", sel,
	)

	q := queries.Raw(query, iD, recordID)

	err := q.Bind(ctx, exec, characterCorporationHistoryObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from character_corporation_history")
	}

	return characterCorporationHistoryObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CharacterCorporationHistory) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no character_corporation_history provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(characterCorporationHistoryColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	characterCorporationHistoryInsertCacheMut.RLock()
	cache, cached := characterCorporationHistoryInsertCache[key]
	characterCorporationHistoryInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			characterCorporationHistoryAllColumns,
			characterCorporationHistoryColumnsWithDefault,
			characterCorporationHistoryColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(characterCorporationHistoryType, characterCorporationHistoryMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(characterCorporationHistoryType, characterCorporationHistoryMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `character_corporation_history` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `character_corporation_history` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `character_corporation_history` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, characterCorporationHistoryPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into character_corporation_history")
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
		return errors.Wrap(err, "models: unable to populate default values for character_corporation_history")
	}

CacheNoHooks:
	if !cached {
		characterCorporationHistoryInsertCacheMut.Lock()
		characterCorporationHistoryInsertCache[key] = cache
		characterCorporationHistoryInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CharacterCorporationHistory.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CharacterCorporationHistory) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	characterCorporationHistoryUpdateCacheMut.RLock()
	cache, cached := characterCorporationHistoryUpdateCache[key]
	characterCorporationHistoryUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			characterCorporationHistoryAllColumns,
			characterCorporationHistoryPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update character_corporation_history, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `character_corporation_history` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, characterCorporationHistoryPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(characterCorporationHistoryType, characterCorporationHistoryMapping, append(wl, characterCorporationHistoryPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update character_corporation_history row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for character_corporation_history")
	}

	if !cached {
		characterCorporationHistoryUpdateCacheMut.Lock()
		characterCorporationHistoryUpdateCache[key] = cache
		characterCorporationHistoryUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q characterCorporationHistoryQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for character_corporation_history")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for character_corporation_history")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CharacterCorporationHistorySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), characterCorporationHistoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `character_corporation_history` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, characterCorporationHistoryPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in characterCorporationHistory slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all characterCorporationHistory")
	}
	return rowsAff, nil
}

var mySQLCharacterCorporationHistoryUniqueColumns = []string{}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CharacterCorporationHistory) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no character_corporation_history provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(characterCorporationHistoryColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCharacterCorporationHistoryUniqueColumns, o)

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

	characterCorporationHistoryUpsertCacheMut.RLock()
	cache, cached := characterCorporationHistoryUpsertCache[key]
	characterCorporationHistoryUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			characterCorporationHistoryAllColumns,
			characterCorporationHistoryColumnsWithDefault,
			characterCorporationHistoryColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			characterCorporationHistoryAllColumns,
			characterCorporationHistoryPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert character_corporation_history, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "character_corporation_history", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `character_corporation_history` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(characterCorporationHistoryType, characterCorporationHistoryMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(characterCorporationHistoryType, characterCorporationHistoryMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for character_corporation_history")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(characterCorporationHistoryType, characterCorporationHistoryMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for character_corporation_history")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for character_corporation_history")
	}

CacheNoHooks:
	if !cached {
		characterCorporationHistoryUpsertCacheMut.Lock()
		characterCorporationHistoryUpsertCache[key] = cache
		characterCorporationHistoryUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CharacterCorporationHistory record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CharacterCorporationHistory) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CharacterCorporationHistory provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), characterCorporationHistoryPrimaryKeyMapping)
	sql := "DELETE FROM `character_corporation_history` WHERE `id`=? AND `record_id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from character_corporation_history")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for character_corporation_history")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q characterCorporationHistoryQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no characterCorporationHistoryQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from character_corporation_history")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for character_corporation_history")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CharacterCorporationHistorySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(characterCorporationHistoryBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), characterCorporationHistoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `character_corporation_history` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, characterCorporationHistoryPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from characterCorporationHistory slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for character_corporation_history")
	}

	if len(characterCorporationHistoryAfterDeleteHooks) != 0 {
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
func (o *CharacterCorporationHistory) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCharacterCorporationHistory(ctx, exec, o.ID, o.RecordID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CharacterCorporationHistorySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CharacterCorporationHistorySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), characterCorporationHistoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `character_corporation_history`.* FROM `character_corporation_history` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, characterCorporationHistoryPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CharacterCorporationHistorySlice")
	}

	*o = slice

	return nil
}

// CharacterCorporationHistoryExists checks if the CharacterCorporationHistory row exists.
func CharacterCorporationHistoryExists(ctx context.Context, exec boil.ContextExecutor, iD uint64, recordID uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `character_corporation_history` where `id`=? AND `record_id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD, recordID)
	}

	row := exec.QueryRowContext(ctx, sql, iD, recordID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if character_corporation_history exists")
	}

	return exists, nil
}