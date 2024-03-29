// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Hold is an object representing the database table.
type Hold struct {
	ID              int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	PatronID        string    `boil:"patron_id" json:"patron_id" toml:"patron_id" yaml:"patron_id"`
	BookID          string    `boil:"book_id" json:"book_id" toml:"book_id" yaml:"book_id"`
	LibraryBranchID string    `boil:"library_branch_id" json:"library_branch_id" toml:"library_branch_id" yaml:"library_branch_id"`
	HoldFrom        time.Time `boil:"hold_from" json:"hold_from" toml:"hold_from" yaml:"hold_from"`
	HoldTill        null.Time `boil:"hold_till" json:"hold_till,omitempty" toml:"hold_till" yaml:"hold_till,omitempty"`
	CreatedAt       time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt       time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	DeletedAt       null.Time `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`

	R *holdR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L holdL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var HoldColumns = struct {
	ID              string
	PatronID        string
	BookID          string
	LibraryBranchID string
	HoldFrom        string
	HoldTill        string
	CreatedAt       string
	UpdatedAt       string
	DeletedAt       string
}{
	ID:              "id",
	PatronID:        "patron_id",
	BookID:          "book_id",
	LibraryBranchID: "library_branch_id",
	HoldFrom:        "hold_from",
	HoldTill:        "hold_till",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
	DeletedAt:       "deleted_at",
}

var HoldTableColumns = struct {
	ID              string
	PatronID        string
	BookID          string
	LibraryBranchID string
	HoldFrom        string
	HoldTill        string
	CreatedAt       string
	UpdatedAt       string
	DeletedAt       string
}{
	ID:              "holds.id",
	PatronID:        "holds.patron_id",
	BookID:          "holds.book_id",
	LibraryBranchID: "holds.library_branch_id",
	HoldFrom:        "holds.hold_from",
	HoldTill:        "holds.hold_till",
	CreatedAt:       "holds.created_at",
	UpdatedAt:       "holds.updated_at",
	DeletedAt:       "holds.deleted_at",
}

// Generated where

type whereHelperint64 struct{ field string }

func (w whereHelperint64) EQ(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint64) NEQ(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint64) LT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint64) LTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint64) GT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint64) GTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint64) IN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint64) NIN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var HoldWhere = struct {
	ID              whereHelperint64
	PatronID        whereHelperstring
	BookID          whereHelperstring
	LibraryBranchID whereHelperstring
	HoldFrom        whereHelpertime_Time
	HoldTill        whereHelpernull_Time
	CreatedAt       whereHelpertime_Time
	UpdatedAt       whereHelpertime_Time
	DeletedAt       whereHelpernull_Time
}{
	ID:              whereHelperint64{field: "\"holds\".\"id\""},
	PatronID:        whereHelperstring{field: "\"holds\".\"patron_id\""},
	BookID:          whereHelperstring{field: "\"holds\".\"book_id\""},
	LibraryBranchID: whereHelperstring{field: "\"holds\".\"library_branch_id\""},
	HoldFrom:        whereHelpertime_Time{field: "\"holds\".\"hold_from\""},
	HoldTill:        whereHelpernull_Time{field: "\"holds\".\"hold_till\""},
	CreatedAt:       whereHelpertime_Time{field: "\"holds\".\"created_at\""},
	UpdatedAt:       whereHelpertime_Time{field: "\"holds\".\"updated_at\""},
	DeletedAt:       whereHelpernull_Time{field: "\"holds\".\"deleted_at\""},
}

// HoldRels is where relationship names are stored.
var HoldRels = struct {
	Patron string
}{
	Patron: "Patron",
}

// holdR is where relationships are stored.
type holdR struct {
	Patron *Patron `boil:"Patron" json:"Patron" toml:"Patron" yaml:"Patron"`
}

// NewStruct creates a new relationship struct
func (*holdR) NewStruct() *holdR {
	return &holdR{}
}

func (r *holdR) GetPatron() *Patron {
	if r == nil {
		return nil
	}
	return r.Patron
}

// holdL is where Load methods for each relationship are stored.
type holdL struct{}

var (
	holdAllColumns            = []string{"id", "patron_id", "book_id", "library_branch_id", "hold_from", "hold_till", "created_at", "updated_at", "deleted_at"}
	holdColumnsWithoutDefault = []string{"patron_id", "book_id", "library_branch_id", "hold_from"}
	holdColumnsWithDefault    = []string{"id", "hold_till", "created_at", "updated_at", "deleted_at"}
	holdPrimaryKeyColumns     = []string{"id"}
	holdGeneratedColumns      = []string{}
)

type (
	// HoldSlice is an alias for a slice of pointers to Hold.
	// This should almost always be used instead of []Hold.
	HoldSlice []*Hold
	// HoldHook is the signature for custom Hold hook methods
	HoldHook func(context.Context, boil.ContextExecutor, *Hold) error

	holdQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	holdType                 = reflect.TypeOf(&Hold{})
	holdMapping              = queries.MakeStructMapping(holdType)
	holdPrimaryKeyMapping, _ = queries.BindMapping(holdType, holdMapping, holdPrimaryKeyColumns)
	holdInsertCacheMut       sync.RWMutex
	holdInsertCache          = make(map[string]insertCache)
	holdUpdateCacheMut       sync.RWMutex
	holdUpdateCache          = make(map[string]updateCache)
	holdUpsertCacheMut       sync.RWMutex
	holdUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var holdAfterSelectHooks []HoldHook

var holdBeforeInsertHooks []HoldHook
var holdAfterInsertHooks []HoldHook

var holdBeforeUpdateHooks []HoldHook
var holdAfterUpdateHooks []HoldHook

var holdBeforeDeleteHooks []HoldHook
var holdAfterDeleteHooks []HoldHook

var holdBeforeUpsertHooks []HoldHook
var holdAfterUpsertHooks []HoldHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Hold) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range holdAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Hold) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range holdBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Hold) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range holdAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Hold) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range holdBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Hold) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range holdAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Hold) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range holdBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Hold) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range holdAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Hold) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range holdBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Hold) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range holdAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddHoldHook registers your hook function for all future operations.
func AddHoldHook(hookPoint boil.HookPoint, holdHook HoldHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		holdAfterSelectHooks = append(holdAfterSelectHooks, holdHook)
	case boil.BeforeInsertHook:
		holdBeforeInsertHooks = append(holdBeforeInsertHooks, holdHook)
	case boil.AfterInsertHook:
		holdAfterInsertHooks = append(holdAfterInsertHooks, holdHook)
	case boil.BeforeUpdateHook:
		holdBeforeUpdateHooks = append(holdBeforeUpdateHooks, holdHook)
	case boil.AfterUpdateHook:
		holdAfterUpdateHooks = append(holdAfterUpdateHooks, holdHook)
	case boil.BeforeDeleteHook:
		holdBeforeDeleteHooks = append(holdBeforeDeleteHooks, holdHook)
	case boil.AfterDeleteHook:
		holdAfterDeleteHooks = append(holdAfterDeleteHooks, holdHook)
	case boil.BeforeUpsertHook:
		holdBeforeUpsertHooks = append(holdBeforeUpsertHooks, holdHook)
	case boil.AfterUpsertHook:
		holdAfterUpsertHooks = append(holdAfterUpsertHooks, holdHook)
	}
}

// One returns a single hold record from the query.
func (q holdQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Hold, error) {
	o := &Hold{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for holds")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Hold records from the query.
func (q holdQuery) All(ctx context.Context, exec boil.ContextExecutor) (HoldSlice, error) {
	var o []*Hold

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Hold slice")
	}

	if len(holdAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Hold records in the query.
func (q holdQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count holds rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q holdQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if holds exists")
	}

	return count > 0, nil
}

// Patron pointed to by the foreign key.
func (o *Hold) Patron(mods ...qm.QueryMod) patronQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.PatronID),
	}

	queryMods = append(queryMods, mods...)

	return Patrons(queryMods...)
}

// LoadPatron allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (holdL) LoadPatron(ctx context.Context, e boil.ContextExecutor, singular bool, maybeHold interface{}, mods queries.Applicator) error {
	var slice []*Hold
	var object *Hold

	if singular {
		var ok bool
		object, ok = maybeHold.(*Hold)
		if !ok {
			object = new(Hold)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeHold)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeHold))
			}
		}
	} else {
		s, ok := maybeHold.(*[]*Hold)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeHold)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeHold))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &holdR{}
		}
		args = append(args, object.PatronID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &holdR{}
			}

			for _, a := range args {
				if a == obj.PatronID {
					continue Outer
				}
			}

			args = append(args, obj.PatronID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`patrons`),
		qm.WhereIn(`patrons.id in ?`, args...),
		qmhelper.WhereIsNull(`patrons.deleted_at`),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Patron")
	}

	var resultSlice []*Patron
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Patron")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for patrons")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for patrons")
	}

	if len(patronAfterSelectHooks) != 0 {
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
		object.R.Patron = foreign
		if foreign.R == nil {
			foreign.R = &patronR{}
		}
		foreign.R.Holds = append(foreign.R.Holds, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.PatronID == foreign.ID {
				local.R.Patron = foreign
				if foreign.R == nil {
					foreign.R = &patronR{}
				}
				foreign.R.Holds = append(foreign.R.Holds, local)
				break
			}
		}
	}

	return nil
}

// SetPatron of the hold to the related item.
// Sets o.R.Patron to related.
// Adds o to related.R.Holds.
func (o *Hold) SetPatron(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Patron) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"holds\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"patron_id"}),
		strmangle.WhereClause("\"", "\"", 2, holdPrimaryKeyColumns),
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

	o.PatronID = related.ID
	if o.R == nil {
		o.R = &holdR{
			Patron: related,
		}
	} else {
		o.R.Patron = related
	}

	if related.R == nil {
		related.R = &patronR{
			Holds: HoldSlice{o},
		}
	} else {
		related.R.Holds = append(related.R.Holds, o)
	}

	return nil
}

// Holds retrieves all the records using an executor.
func Holds(mods ...qm.QueryMod) holdQuery {
	mods = append(mods, qm.From("\"holds\""), qmhelper.WhereIsNull("\"holds\".\"deleted_at\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"holds\".*"})
	}

	return holdQuery{q}
}

// FindHold retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindHold(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Hold, error) {
	holdObj := &Hold{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"holds\" where \"id\"=$1 and \"deleted_at\" is null", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, holdObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from holds")
	}

	if err = holdObj.doAfterSelectHooks(ctx, exec); err != nil {
		return holdObj, err
	}

	return holdObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Hold) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no holds provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(holdColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	holdInsertCacheMut.RLock()
	cache, cached := holdInsertCache[key]
	holdInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			holdAllColumns,
			holdColumnsWithDefault,
			holdColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(holdType, holdMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(holdType, holdMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"holds\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"holds\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into holds")
	}

	if !cached {
		holdInsertCacheMut.Lock()
		holdInsertCache[key] = cache
		holdInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Hold.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Hold) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	holdUpdateCacheMut.RLock()
	cache, cached := holdUpdateCache[key]
	holdUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			holdAllColumns,
			holdPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update holds, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"holds\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, holdPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(holdType, holdMapping, append(wl, holdPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update holds row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for holds")
	}

	if !cached {
		holdUpdateCacheMut.Lock()
		holdUpdateCache[key] = cache
		holdUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q holdQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for holds")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for holds")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o HoldSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), holdPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"holds\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, holdPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in hold slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all hold")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Hold) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no holds provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(holdColumnsWithDefault, o)

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

	holdUpsertCacheMut.RLock()
	cache, cached := holdUpsertCache[key]
	holdUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			holdAllColumns,
			holdColumnsWithDefault,
			holdColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			holdAllColumns,
			holdPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert holds, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(holdPrimaryKeyColumns))
			copy(conflict, holdPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"holds\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(holdType, holdMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(holdType, holdMapping, ret)
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
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert holds")
	}

	if !cached {
		holdUpsertCacheMut.Lock()
		holdUpsertCache[key] = cache
		holdUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Hold record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Hold) Delete(ctx context.Context, exec boil.ContextExecutor, hardDelete bool) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Hold provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	var (
		sql  string
		args []interface{}
	)
	if hardDelete {
		args = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), holdPrimaryKeyMapping)
		sql = "DELETE FROM \"holds\" WHERE \"id\"=$1"
	} else {
		currTime := time.Now().In(boil.GetLocation())
		o.DeletedAt = null.TimeFrom(currTime)
		wl := []string{"deleted_at"}
		sql = fmt.Sprintf("UPDATE \"holds\" SET %s WHERE \"id\"=$2",
			strmangle.SetParamNames("\"", "\"", 1, wl),
		)
		valueMapping, err := queries.BindMapping(holdType, holdMapping, append(wl, holdPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
		args = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), valueMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from holds")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for holds")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q holdQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor, hardDelete bool) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no holdQuery provided for delete all")
	}

	if hardDelete {
		queries.SetDelete(q.Query)
	} else {
		currTime := time.Now().In(boil.GetLocation())
		queries.SetUpdate(q.Query, M{"deleted_at": currTime})
	}

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from holds")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for holds")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o HoldSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor, hardDelete bool) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(holdBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var (
		sql  string
		args []interface{}
	)
	if hardDelete {
		for _, obj := range o {
			pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), holdPrimaryKeyMapping)
			args = append(args, pkeyArgs...)
		}
		sql = "DELETE FROM \"holds\" WHERE " +
			strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, holdPrimaryKeyColumns, len(o))
	} else {
		currTime := time.Now().In(boil.GetLocation())
		for _, obj := range o {
			pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), holdPrimaryKeyMapping)
			args = append(args, pkeyArgs...)
			obj.DeletedAt = null.TimeFrom(currTime)
		}
		wl := []string{"deleted_at"}
		sql = fmt.Sprintf("UPDATE \"holds\" SET %s WHERE "+
			strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 2, holdPrimaryKeyColumns, len(o)),
			strmangle.SetParamNames("\"", "\"", 1, wl),
		)
		args = append([]interface{}{currTime}, args...)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from hold slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for holds")
	}

	if len(holdAfterDeleteHooks) != 0 {
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
func (o *Hold) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindHold(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *HoldSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := HoldSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), holdPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"holds\".* FROM \"holds\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, holdPrimaryKeyColumns, len(*o)) +
		"and \"deleted_at\" is null"

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in HoldSlice")
	}

	*o = slice

	return nil
}

// HoldExists checks if the Hold row exists.
func HoldExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"holds\" where \"id\"=$1 and \"deleted_at\" is null limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if holds exists")
	}

	return exists, nil
}

// Exists checks if the Hold row exists.
func (o *Hold) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return HoldExists(ctx, exec, o.ID)
}
