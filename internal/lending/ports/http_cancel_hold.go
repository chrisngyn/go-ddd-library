package ports

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/chiennguyen196/go-library/internal/common/server/httperr"
	"github.com/chiennguyen196/go-library/internal/lending/app/command"
	"github.com/chiennguyen196/go-library/internal/lending/domain"
)

func (h HttpServer) CancelHold(w http.ResponseWriter, r *http.Request, patronId string) {
	var req CancelHoldJSONRequestBody
	if err := render.Decode(r, &req); err != nil {
		httperr.BadRequest("decode-fail", err, w, r)
		return
	}

	cmd := toCancelHoldCommand(patronId, req)

	if err := h.app.Commands.CancelHold.Handle(r.Context(), cmd); err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	respondSuccess(w, r)
}

func toCancelHoldCommand(patronID string, req CancelHoldJSONRequestBody) command.CancelHoldCommand {
	return command.CancelHoldCommand{
		PatronID: domain.PatronID(patronID),
		BookID:   domain.BookID(req.BookId),
	}
}
