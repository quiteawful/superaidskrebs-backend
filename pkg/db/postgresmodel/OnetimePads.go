// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package postgresmodel

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// OnetimePad is an object representing the database table.
type OnetimePad struct {
	Time   time.Time `boil:"Time" json:"Time" toml:"Time" yaml:"Time"`
	Key    string    `boil:"Key" json:"Key" toml:"Key" yaml:"Key"`
	UserID int       `boil:"UserID" json:"UserID" toml:"UserID" yaml:"UserID"`
	Type   string    `boil:"Type" json:"Type" toml:"Type" yaml:"Type"`

	R *onetimePadR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L onetimePadL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var OnetimePadColumns = struct {
	Time   string
	Key    string
	UserID string
	Type   string
}{
	Time:   "Time",
	Key:    "Key",
	UserID: "UserID",
	Type:   "Type",
}

var OnetimePadTableColumns = struct {
	Time   string
	Key    string
	UserID string
	Type   string
}{
	Time:   "OnetimePads.Time",
	Key:    "OnetimePads.Key",
	UserID: "OnetimePads.UserID",
	Type:   "OnetimePads.Type",
}

// Generated where

var OnetimePadWhere = struct {
	Time   whereHelpertime_Time
	Key    whereHelperstring
	UserID whereHelperint
	Type   whereHelperstring
}{
	Time:   whereHelpertime_Time{field: "\"OnetimePads\".\"Time\""},
	Key:    whereHelperstring{field: "\"OnetimePads\".\"Key\""},
	UserID: whereHelperint{field: "\"OnetimePads\".\"UserID\""},
	Type:   whereHelperstring{field: "\"OnetimePads\".\"Type\""},
}

// OnetimePadRels is where relationship names are stored.
var OnetimePadRels = struct {
	UserIDUser string
}{
	UserIDUser: "UserIDUser",
}

// onetimePadR is where relationships are stored.
type onetimePadR struct {
	UserIDUser *User `boil:"UserIDUser" json:"UserIDUser" toml:"UserIDUser" yaml:"UserIDUser"`
}

// NewStruct creates a new relationship struct
func (*onetimePadR) NewStruct() *onetimePadR {
	return &onetimePadR{}
}

// onetimePadL is where Load methods for each relationship are stored.
type onetimePadL struct{}

var (
	onetimePadAllColumns            = []string{"Time", "Key", "UserID", "Type"}
	onetimePadColumnsWithoutDefault = []string{"Time", "Key", "UserID", "Type"}
	onetimePadColumnsWithDefault    = []string{}
	onetimePadPrimaryKeyColumns     = []string{"Time"}
)

type (
	// OnetimePadSlice is an alias for a slice of pointers to OnetimePad.
	// This should almost always be used instead of []OnetimePad.
	OnetimePadSlice []*OnetimePad
	// OnetimePadHook is the signature for custom OnetimePad hook methods
	OnetimePadHook func(context.Context, boil.ContextExecutor, *OnetimePad) error

	onetimePadQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	onetimePadType                 = reflect.TypeOf(&OnetimePad{})
	onetimePadMapping              = queries.MakeStructMapping(onetimePadType)
	onetimePadPrimaryKeyMapping, _ = queries.BindMapping(onetimePadType, onetimePadMapping, onetimePadPrimaryKeyColumns)
	onetimePadInsertCacheMut       sync.RWMutex
	onetimePadInsertCache          = make(map[string]insertCache)
	onetimePadUpdateCacheMut       sync.RWMutex
	onetimePadUpdateCache          = make(map[string]updateCache)
	onetimePadUpsertCacheMut       sync.RWMutex
	onetimePadUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var onetimePadBeforeInsertHooks []OnetimePadHook
var onetimePadBeforeUpdateHooks []OnetimePadHook
var onetimePadBeforeDeleteHooks []OnetimePadHook
var onetimePadBeforeUpsertHooks []OnetimePadHook

var onetimePadAfterInsertHooks []OnetimePadHook
var onetimePadAfterSelectHooks []OnetimePadHook
var onetimePadAfterUpdateHooks []OnetimePadHook
var onetimePadAfterDeleteHooks []OnetimePadHook
var onetimePadAfterUpsertHooks []OnetimePadHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *OnetimePad) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range onetimePadBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *OnetimePad) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range onetimePadBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *OnetimePad) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range onetimePadBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *OnetimePad) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range onetimePadBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *OnetimePad) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range onetimePadAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *OnetimePad) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range onetimePadAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *OnetimePad) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range onetimePadAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *OnetimePad) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range onetimePadAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *OnetimePad) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range onetimePadAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddOnetimePadHook registers your hook function for all future operations.
func AddOnetimePadHook(hookPoint boil.HookPoint, onetimePadHook OnetimePadHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		onetimePadBeforeInsertHooks = append(onetimePadBeforeInsertHooks, onetimePadHook)
	case boil.BeforeUpdateHook:
		onetimePadBeforeUpdateHooks = append(onetimePadBeforeUpdateHooks, onetimePadHook)
	case boil.BeforeDeleteHook:
		onetimePadBeforeDeleteHooks = append(onetimePadBeforeDeleteHooks, onetimePadHook)
	case boil.BeforeUpsertHook:
		onetimePadBeforeUpsertHooks = append(onetimePadBeforeUpsertHooks, onetimePadHook)
	case boil.AfterInsertHook:
		onetimePadAfterInsertHooks = append(onetimePadAfterInsertHooks, onetimePadHook)
	case boil.AfterSelectHook:
		onetimePadAfterSelectHooks = append(onetimePadAfterSelectHooks, onetimePadHook)
	case boil.AfterUpdateHook:
		onetimePadAfterUpdateHooks = append(onetimePadAfterUpdateHooks, onetimePadHook)
	case boil.AfterDeleteHook:
		onetimePadAfterDeleteHooks = append(onetimePadAfterDeleteHooks, onetimePadHook)
	case boil.AfterUpsertHook:
		onetimePadAfterUpsertHooks = append(onetimePadAfterUpsertHooks, onetimePadHook)
	}
}

// One returns a single onetimePad record from the query.
func (q onetimePadQuery) One(ctx context.Context, exec boil.ContextExecutor) (*OnetimePad, error) {
	o := &OnetimePad{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "postgresmodel: failed to execute a one query for OnetimePads")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all OnetimePad records from the query.
func (q onetimePadQuery) All(ctx context.Context, exec boil.ContextExecutor) (OnetimePadSlice, error) {
	var o []*OnetimePad

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "postgresmodel: failed to assign all query results to OnetimePad slice")
	}

	if len(onetimePadAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all OnetimePad records in the query.
func (q onetimePadQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "postgresmodel: failed to count OnetimePads rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q onetimePadQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "postgresmodel: failed to check if OnetimePads exists")
	}

	return count > 0, nil
}

// UserIDUser pointed to by the foreign key.
func (o *OnetimePad) UserIDUser(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"UserID\" = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "\"Users\"")

	return query
}

// LoadUserIDUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (onetimePadL) LoadUserIDUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeOnetimePad interface{}, mods queries.Applicator) error {
	var slice []*OnetimePad
	var object *OnetimePad

	if singular {
		object = maybeOnetimePad.(*OnetimePad)
	} else {
		slice = *maybeOnetimePad.(*[]*OnetimePad)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &onetimePadR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &onetimePadR{}
			}

			for _, a := range args {
				if a == obj.UserID {
					continue Outer
				}
			}

			args = append(args, obj.UserID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`Users`),
		qm.WhereIn(`Users.UserID in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for Users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for Users")
	}

	if len(onetimePadAfterSelectHooks) != 0 {
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
		object.R.UserIDUser = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.UserIDOnetimePads = append(foreign.R.UserIDOnetimePads, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.UserID {
				local.R.UserIDUser = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.UserIDOnetimePads = append(foreign.R.UserIDOnetimePads, local)
				break
			}
		}
	}

	return nil
}

// SetUserIDUser of the onetimePad to the related item.
// Sets o.R.UserIDUser to related.
// Adds o to related.R.UserIDOnetimePads.
func (o *OnetimePad) SetUserIDUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"OnetimePads\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"UserID"}),
		strmangle.WhereClause("\"", "\"", 2, onetimePadPrimaryKeyColumns),
	)
	values := []interface{}{related.UserID, o.Time}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.UserID = related.UserID
	if o.R == nil {
		o.R = &onetimePadR{
			UserIDUser: related,
		}
	} else {
		o.R.UserIDUser = related
	}

	if related.R == nil {
		related.R = &userR{
			UserIDOnetimePads: OnetimePadSlice{o},
		}
	} else {
		related.R.UserIDOnetimePads = append(related.R.UserIDOnetimePads, o)
	}

	return nil
}

// OnetimePads retrieves all the records using an executor.
func OnetimePads(mods ...qm.QueryMod) onetimePadQuery {
	mods = append(mods, qm.From("\"OnetimePads\""))
	return onetimePadQuery{NewQuery(mods...)}
}

// FindOnetimePad retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindOnetimePad(ctx context.Context, exec boil.ContextExecutor, time time.Time, selectCols ...string) (*OnetimePad, error) {
	onetimePadObj := &OnetimePad{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"OnetimePads\" where \"Time\"=$1", sel,
	)

	q := queries.Raw(query, time)

	err := q.Bind(ctx, exec, onetimePadObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "postgresmodel: unable to select from OnetimePads")
	}

	if err = onetimePadObj.doAfterSelectHooks(ctx, exec); err != nil {
		return onetimePadObj, err
	}

	return onetimePadObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *OnetimePad) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("postgresmodel: no OnetimePads provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(onetimePadColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	onetimePadInsertCacheMut.RLock()
	cache, cached := onetimePadInsertCache[key]
	onetimePadInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			onetimePadAllColumns,
			onetimePadColumnsWithDefault,
			onetimePadColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(onetimePadType, onetimePadMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(onetimePadType, onetimePadMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"OnetimePads\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"OnetimePads\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
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

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "postgresmodel: unable to insert into OnetimePads")
	}

	if !cached {
		onetimePadInsertCacheMut.Lock()
		onetimePadInsertCache[key] = cache
		onetimePadInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the OnetimePad.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *OnetimePad) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	onetimePadUpdateCacheMut.RLock()
	cache, cached := onetimePadUpdateCache[key]
	onetimePadUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			onetimePadAllColumns,
			onetimePadPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("postgresmodel: unable to update OnetimePads, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"OnetimePads\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, onetimePadPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(onetimePadType, onetimePadMapping, append(wl, onetimePadPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "postgresmodel: unable to update OnetimePads row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "postgresmodel: failed to get rows affected by update for OnetimePads")
	}

	if !cached {
		onetimePadUpdateCacheMut.Lock()
		onetimePadUpdateCache[key] = cache
		onetimePadUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q onetimePadQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "postgresmodel: unable to update all for OnetimePads")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "postgresmodel: unable to retrieve rows affected for OnetimePads")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o OnetimePadSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("postgresmodel: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), onetimePadPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"OnetimePads\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, onetimePadPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "postgresmodel: unable to update all in onetimePad slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "postgresmodel: unable to retrieve rows affected all in update all onetimePad")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *OnetimePad) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("postgresmodel: no OnetimePads provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(onetimePadColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
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
	key := buf.String()
	strmangle.PutBuffer(buf)

	onetimePadUpsertCacheMut.RLock()
	cache, cached := onetimePadUpsertCache[key]
	onetimePadUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			onetimePadAllColumns,
			onetimePadColumnsWithDefault,
			onetimePadColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			onetimePadAllColumns,
			onetimePadPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("postgresmodel: unable to upsert OnetimePads, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(onetimePadPrimaryKeyColumns))
			copy(conflict, onetimePadPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"OnetimePads\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(onetimePadType, onetimePadMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(onetimePadType, onetimePadMapping, ret)
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

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "postgresmodel: unable to upsert OnetimePads")
	}

	if !cached {
		onetimePadUpsertCacheMut.Lock()
		onetimePadUpsertCache[key] = cache
		onetimePadUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single OnetimePad record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *OnetimePad) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("postgresmodel: no OnetimePad provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), onetimePadPrimaryKeyMapping)
	sql := "DELETE FROM \"OnetimePads\" WHERE \"Time\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "postgresmodel: unable to delete from OnetimePads")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "postgresmodel: failed to get rows affected by delete for OnetimePads")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q onetimePadQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("postgresmodel: no onetimePadQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "postgresmodel: unable to delete all from OnetimePads")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "postgresmodel: failed to get rows affected by deleteall for OnetimePads")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o OnetimePadSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(onetimePadBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), onetimePadPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"OnetimePads\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, onetimePadPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "postgresmodel: unable to delete all from onetimePad slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "postgresmodel: failed to get rows affected by deleteall for OnetimePads")
	}

	if len(onetimePadAfterDeleteHooks) != 0 {
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
func (o *OnetimePad) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindOnetimePad(ctx, exec, o.Time)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OnetimePadSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := OnetimePadSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), onetimePadPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"OnetimePads\".* FROM \"OnetimePads\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, onetimePadPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "postgresmodel: unable to reload all in OnetimePadSlice")
	}

	*o = slice

	return nil
}

// OnetimePadExists checks if the OnetimePad row exists.
func OnetimePadExists(ctx context.Context, exec boil.ContextExecutor, time time.Time) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"OnetimePads\" where \"Time\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, time)
	}
	row := exec.QueryRowContext(ctx, sql, time)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "postgresmodel: unable to check if OnetimePads exists")
	}

	return exists, nil
}
