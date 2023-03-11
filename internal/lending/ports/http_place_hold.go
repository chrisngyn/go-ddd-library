package ports

import (
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/pkg/errors"

	"github.com/chiennguyen196/go-library/internal/common/server/httperr"
	"github.com/chiennguyen196/go-library/internal/lending/app/command"
	"github.com/chiennguyen196/go-library/internal/lending/domain"
)

func (h HttpServer) PlaceHold(w http.ResponseWriter, r *http.Request, patronId string) {
	var req PlaceHoldJSONRequestBody
	if err := render.Decode(r, &req); err != nil {
		httperr.BadRequest("decode-fail", err, w, r)
		return
	}

	cmd, err := buildPlaceOnHoldCommand(patronId, req)
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	if err := h.app.Commands.PlaceOnHold.Handle(r.Context(), cmd); err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}
	respondSuccess(w, r)
}

func buildPlaceOnHoldCommand(patronID string, req PlaceHoldJSONRequestBody) (cmd command.PlaceOnHoldCommand, err error) {
	holdDuration, err := domain.NewHoldDuration(time.Now(), req.NumOfDays)
	if err != nil {
		return cmd, errors.Wrap(err, "create hold duration")
	}
	return command.PlaceOnHoldCommand{
		PatronID:     domain.PatronID(patronID),
		BookID:       domain.BookID(req.BookId),
		HoldDuration: holdDuration,
	}, nil
}
