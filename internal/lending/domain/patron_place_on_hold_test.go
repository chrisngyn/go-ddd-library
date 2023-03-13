package domain_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/chiennguyen196/go-library/internal/lending/domain"
)

func TestPatron_PlaceOnHold(t *testing.T) {
	t.Run("regular patron hold a normal book", func(t *testing.T) {
		t.Parallel()
		patron := newExamplePatron(t, domain.PatronTypeRegular, nil, nil)
		book := newExampleBookInformation(t, domain.BookTypeCirculating)
		duration := newExampleCloseEndedHoldDuration(t)

		err := patron.PlaceOnHold(book, duration)
		require.NoError(t, err)

		assert.Equal(t, 1, len(patron.Holds()))
		hold := patron.Holds()[0]
		assert.Equal(t, book.BookID, hold.BookID)
		assert.Equal(t, book.PlacedAt, hold.PlacedAt)
		assert.Equal(t, duration, hold.HoldDuration)
	})

	t.Run("researcher patron hold a restricted book", func(t *testing.T) {
		t.Parallel()
		patron := newExamplePatron(t, domain.PatronTypeResearcher, nil, nil)
		book := newExampleBookInformation(t, domain.BookTypeRestricted)
		duration := newExampleOpenEndedHoldDuration(t)

		err := patron.PlaceOnHold(book, duration)
		require.NoError(t, err)

		assert.Equal(t, 1, len(patron.Holds()))
		hold := patron.Holds()[0]
		assert.Equal(t, book.BookID, hold.BookID)
		assert.Equal(t, book.PlacedAt, hold.PlacedAt)
		assert.Equal(t, duration, hold.HoldDuration)
	})
}

func TestPatron_PlaceOnHold_invalid(t *testing.T) {
	type args struct {
		patron   domain.Patron
		book     domain.BookInformation
		duration domain.HoldDuration
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "regular patron cannot hold restricted book",
			args: args{
				patron:   newExamplePatron(t, domain.PatronTypeRegular, nil, nil),
				book:     newExampleBookInformation(t, domain.BookTypeRestricted),
				duration: newExampleCloseEndedHoldDuration(t),
			},
			wantErr: domain.ErrRegularPatronCannotHoldRestrictedBook,
		},
		{
			name: "patron rejected by overdue checkout policy",
			args: args{
				patron: newExamplePatron(t, domain.PatronTypeResearcher, nil, map[domain.LibraryBranchID][]domain.BookID{
					"library_branch_1": {domain.BookID(uuid.NewString()), domain.BookID(uuid.NewString())},
				}),
				book:     newExampleBookInformationWithLibraryBranchID(t, "library_branch_1"),
				duration: newExampleOpenEndedHoldDuration(t),
			},
			wantErr: domain.ErrMaxCountOfOverdueCheckoutsReached,
		},
		{
			name: "regular patron reach max holds",
			args: args{
				patron: newExamplePatron(t, domain.PatronTypeRegular, []domain.Hold{
					newExampleHold(t),
					newExampleHold(t),
					newExampleHold(t),
					newExampleHold(t),
					newExampleHold(t),
				}, nil),
				book:     newExampleBookInformation(t, domain.BookTypeCirculating),
				duration: newExampleCloseEndedHoldDuration(t),
			},
			wantErr: domain.ErrMaxHoldsReached,
		},
		{
			name: "only researcher can hold open-ended duration",
			args: args{
				patron:   newExamplePatron(t, domain.PatronTypeRegular, nil, nil),
				book:     newExampleBookInformation(t, domain.BookTypeCirculating),
				duration: newExampleOpenEndedHoldDuration(t),
			},
			wantErr: domain.ErrOnlyResearcherCanPlaceOpenEndedHold,
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

func newExamplePatron(t *testing.T, patronType domain.PatronType, holds []domain.Hold, overdueCheckouts map[domain.LibraryBranchID][]domain.BookID) domain.Patron {
	t.Helper()
	patron, err := domain.NewPatron(domain.PatronID(uuid.NewString()), patronType, holds, overdueCheckouts)
	if err != nil {
		t.Fatalf("Fail to create a new patron: %v", err)
	}
	return patron
}

func newExampleBookInformation(t *testing.T, bookType domain.BookType) domain.BookInformation {
	t.Helper()
	return domain.BookInformation{
		BookID:   domain.BookID(uuid.NewString()),
		BookType: bookType,
		PlacedAt: domain.LibraryBranchID(uuid.NewString()),
	}
}

func newExampleBookInformationWithLibraryBranchID(t *testing.T, placedAt domain.LibraryBranchID) domain.BookInformation {
	t.Helper()
	return domain.BookInformation{
		BookID:   domain.BookID(uuid.NewString()),
		BookType: domain.BookTypeCirculating,
		PlacedAt: placedAt,
	}
}

func newExampleCloseEndedHoldDuration(t *testing.T) domain.HoldDuration {
	t.Helper()
	d, err := domain.NewHoldDuration(time.Now(), 10)
	if err != nil {
		t.Fatalf("Fail create a close-ended hold duration: %v", err)
	}
	return d
}

func newExampleOpenEndedHoldDuration(t *testing.T) domain.HoldDuration {
	t.Helper()
	d, err := domain.NewHoldDuration(time.Now(), 0)
	if err != nil {
		t.Fatalf("Fail create a open-ended hold duration: %v", err)
	}
	return d
}

func newExampleHold(t *testing.T) domain.Hold {
	t.Helper()
	h, err := domain.NewHold(domain.BookID(uuid.NewString()), domain.LibraryBranchID(uuid.NewString()), newExampleCloseEndedHoldDuration(t))
	if err != nil {
		t.Fatalf("Fail to create new hold: %v", err)
	}
	return h
}
