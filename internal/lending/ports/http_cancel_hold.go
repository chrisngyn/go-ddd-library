package ports

import (
	"net/http"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/go-chi/render"

	"github.com/chiennguyen196/go-library/internal/common/server/httperr"
	"github.com/chiennguyen196/go-library/internal/lending/app/command"
)

func (h HttpServer) CancelHold(w http.ResponseWriter, r *http.Request, patronId openapi_types.UUID) {
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

func toCancelHoldCommand(patronID openapi_types.UUID, req CancelHoldJSONRequestBody) command.CancelHoldCommand {
	return command.CancelHoldCommand{
		PatronID: patronID,
		BookID:   req.BookId,
	}
}
