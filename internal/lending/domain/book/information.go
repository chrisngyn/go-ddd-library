package book

import (
	"time"

	"github.com/google/uuid"

	commonErrors "github.com/chiennguyen196/go-library/internal/common/errors"
)

type Type struct {
	s string
}

func (t Type) IsZero() bool {
	return t == Type{}
}

var (
	TypeRestricted  = Type{"Restricted"}
	TypeCirculating = Type{"Circulating"}
)

type Status struct {
	s string
}

var (
	StatusAvailable  = Status{"Available"}
	StatusOnHold     = Status{"OnHold"}
	StatusCheckedOut = Status{"CheckedOut"}
)

type Information struct {
	BookID   uuid.UUID
	BookType Type
	PlacedAt uuid.UUID
}

func NewBookInformation(bookID uuid.UUID, bookType Type, placedAt uuid.UUID) (Information, error) {
	if bookID == uuid.Nil {
		return Information{}, commonErrors.NewIncorrectInputError("missing-book-id", "missing book id")
	}
	if bookType.IsZero() {
		return Information{}, commonErrors.NewIncorrectInputError("missing-book-type", "missing book type")
	}
	if placedAt == uuid.Nil {
		return Information{}, commonErrors.NewIncorrectInputError("missing-placed-at", "missing placed at")
	}
	return Information{
		BookID:   bookID,
		BookType: bookType,
		PlacedAt: placedAt,
	}, nil
}

func (b Information) IsRestricted() bool {
	return b.BookType == TypeRestricted
}

type HoldInformation struct {
	ByPatron uuid.UUID
	Till     time.Time
}

func (h HoldInformation) IsZero() bool {
	return h == HoldInformation{}
}

type CheckedOutInformation struct {
	ByPatron uuid.UUID
	At       time.Time
}

func (c CheckedOutInformation) IsZero() bool {
	return c == CheckedOutInformation{}
}
