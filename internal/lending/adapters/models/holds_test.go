// Code generated by SQLBoiler 4.14.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testHolds(t *testing.T) {
	t.Parallel()

	query := Holds()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testHoldsSoftDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Hold{}
	if err = randomize.Struct(seed, o, holdDBTypes, true, holdColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hold struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx, false); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Holds().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testHoldsQuerySoftDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Hold{}
	if err = randomize.Struct(seed, o, holdDBTypes, true, holdColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hold struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Holds().DeleteAll(ctx, tx, false); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Holds().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testHoldsSliceSoftDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Hold{}
	if err = randomize.Struct(seed, o, holdDBTypes, true, holdColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hold struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := HoldSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx, false); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Holds().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testHoldsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Hold{}
	if err = randomize.Struct(seed, o, holdDBTypes, true, holdColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hold struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx, true); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Holds().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testHoldsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Hold{}
	if err = randomize.Struct(seed, o, holdDBTypes, true, holdColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hold struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Holds().DeleteAll(ctx, tx, true); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Holds().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testHoldsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Hold{}
	if err = randomize.Struct(seed, o, holdDBTypes, true, holdColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hold struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := HoldSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx, true); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Holds().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testHoldsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Hold{}
	if err = randomize.Struct(seed, o, holdDBTypes, true, holdColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hold struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := HoldExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Hold exists: %s", err)
	}
	if !e {
		t.Errorf("Expected HoldExists to return true, but got false.")
	}
}

func testHoldsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Hold{}
	if err = randomize.Struct(seed, o, holdDBTypes, true, holdColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hold struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	holdFound, err := FindHold(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if holdFound == nil {
		t.Error("want a record, got nil")
	}
}

func testHoldsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Hold{}
	if err = randomize.Struct(seed, o, holdDBTypes, true, holdColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hold struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Holds().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testHoldsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Hold{}
	if err = randomize.Struct(seed, o, holdDBTypes, true, holdColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hold struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Holds().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testHoldsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	holdOne := &Hold{}
	holdTwo := &Hold{}
	if err = randomize.Struct(seed, holdOne, holdDBTypes, false, holdColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hold struct: %s", err)
	}
	if err = randomize.Struct(seed, holdTwo, holdDBTypes, false, holdColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hold struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = holdOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = holdTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Holds().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testHoldsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	holdOne := &Hold{}
	holdTwo := &Hold{}
	if err = randomize.Struct(seed, holdOne, holdDBTypes, false, holdColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hold struct: %s", err)
	}
	if err = randomize.Struct(seed, holdTwo, holdDBTypes, false, holdColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hold struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = holdOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = holdTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Holds().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func holdBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Hold) error {
	*o = Hold{}
	return nil
}

func holdAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Hold) error {
	*o = Hold{}
	return nil
}

func holdAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Hold) error {
	*o = Hold{}
	return nil
}

func holdBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Hold) error {
	*o = Hold{}
	return nil
}

func holdAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Hold) error {
	*o = Hold{}
	return nil
}

func holdBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Hold) error {
	*o = Hold{}
	return nil
}

func holdAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Hold) error {
	*o = Hold{}
	return nil
}

func holdBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Hold) error {
	*o = Hold{}
	return nil
}

func holdAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Hold) error {
	*o = Hold{}
	return nil
}

func testHoldsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Hold{}
	o := &Hold{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, holdDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Hold object: %s", err)
	}

	AddHoldHook(boil.BeforeInsertHook, holdBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	holdBeforeInsertHooks = []HoldHook{}

	AddHoldHook(boil.AfterInsertHook, holdAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	holdAfterInsertHooks = []HoldHook{}

	AddHoldHook(boil.AfterSelectHook, holdAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	holdAfterSelectHooks = []HoldHook{}

	AddHoldHook(boil.BeforeUpdateHook, holdBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	holdBeforeUpdateHooks = []HoldHook{}

	AddHoldHook(boil.AfterUpdateHook, holdAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	holdAfterUpdateHooks = []HoldHook{}

	AddHoldHook(boil.BeforeDeleteHook, holdBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	holdBeforeDeleteHooks = []HoldHook{}

	AddHoldHook(boil.AfterDeleteHook, holdAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	holdAfterDeleteHooks = []HoldHook{}

	AddHoldHook(boil.BeforeUpsertHook, holdBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	holdBeforeUpsertHooks = []HoldHook{}

	AddHoldHook(boil.AfterUpsertHook, holdAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	holdAfterUpsertHooks = []HoldHook{}
}

func testHoldsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Hold{}
	if err = randomize.Struct(seed, o, holdDBTypes, true, holdColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hold struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Holds().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testHoldsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Hold{}
	if err = randomize.Struct(seed, o, holdDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Hold struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(holdColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Holds().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testHoldToOnePatronUsingPatron(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Hold
	var foreign Patron

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, holdDBTypes, false, holdColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hold struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, patronDBTypes, false, patronColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patron struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.PatronID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Patron().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	ranAfterSelectHook := false
	AddPatronHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *Patron) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := HoldSlice{&local}
	if err = local.L.LoadPatron(ctx, tx, false, (*[]*Hold)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Patron == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Patron = nil
	if err = local.L.LoadPatron(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Patron == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testHoldToOneSetOpPatronUsingPatron(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Hold
	var b, c Patron

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, holdDBTypes, false, strmangle.SetComplement(holdPrimaryKeyColumns, holdColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, patronDBTypes, false, strmangle.SetComplement(patronPrimaryKeyColumns, patronColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, patronDBTypes, false, strmangle.SetComplement(patronPrimaryKeyColumns, patronColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Patron{&b, &c} {
		err = a.SetPatron(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Patron != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Holds[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.PatronID != x.ID {
			t.Error("foreign key was wrong value", a.PatronID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.PatronID))
		reflect.Indirect(reflect.ValueOf(&a.PatronID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.PatronID != x.ID {
			t.Error("foreign key was wrong value", a.PatronID, x.ID)
		}
	}
}

func testHoldsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Hold{}
	if err = randomize.Struct(seed, o, holdDBTypes, true, holdColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hold struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testHoldsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Hold{}
	if err = randomize.Struct(seed, o, holdDBTypes, true, holdColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hold struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := HoldSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testHoldsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Hold{}
	if err = randomize.Struct(seed, o, holdDBTypes, true, holdColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hold struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Holds().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	holdDBTypes = map[string]string{`ID`: `bigint`, `PatronID`: `character varying`, `BookID`: `uuid`, `LibraryBranchID`: `uuid`, `HoldFrom`: `timestamp without time zone`, `HoldTill`: `timestamp without time zone`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`, `DeletedAt`: `timestamp without time zone`}
	_           = bytes.MinRead
)

func testHoldsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(holdPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(holdAllColumns) == len(holdPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Hold{}
	if err = randomize.Struct(seed, o, holdDBTypes, true, holdColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hold struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Holds().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, holdDBTypes, true, holdPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Hold struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testHoldsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(holdAllColumns) == len(holdPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Hold{}
	if err = randomize.Struct(seed, o, holdDBTypes, true, holdColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hold struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Holds().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, holdDBTypes, true, holdPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Hold struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(holdAllColumns, holdPrimaryKeyColumns) {
		fields = holdAllColumns
	} else {
		fields = strmangle.SetComplement(
			holdAllColumns,
			holdPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := HoldSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testHoldsUpsert(t *testing.T) {
	t.Parallel()

	if len(holdAllColumns) == len(holdPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Hold{}
	if err = randomize.Struct(seed, &o, holdDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Hold struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Hold: %s", err)
	}

	count, err := Holds().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, holdDBTypes, false, holdPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Hold struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Hold: %s", err)
	}

	count, err = Holds().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}