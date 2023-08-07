package patron_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/chiennguyen196/go-library/internal/lending/domain/book"
	"github.com/chiennguyen196/go-library/internal/lending/domain/patron"
)

func TestPatron_PlaceOnHold(t *testing.T) {
	t.Run("regular patron hold a normal book", func(t *testing.T) {
		t.Parallel()
		aPatron := newExamplePatron(t, patron.TypeRegular, nil, nil)
		aBook := newExampleBookInformation(t, book.TypeCirculating)
		duration := newExampleCloseEndedHoldDuration(t)

		err := aPatron.PlaceOnHold(aBook, duration)
		require.NoError(t, err)

		assert.Equal(t, 1, len(aPatron.Holds()))
		aHold := aPatron.Holds()[0]
		assert.Equal(t, aBook.BookID, aHold.BookID)
		assert.Equal(t, aBook.PlacedAt, aHold.PlacedAt)
		assert.Equal(t, duration, aHold.HoldDuration)
	})

	t.Run("researcher patron hold a restricted book", func(t *testing.T) {
		t.Parallel()
		aPatron := newExamplePatron(t, patron.TypeResearcher, nil, nil)
		aBook := newExampleBookInformation(t, book.TypeRestricted)
		duration := newExampleOpenEndedHoldDuration(t)

		err := aPatron.PlaceOnHold(aBook, duration)
		require.NoError(t, err)

		assert.Equal(t, 1, len(aPatron.Holds()))
		aHold := aPatron.Holds()[0]
		assert.Equal(t, aBook.BookID, aHold.BookID)
		assert.Equal(t, aBook.PlacedAt, aHold.PlacedAt)
		assert.Equal(t, duration, aHold.HoldDuration)
	})
}

func TestPatron_PlaceOnHold_invalid(t *testing.T) {
	exampleLibraryBranchID := uuid.New()
	type args struct {
		patron   patron.Patron
		book     book.Information
		duration patron.HoldDuration
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "regular patron cannot hold restricted book",
			args: args{
				patron:   newExamplePatron(t, patron.TypeRegular, nil, nil),
				book:     newExampleBookInformation(t, book.TypeRestricted),
				duration: newExampleCloseEndedHoldDuration(t),
			},
			wantErr: patron.ErrRegularPatronCannotHoldRestrictedBook,
		},
		{
			name: "patron rejected by overdue checkout policy",
			args: args{
				patron: newExamplePatron(t, patron.TypeResearcher, nil, map[uuid.UUID][]uuid.UUID{
					exampleLibraryBranchID: {uuid.New(), uuid.New()},
				}),
				book:     newExampleBookInformationWithLibraryBranchID(t, exampleLibraryBranchID),
				duration: newExampleOpenEndedHoldDuration(t),
			},
			wantErr: patron.ErrMaxCountOfOverdueCheckoutsReached,
		},
		{
			name: "regular patron reach max holds",
			args: args{
				patron: newExamplePatron(t, patron.TypeRegular, []patron.Hold{
					newExampleHold(t),
					newExampleHold(t),
					newExampleHold(t),
					newExampleHold(t),
					newExampleHold(t),
				}, nil),
				book:     newExampleBookInformation(t, book.TypeCirculating),
				duration: newExampleCloseEndedHoldDuration(t),
			},
			wantErr: patron.ErrMaxHoldsReached,
		},
		{
			name: "only researcher can hold open-ended duration",
			args: args{
				patron:   newExamplePatron(t, patron.TypeRegular, nil, nil),
				book:     newExampleBookInformation(t, book.TypeCirculating),
				duration: newExampleOpenEndedHoldDuration(t),
			},
			wantErr: patron.ErrOnlyResearcherCanPlaceOpenEndedHold,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := tt.args.patron.PlaceOnHold(tt.args.book, tt.args.duration)
			require.Error(t, err)

			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}

func newExamplePatron(t *testing.T, patronType patron.Type, holds []patron.Hold, overdueCheckouts map[uuid.UUID][]uuid.UUID) patron.Patron {
	t.Helper()
	aPatron, err := patron.NewPatron(uuid.New(), patronType, holds, overdueCheckouts)
	if err != nil {
		t.Fatalf("Fail to create a new patron: %v", err)
	}
	return aPatron
}

func newExampleBookInformation(t *testing.T, bookType book.Type) book.Information {
	t.Helper()
	return book.Information{
		BookID:   uuid.New(),
		BookType: bookType,
		PlacedAt: uuid.New(),
	}
}

func newExampleBookInformationWithLibraryBranchID(t *testing.T, placedAt uuid.UUID) book.Information {
	t.Helper()
	return book.Information{
		BookID:   uuid.New(),
		BookType: book.TypeCirculating,
		PlacedAt: placedAt,
	}
}

func newExampleCloseEndedHoldDuration(t *testing.T) patron.HoldDuration {
	t.Helper()
	d, err := patron.NewHoldDuration(time.Now(), 10)
	if err != nil {
		t.Fatalf("Fail create a close-ended hold duration: %v", err)
	}
	return d
}

func newExampleOpenEndedHoldDuration(t *testing.T) patron.HoldDuration {
	t.Helper()
	d, err := patron.NewHoldDuration(time.Now(), 0)
	if err != nil {
		t.Fatalf("Fail create a open-ended hold duration: %v", err)
	}
	return d
}

func newExampleHold(t *testing.T) patron.Hold {
	t.Helper()
	h, err := patron.NewHold(uuid.New(), uuid.New(), newExampleCloseEndedHoldDuration(t))
	if err != nil {
		t.Fatalf("Fail to create new hold: %v", err)
	}
	return h
}
