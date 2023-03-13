package ports

import (
	"net/http"

	"github.com/chiennguyen196/go-library/internal/common/server/httperr"
	"github.com/chiennguyen196/go-library/internal/lending/app/query"
	"github.com/chiennguyen196/go-library/internal/lending/domain"
)

func (h HttpServer) GetPatronProfile(w http.ResponseWriter, r *http.Request, patronId string) {
	q := query.PatronProfileQuery{PatronID: domain.PatronID(patronId)}

	profile, err := h.app.Queries.PatronProfile.Handle(r.Context(), q)
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}
	respondWithData(w, r, map[string]any{
		"patronProfile": toResponsePatronProfile(profile),
	})
}

func toResponsePatronProfile(profile query.PatronProfile) PatronProfile {
	return PatronProfile{
		CheckedOuts:      toResponseCheckedOuts(profile.CheckedOuts),
		Holds:            toResponseHolds(profile.Holds),
		OverdueCheckouts: toResponseOverdueCheckouts(profile.OverdueCheckouts),
		PatronId:         profile.PatronID,
		PatronType:       toResponsePatronType(profile.PatronType),
	}
}

func toResponsePatronType(t domain.PatronType) PatronType {
	switch t {
	case domain.PatronTypeRegular:
		return Regular
	case domain.PatronTypeResearcher:
		return Researcher
	default:
		return ""
	}
}

func toResponseHolds(holds []domain.Hold) []Hold {
	result := make([]Hold, 0, len(holds))
	for _, h := range holds {
		result = append(result, Hold{
			BookId:          string(h.BookID),
			From:            h.HoldDuration.From(),
			LibraryBranchId: string(h.PlacedAt),
			Till:            h.HoldDuration.Till(),
			IsOpenEnded:     h.HoldDuration.IsOpenEnded(),
		})
	}
	return result
}

func toResponseCheckedOuts(checkedOuts []query.CheckedOut) []CheckedOut {
	result := make([]CheckedOut, 0, len(checkedOuts))
	for _, c := range checkedOuts {
		result = append(result, CheckedOut{
			BookId:          c.BookID,
			CheckedOutAt:    c.At,
			LibraryBranchId: c.LibraryBranchID,
		})
	}
	return result
}

func toResponseOverdueCheckouts(overdueCheckouts []query.OverdueCheckout) []OverdueCheckout {
	result := make([]OverdueCheckout, 0, len(overdueCheckouts))
	for _, c := range overdueCheckouts {
		result = append(result, OverdueCheckout{
			BookId:          string(c.BookID),
			LibraryBranchId: c.LibraryBranchID,
		})
	}
	return result
}
