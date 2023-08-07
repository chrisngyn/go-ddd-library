package ports

import (
	"net/http"
	"time"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/go-chi/render"
	"github.com/pkg/errors"

	"github.com/chiennguyen196/go-library/internal/common/server/httperr"
	"github.com/chiennguyen196/go-library/internal/lending/app/command"
	"github.com/chiennguyen196/go-library/internal/lending/domain/patron"
)

func (h HttpServer) PlaceHold(w http.ResponseWriter, r *http.Request, patronId openapi_types.UUID) {
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

func buildPlaceOnHoldCommand(patronID openapi_types.UUID, req PlaceHoldJSONRequestBody) (cmd command.PlaceOnHoldCommand, err error) {
	holdDuration, err := patron.NewHoldDuration(time.Now(), req.NumOfDays)
	if err != nil {
		return cmd, errors.Wrap(err, "create hold duration")
	}
	return command.PlaceOnHoldCommand{
		PatronID:     patronID,
		BookID:       req.BookId,
		HoldDuration: holdDuration,
	}, nil
}
