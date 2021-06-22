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

// Tag is an object representing the database table.
type Tag struct {
	TagID int `boil:"TagID" json:"TagID" toml:"TagID" yaml:"TagID"`
	Name  int `boil:"Name" json:"Name" toml:"Name" yaml:"Name"`

	R *tagR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L tagL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TagColumns = struct {
	TagID string
	Name  string
}{
	TagID: "TagID",
	Name:  "Name",
}

var TagTableColumns = struct {
	TagID string
	Name  string
}{
	TagID: "Tags.TagID",
	Name:  "Tags.Name",
}

// Generated where

var TagWhere = struct {
	TagID whereHelperint
	Name  whereHelperint
}{
	TagID: whereHelperint{field: "\"Tags\".\"TagID\""},
	Name:  whereHelperint{field: "\"Tags\".\"Name\""},
}

// TagRels is where relationship names are stored.
var TagRels = struct {
	LinkIDLinks string
}{
	LinkIDLinks: "LinkIDLinks",
}

// tagR is where relationships are stored.
type tagR struct {
	LinkIDLinks LinkSlice `boil:"LinkIDLinks" json:"LinkIDLinks" toml:"LinkIDLinks" yaml:"LinkIDLinks"`
}

// NewStruct creates a new relationship struct
func (*tagR) NewStruct() *tagR {
	return &tagR{}
}

// tagL is where Load methods for each relationship are stored.
type tagL struct{}

var (
	tagAllColumns            = []string{"TagID", "Name"}
	tagColumnsWithoutDefault = []string{"Name"}
	tagColumnsWithDefault    = []string{"TagID"}
	tagPrimaryKeyColumns     = []string{"TagID"}
)

type (
	// TagSlice is an alias for a slice of pointers to Tag.
	// This should almost always be used instead of []Tag.
	TagSlice []*Tag
	// TagHook is the signature for custom Tag hook methods
	TagHook func(context.Context, boil.ContextExecutor, *Tag) error

	tagQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	tagType                 = reflect.TypeOf(&Tag{})
	tagMapping              = queries.MakeStructMapping(tagType)
	tagPrimaryKeyMapping, _ = queries.BindMapping(tagType, tagMapping, tagPrimaryKeyColumns)
	tagInsertCacheMut       sync.RWMutex
	tagInsertCache          = make(map[string]insertCache)
	tagUpdateCacheMut       sync.RWMutex
	tagUpdateCache          = make(map[string]updateCache)
	tagUpsertCacheMut       sync.RWMutex
	tagUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var tagBeforeInsertHooks []TagHook
var tagBeforeUpdateHooks []TagHook
var tagBeforeDeleteHooks []TagHook
var tagBeforeUpsertHooks []TagHook

var tagAfterInsertHooks []TagHook
var tagAfterSelectHooks []TagHook
var tagAfterUpdateHooks []TagHook
var tagAfterDeleteHooks []TagHook
var tagAfterUpsertHooks []TagHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Tag) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tagBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Tag) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tagBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Tag) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tagBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Tag) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tagBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Tag) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tagAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Tag) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tagAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Tag) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tagAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Tag) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tagAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Tag) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tagAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTagHook registers your hook function for all future operations.
func AddTagHook(hookPoint boil.HookPoint, tagHook TagHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		tagBeforeInsertHooks = append(tagBeforeInsertHooks, tagHook)
	case boil.BeforeUpdateHook:
		tagBeforeUpdateHooks = append(tagBeforeUpdateHooks, tagHook)
	case boil.BeforeDeleteHook:
		tagBeforeDeleteHooks = append(tagBeforeDeleteHooks, tagHook)
	case boil.BeforeUpsertHook:
		tagBeforeUpsertHooks = append(tagBeforeUpsertHooks, tagHook)
	case boil.AfterInsertHook:
		tagAfterInsertHooks = append(tagAfterInsertHooks, tagHook)
	case boil.AfterSelectHook:
		tagAfterSelectHooks = append(tagAfterSelectHooks, tagHook)
	case boil.AfterUpdateHook:
		tagAfterUpdateHooks = append(tagAfterUpdateHooks, tagHook)
	case boil.AfterDeleteHook:
		tagAfterDeleteHooks = append(tagAfterDeleteHooks, tagHook)
	case boil.AfterUpsertHook:
		tagAfterUpsertHooks = append(tagAfterUpsertHooks, tagHook)
	}
}

// One returns a single tag record from the query.
func (q tagQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Tag, error) {
	o := &Tag{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "postgresmodel: failed to execute a one query for Tags")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Tag records from the query.
func (q tagQuery) All(ctx context.Context, exec boil.ContextExecutor) (TagSlice, error) {
	var o []*Tag

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "postgresmodel: failed to assign all query results to Tag slice")
	}

	if len(tagAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Tag records in the query.
func (q tagQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "postgresmodel: failed to count Tags rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q tagQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "postgresmodel: failed to check if Tags exists")
	}

	return count > 0, nil
}

// LinkIDLinks retrieves all the Link's Links with an executor via LinkID column.
func (o *Tag) LinkIDLinks(mods ...qm.QueryMod) linkQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.InnerJoin("\"LinkHasTag\" on \"Links\".\"LinkID\" = \"LinkHasTag\".\"LinkID\""),
		qm.Where("\"LinkHasTag\".\"TagID\"=?", o.TagID),
	)

	query := Links(queryMods...)
	queries.SetFrom(query.Query, "\"Links\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"Links\".*"})
	}

	return query
}

// LoadLinkIDLinks allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (tagL) LoadLinkIDLinks(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTag interface{}, mods queries.Applicator) error {
	var slice []*Tag
	var object *Tag

	if singular {
		object = maybeTag.(*Tag)
	} else {
		slice = *maybeTag.(*[]*Tag)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &tagR{}
		}
		args = append(args, object.TagID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &tagR{}
			}

			for _, a := range args {
				if a == obj.TagID {
					continue Outer
				}
			}

			args = append(args, obj.TagID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.Select("\"Links\".LinkID, \"Links\".ChannelID, \"Links\".URL, \"Links\".Image, \"a\".\"TagID\""),
		qm.From("\"Links\""),
		qm.InnerJoin("\"LinkHasTag\" as \"a\" on \"Links\".\"LinkID\" = \"a\".\"LinkID\""),
		qm.WhereIn("\"a\".\"TagID\" in ?", args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Links")
	}

	var resultSlice []*Link

	var localJoinCols []int
	for results.Next() {
		one := new(Link)
		var localJoinCol int

		err = results.Scan(&one.LinkID, &one.ChannelID, &one.URL, &one.Image, &localJoinCol)
		if err != nil {
			return errors.Wrap(err, "failed to scan eager loaded results for Links")
		}
		if err = results.Err(); err != nil {
			return errors.Wrap(err, "failed to plebian-bind eager loaded slice Links")
		}

		resultSlice = append(resultSlice, one)
		localJoinCols = append(localJoinCols, localJoinCol)
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on Links")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for Links")
	}

	if len(linkAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.LinkIDLinks = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &linkR{}
			}
			foreign.R.TagIDTags = append(foreign.R.TagIDTags, object)
		}
		return nil
	}

	for i, foreign := range resultSlice {
		localJoinCol := localJoinCols[i]
		for _, local := range slice {
			if local.TagID == localJoinCol {
				local.R.LinkIDLinks = append(local.R.LinkIDLinks, foreign)
				if foreign.R == nil {
					foreign.R = &linkR{}
				}
				foreign.R.TagIDTags = append(foreign.R.TagIDTags, local)
				break
			}
		}
	}

	return nil
}

// AddLinkIDLinks adds the given related objects to the existing relationships
// of the Tag, optionally inserting them as new records.
// Appends related to o.R.LinkIDLinks.
// Sets related.R.TagIDTags appropriately.
func (o *Tag) AddLinkIDLinks(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Link) error {
	var err error
	for _, rel := range related {
		if insert {
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		}
	}

	for _, rel := range related {
		query := "insert into \"LinkHasTag\" (\"TagID\", \"LinkID\") values ($1, $2)"
		values := []interface{}{o.TagID, rel.LinkID}

		if boil.IsDebug(ctx) {
			writer := boil.DebugWriterFrom(ctx)
			fmt.Fprintln(writer, query)
			fmt.Fprintln(writer, values)
		}
		_, err = exec.ExecContext(ctx, query, values...)
		if err != nil {
			return errors.Wrap(err, "failed to insert into join table")
		}
	}
	if o.R == nil {
		o.R = &tagR{
			LinkIDLinks: related,
		}
	} else {
		o.R.LinkIDLinks = append(o.R.LinkIDLinks, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &linkR{
				TagIDTags: TagSlice{o},
			}
		} else {
			rel.R.TagIDTags = append(rel.R.TagIDTags, o)
		}
	}
	return nil
}

// SetLinkIDLinks removes all previously related items of the
// Tag replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.TagIDTags's LinkIDLinks accordingly.
// Replaces o.R.LinkIDLinks with related.
// Sets related.R.TagIDTags's LinkIDLinks accordingly.
func (o *Tag) SetLinkIDLinks(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Link) error {
	query := "delete from \"LinkHasTag\" where \"TagID\" = $1"
	values := []interface{}{o.TagID}
	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, query)
		fmt.Fprintln(writer, values)
	}
	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	removeLinkIDLinksFromTagIDTagsSlice(o, related)
	if o.R != nil {
		o.R.LinkIDLinks = nil
	}
	return o.AddLinkIDLinks(ctx, exec, insert, related...)
}

// RemoveLinkIDLinks relationships from objects passed in.
// Removes related items from R.LinkIDLinks (uses pointer comparison, removal does not keep order)
// Sets related.R.TagIDTags.
func (o *Tag) RemoveLinkIDLinks(ctx context.Context, exec boil.ContextExecutor, related ...*Link) error {
	if len(related) == 0 {
		return nil
	}

	var err error
	query := fmt.Sprintf(
		"delete from \"LinkHasTag\" where \"TagID\" = $1 and \"LinkID\" in (%s)",
		strmangle.Placeholders(dialect.UseIndexPlaceholders, len(related), 2, 1),
	)
	values := []interface{}{o.TagID}
	for _, rel := range related {
		values = append(values, rel.LinkID)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, query)
		fmt.Fprintln(writer, values)
	}
	_, err = exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}
	removeLinkIDLinksFromTagIDTagsSlice(o, related)
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.LinkIDLinks {
			if rel != ri {
				continue
			}

			ln := len(o.R.LinkIDLinks)
			if ln > 1 && i < ln-1 {
				o.R.LinkIDLinks[i] = o.R.LinkIDLinks[ln-1]
			}
			o.R.LinkIDLinks = o.R.LinkIDLinks[:ln-1]
			break
		}
	}

	return nil
}

func removeLinkIDLinksFromTagIDTagsSlice(o *Tag, related []*Link) {
	for _, rel := range related {
		if rel.R == nil {
			continue
		}
		for i, ri := range rel.R.TagIDTags {
			if o.TagID != ri.TagID {
				continue
			}

			ln := len(rel.R.TagIDTags)
			if ln > 1 && i < ln-1 {
				rel.R.TagIDTags[i] = rel.R.TagIDTags[ln-1]
			}
			rel.R.TagIDTags = rel.R.TagIDTags[:ln-1]
			break
		}
	}
}

// Tags retrieves all the records using an executor.
func Tags(mods ...qm.QueryMod) tagQuery {
	mods = append(mods, qm.From("\"Tags\""))
	return tagQuery{NewQuery(mods...)}
}

// FindTag retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTag(ctx context.Context, exec boil.ContextExecutor, tagID int, selectCols ...string) (*Tag, error) {
	tagObj := &Tag{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"Tags\" where \"TagID\"=$1", sel,
	)

	q := queries.Raw(query, tagID)

	err := q.Bind(ctx, exec, tagObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "postgresmodel: unable to select from Tags")
	}

	if err = tagObj.doAfterSelectHooks(ctx, exec); err != nil {
		return tagObj, err
	}

	return tagObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Tag) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("postgresmodel: no Tags provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(tagColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	tagInsertCacheMut.RLock()
	cache, cached := tagInsertCache[key]
	tagInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			tagAllColumns,
			tagColumnsWithDefault,
			tagColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(tagType, tagMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(tagType, tagMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"Tags\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"Tags\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "postgresmodel: unable to insert into Tags")
	}

	if !cached {
		tagInsertCacheMut.Lock()
		tagInsertCache[key] = cache
		tagInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Tag.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Tag) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	tagUpdateCacheMut.RLock()
	cache, cached := tagUpdateCache[key]
	tagUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			tagAllColumns,
			tagPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("postgresmodel: unable to update Tags, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"Tags\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, tagPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(tagType, tagMapping, append(wl, tagPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "postgresmodel: unable to update Tags row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "postgresmodel: failed to get rows affected by update for Tags")
	}

	if !cached {
		tagUpdateCacheMut.Lock()
		tagUpdateCache[key] = cache
		tagUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q tagQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "postgresmodel: unable to update all for Tags")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "postgresmodel: unable to retrieve rows affected for Tags")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TagSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tagPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"Tags\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, tagPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "postgresmodel: unable to update all in tag slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "postgresmodel: unable to retrieve rows affected all in update all tag")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Tag) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("postgresmodel: no Tags provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(tagColumnsWithDefault, o)

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

	tagUpsertCacheMut.RLock()
	cache, cached := tagUpsertCache[key]
	tagUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			tagAllColumns,
			tagColumnsWithDefault,
			tagColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			tagAllColumns,
			tagPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("postgresmodel: unable to upsert Tags, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(tagPrimaryKeyColumns))
			copy(conflict, tagPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"Tags\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(tagType, tagMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(tagType, tagMapping, ret)
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
		return errors.Wrap(err, "postgresmodel: unable to upsert Tags")
	}

	if !cached {
		tagUpsertCacheMut.Lock()
		tagUpsertCache[key] = cache
		tagUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Tag record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Tag) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("postgresmodel: no Tag provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), tagPrimaryKeyMapping)
	sql := "DELETE FROM \"Tags\" WHERE \"TagID\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "postgresmodel: unable to delete from Tags")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "postgresmodel: failed to get rows affected by delete for Tags")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q tagQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("postgresmodel: no tagQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "postgresmodel: unable to delete all from Tags")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "postgresmodel: failed to get rows affected by deleteall for Tags")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TagSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(tagBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tagPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"Tags\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, tagPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "postgresmodel: unable to delete all from tag slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "postgresmodel: failed to get rows affected by deleteall for Tags")
	}

	if len(tagAfterDeleteHooks) != 0 {
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
func (o *Tag) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTag(ctx, exec, o.TagID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TagSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TagSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tagPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"Tags\".* FROM \"Tags\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, tagPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "postgresmodel: unable to reload all in TagSlice")
	}

	*o = slice

	return nil
}

// TagExists checks if the Tag row exists.
func TagExists(ctx context.Context, exec boil.ContextExecutor, tagID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"Tags\" where \"TagID\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, tagID)
	}
	row := exec.QueryRowContext(ctx, sql, tagID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "postgresmodel: unable to check if Tags exists")
	}

	return exists, nil
}
