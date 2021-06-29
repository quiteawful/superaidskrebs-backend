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

// ChannelhasUser is an object representing the database table.
type ChannelhasUser struct {
	ChannelID int `boil:"ChannelID" json:"ChannelID" toml:"ChannelID" yaml:"ChannelID"`
	UserID    int `boil:"UserID" json:"UserID" toml:"UserID" yaml:"UserID"`
	Right     int `boil:"Right" json:"Right" toml:"Right" yaml:"Right"`

	R *channelhasUserR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L channelhasUserL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ChannelhasUserColumns = struct {
	ChannelID string
	UserID    string
	Right     string
}{
	ChannelID: "ChannelID",
	UserID:    "UserID",
	Right:     "Right",
}

var ChannelhasUserTableColumns = struct {
	ChannelID string
	UserID    string
	Right     string
}{
	ChannelID: "ChannelhasUsers.ChannelID",
	UserID:    "ChannelhasUsers.UserID",
	Right:     "ChannelhasUsers.Right",
}

// Generated where

var ChannelhasUserWhere = struct {
	ChannelID whereHelperint
	UserID    whereHelperint
	Right     whereHelperint
}{
	ChannelID: whereHelperint{field: "\"ChannelhasUsers\".\"ChannelID\""},
	UserID:    whereHelperint{field: "\"ChannelhasUsers\".\"UserID\""},
	Right:     whereHelperint{field: "\"ChannelhasUsers\".\"Right\""},
}

// ChannelhasUserRels is where relationship names are stored.
var ChannelhasUserRels = struct {
	UserIDUser       string
	ChannelIDChannel string
}{
	UserIDUser:       "UserIDUser",
	ChannelIDChannel: "ChannelIDChannel",
}

// channelhasUserR is where relationships are stored.
type channelhasUserR struct {
	UserIDUser       *User    `boil:"UserIDUser" json:"UserIDUser" toml:"UserIDUser" yaml:"UserIDUser"`
	ChannelIDChannel *Channel `boil:"ChannelIDChannel" json:"ChannelIDChannel" toml:"ChannelIDChannel" yaml:"ChannelIDChannel"`
}

// NewStruct creates a new relationship struct
func (*channelhasUserR) NewStruct() *channelhasUserR {
	return &channelhasUserR{}
}

// channelhasUserL is where Load methods for each relationship are stored.
type channelhasUserL struct{}

var (
	channelhasUserAllColumns            = []string{"ChannelID", "UserID", "Right"}
	channelhasUserColumnsWithoutDefault = []string{"ChannelID", "UserID", "Right"}
	channelhasUserColumnsWithDefault    = []string{}
	channelhasUserPrimaryKeyColumns     = []string{"ChannelID", "UserID"}
)

type (
	// ChannelhasUserSlice is an alias for a slice of pointers to ChannelhasUser.
	// This should almost always be used instead of []ChannelhasUser.
	ChannelhasUserSlice []*ChannelhasUser
	// ChannelhasUserHook is the signature for custom ChannelhasUser hook methods
	ChannelhasUserHook func(context.Context, boil.ContextExecutor, *ChannelhasUser) error

	channelhasUserQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	channelhasUserType                 = reflect.TypeOf(&ChannelhasUser{})
	channelhasUserMapping              = queries.MakeStructMapping(channelhasUserType)
	channelhasUserPrimaryKeyMapping, _ = queries.BindMapping(channelhasUserType, channelhasUserMapping, channelhasUserPrimaryKeyColumns)
	channelhasUserInsertCacheMut       sync.RWMutex
	channelhasUserInsertCache          = make(map[string]insertCache)
	channelhasUserUpdateCacheMut       sync.RWMutex
	channelhasUserUpdateCache          = make(map[string]updateCache)
	channelhasUserUpsertCacheMut       sync.RWMutex
	channelhasUserUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var channelhasUserBeforeInsertHooks []ChannelhasUserHook
var channelhasUserBeforeUpdateHooks []ChannelhasUserHook
var channelhasUserBeforeDeleteHooks []ChannelhasUserHook
var channelhasUserBeforeUpsertHooks []ChannelhasUserHook

var channelhasUserAfterInsertHooks []ChannelhasUserHook
var channelhasUserAfterSelectHooks []ChannelhasUserHook
var channelhasUserAfterUpdateHooks []ChannelhasUserHook
var channelhasUserAfterDeleteHooks []ChannelhasUserHook
var channelhasUserAfterUpsertHooks []ChannelhasUserHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ChannelhasUser) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range channelhasUserBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ChannelhasUser) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range channelhasUserBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ChannelhasUser) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range channelhasUserBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ChannelhasUser) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range channelhasUserBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ChannelhasUser) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range channelhasUserAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ChannelhasUser) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range channelhasUserAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ChannelhasUser) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range channelhasUserAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ChannelhasUser) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range channelhasUserAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ChannelhasUser) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range channelhasUserAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddChannelhasUserHook registers your hook function for all future operations.
func AddChannelhasUserHook(hookPoint boil.HookPoint, channelhasUserHook ChannelhasUserHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		channelhasUserBeforeInsertHooks = append(channelhasUserBeforeInsertHooks, channelhasUserHook)
	case boil.BeforeUpdateHook:
		channelhasUserBeforeUpdateHooks = append(channelhasUserBeforeUpdateHooks, channelhasUserHook)
	case boil.BeforeDeleteHook:
		channelhasUserBeforeDeleteHooks = append(channelhasUserBeforeDeleteHooks, channelhasUserHook)
	case boil.BeforeUpsertHook:
		channelhasUserBeforeUpsertHooks = append(channelhasUserBeforeUpsertHooks, channelhasUserHook)
	case boil.AfterInsertHook:
		channelhasUserAfterInsertHooks = append(channelhasUserAfterInsertHooks, channelhasUserHook)
	case boil.AfterSelectHook:
		channelhasUserAfterSelectHooks = append(channelhasUserAfterSelectHooks, channelhasUserHook)
	case boil.AfterUpdateHook:
		channelhasUserAfterUpdateHooks = append(channelhasUserAfterUpdateHooks, channelhasUserHook)
	case boil.AfterDeleteHook:
		channelhasUserAfterDeleteHooks = append(channelhasUserAfterDeleteHooks, channelhasUserHook)
	case boil.AfterUpsertHook:
		channelhasUserAfterUpsertHooks = append(channelhasUserAfterUpsertHooks, channelhasUserHook)
	}
}

// One returns a single channelhasUser record from the query.
func (q channelhasUserQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ChannelhasUser, error) {
	o := &ChannelhasUser{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "postgresmodel: failed to execute a one query for ChannelhasUsers")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ChannelhasUser records from the query.
func (q channelhasUserQuery) All(ctx context.Context, exec boil.ContextExecutor) (ChannelhasUserSlice, error) {
	var o []*ChannelhasUser

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "postgresmodel: failed to assign all query results to ChannelhasUser slice")
	}

	if len(channelhasUserAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ChannelhasUser records in the query.
func (q channelhasUserQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "postgresmodel: failed to count ChannelhasUsers rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q channelhasUserQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "postgresmodel: failed to check if ChannelhasUsers exists")
	}

	return count > 0, nil
}

// UserIDUser pointed to by the foreign key.
func (o *ChannelhasUser) UserIDUser(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"UserID\" = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "\"Users\"")

	return query
}

// ChannelIDChannel pointed to by the foreign key.
func (o *ChannelhasUser) ChannelIDChannel(mods ...qm.QueryMod) channelQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"ChannelID\" = ?", o.ChannelID),
	}

	queryMods = append(queryMods, mods...)

	query := Channels(queryMods...)
	queries.SetFrom(query.Query, "\"Channels\"")

	return query
}

// LoadUserIDUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (channelhasUserL) LoadUserIDUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeChannelhasUser interface{}, mods queries.Applicator) error {
	var slice []*ChannelhasUser
	var object *ChannelhasUser

	if singular {
		object = maybeChannelhasUser.(*ChannelhasUser)
	} else {
		slice = *maybeChannelhasUser.(*[]*ChannelhasUser)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &channelhasUserR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &channelhasUserR{}
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

	if len(channelhasUserAfterSelectHooks) != 0 {
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
		foreign.R.UserIDChannelhasUsers = append(foreign.R.UserIDChannelhasUsers, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.UserID {
				local.R.UserIDUser = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.UserIDChannelhasUsers = append(foreign.R.UserIDChannelhasUsers, local)
				break
			}
		}
	}

	return nil
}

// LoadChannelIDChannel allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (channelhasUserL) LoadChannelIDChannel(ctx context.Context, e boil.ContextExecutor, singular bool, maybeChannelhasUser interface{}, mods queries.Applicator) error {
	var slice []*ChannelhasUser
	var object *ChannelhasUser

	if singular {
		object = maybeChannelhasUser.(*ChannelhasUser)
	} else {
		slice = *maybeChannelhasUser.(*[]*ChannelhasUser)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &channelhasUserR{}
		}
		args = append(args, object.ChannelID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &channelhasUserR{}
			}

			for _, a := range args {
				if a == obj.ChannelID {
					continue Outer
				}
			}

			args = append(args, obj.ChannelID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`Channels`),
		qm.WhereIn(`Channels.ChannelID in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Channel")
	}

	var resultSlice []*Channel
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Channel")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for Channels")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for Channels")
	}

	if len(channelhasUserAfterSelectHooks) != 0 {
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
		object.R.ChannelIDChannel = foreign
		if foreign.R == nil {
			foreign.R = &channelR{}
		}
		foreign.R.ChannelIDChannelhasUsers = append(foreign.R.ChannelIDChannelhasUsers, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ChannelID == foreign.ChannelID {
				local.R.ChannelIDChannel = foreign
				if foreign.R == nil {
					foreign.R = &channelR{}
				}
				foreign.R.ChannelIDChannelhasUsers = append(foreign.R.ChannelIDChannelhasUsers, local)
				break
			}
		}
	}

	return nil
}

// SetUserIDUser of the channelhasUser to the related item.
// Sets o.R.UserIDUser to related.
// Adds o to related.R.UserIDChannelhasUsers.
func (o *ChannelhasUser) SetUserIDUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"ChannelhasUsers\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"UserID"}),
		strmangle.WhereClause("\"", "\"", 2, channelhasUserPrimaryKeyColumns),
	)
	values := []interface{}{related.UserID, o.ChannelID, o.UserID}

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
		o.R = &channelhasUserR{
			UserIDUser: related,
		}
	} else {
		o.R.UserIDUser = related
	}

	if related.R == nil {
		related.R = &userR{
			UserIDChannelhasUsers: ChannelhasUserSlice{o},
		}
	} else {
		related.R.UserIDChannelhasUsers = append(related.R.UserIDChannelhasUsers, o)
	}

	return nil
}

// SetChannelIDChannel of the channelhasUser to the related item.
// Sets o.R.ChannelIDChannel to related.
// Adds o to related.R.ChannelIDChannelhasUsers.
func (o *ChannelhasUser) SetChannelIDChannel(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Channel) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"ChannelhasUsers\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"ChannelID"}),
		strmangle.WhereClause("\"", "\"", 2, channelhasUserPrimaryKeyColumns),
	)
	values := []interface{}{related.ChannelID, o.ChannelID, o.UserID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ChannelID = related.ChannelID
	if o.R == nil {
		o.R = &channelhasUserR{
			ChannelIDChannel: related,
		}
	} else {
		o.R.ChannelIDChannel = related
	}

	if related.R == nil {
		related.R = &channelR{
			ChannelIDChannelhasUsers: ChannelhasUserSlice{o},
		}
	} else {
		related.R.ChannelIDChannelhasUsers = append(related.R.ChannelIDChannelhasUsers, o)
	}

	return nil
}

// ChannelhasUsers retrieves all the records using an executor.
func ChannelhasUsers(mods ...qm.QueryMod) channelhasUserQuery {
	mods = append(mods, qm.From("\"ChannelhasUsers\""))
	return channelhasUserQuery{NewQuery(mods...)}
}

// FindChannelhasUser retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindChannelhasUser(ctx context.Context, exec boil.ContextExecutor, channelID int, userID int, selectCols ...string) (*ChannelhasUser, error) {
	channelhasUserObj := &ChannelhasUser{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"ChannelhasUsers\" where \"ChannelID\"=$1 AND \"UserID\"=$2", sel,
	)

	q := queries.Raw(query, channelID, userID)

	err := q.Bind(ctx, exec, channelhasUserObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "postgresmodel: unable to select from ChannelhasUsers")
	}

	if err = channelhasUserObj.doAfterSelectHooks(ctx, exec); err != nil {
		return channelhasUserObj, err
	}

	return channelhasUserObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ChannelhasUser) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("postgresmodel: no ChannelhasUsers provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(channelhasUserColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	channelhasUserInsertCacheMut.RLock()
	cache, cached := channelhasUserInsertCache[key]
	channelhasUserInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			channelhasUserAllColumns,
			channelhasUserColumnsWithDefault,
			channelhasUserColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(channelhasUserType, channelhasUserMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(channelhasUserType, channelhasUserMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"ChannelhasUsers\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"ChannelhasUsers\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "postgresmodel: unable to insert into ChannelhasUsers")
	}

	if !cached {
		channelhasUserInsertCacheMut.Lock()
		channelhasUserInsertCache[key] = cache
		channelhasUserInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ChannelhasUser.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ChannelhasUser) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	channelhasUserUpdateCacheMut.RLock()
	cache, cached := channelhasUserUpdateCache[key]
	channelhasUserUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			channelhasUserAllColumns,
			channelhasUserPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("postgresmodel: unable to update ChannelhasUsers, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"ChannelhasUsers\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, channelhasUserPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(channelhasUserType, channelhasUserMapping, append(wl, channelhasUserPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "postgresmodel: unable to update ChannelhasUsers row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "postgresmodel: failed to get rows affected by update for ChannelhasUsers")
	}

	if !cached {
		channelhasUserUpdateCacheMut.Lock()
		channelhasUserUpdateCache[key] = cache
		channelhasUserUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q channelhasUserQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "postgresmodel: unable to update all for ChannelhasUsers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "postgresmodel: unable to retrieve rows affected for ChannelhasUsers")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ChannelhasUserSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), channelhasUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"ChannelhasUsers\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, channelhasUserPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "postgresmodel: unable to update all in channelhasUser slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "postgresmodel: unable to retrieve rows affected all in update all channelhasUser")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ChannelhasUser) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("postgresmodel: no ChannelhasUsers provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(channelhasUserColumnsWithDefault, o)

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

	channelhasUserUpsertCacheMut.RLock()
	cache, cached := channelhasUserUpsertCache[key]
	channelhasUserUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			channelhasUserAllColumns,
			channelhasUserColumnsWithDefault,
			channelhasUserColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			channelhasUserAllColumns,
			channelhasUserPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("postgresmodel: unable to upsert ChannelhasUsers, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(channelhasUserPrimaryKeyColumns))
			copy(conflict, channelhasUserPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"ChannelhasUsers\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(channelhasUserType, channelhasUserMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(channelhasUserType, channelhasUserMapping, ret)
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
		return errors.Wrap(err, "postgresmodel: unable to upsert ChannelhasUsers")
	}

	if !cached {
		channelhasUserUpsertCacheMut.Lock()
		channelhasUserUpsertCache[key] = cache
		channelhasUserUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single ChannelhasUser record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ChannelhasUser) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("postgresmodel: no ChannelhasUser provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), channelhasUserPrimaryKeyMapping)
	sql := "DELETE FROM \"ChannelhasUsers\" WHERE \"ChannelID\"=$1 AND \"UserID\"=$2"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "postgresmodel: unable to delete from ChannelhasUsers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "postgresmodel: failed to get rows affected by delete for ChannelhasUsers")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q channelhasUserQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("postgresmodel: no channelhasUserQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "postgresmodel: unable to delete all from ChannelhasUsers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "postgresmodel: failed to get rows affected by deleteall for ChannelhasUsers")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ChannelhasUserSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(channelhasUserBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), channelhasUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"ChannelhasUsers\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, channelhasUserPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "postgresmodel: unable to delete all from channelhasUser slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "postgresmodel: failed to get rows affected by deleteall for ChannelhasUsers")
	}

	if len(channelhasUserAfterDeleteHooks) != 0 {
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
func (o *ChannelhasUser) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindChannelhasUser(ctx, exec, o.ChannelID, o.UserID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ChannelhasUserSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ChannelhasUserSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), channelhasUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"ChannelhasUsers\".* FROM \"ChannelhasUsers\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, channelhasUserPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "postgresmodel: unable to reload all in ChannelhasUserSlice")
	}

	*o = slice

	return nil
}

// ChannelhasUserExists checks if the ChannelhasUser row exists.
func ChannelhasUserExists(ctx context.Context, exec boil.ContextExecutor, channelID int, userID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"ChannelhasUsers\" where \"ChannelID\"=$1 AND \"UserID\"=$2 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, channelID, userID)
	}
	row := exec.QueryRowContext(ctx, sql, channelID, userID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "postgresmodel: unable to check if ChannelhasUsers exists")
	}

	return exists, nil
}
