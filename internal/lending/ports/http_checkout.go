package ports

import (
	"net/http"
	"time"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/go-chi/render"

	"github.com/chiennguyen196/go-library/internal/common/server/httperr"
	"github.com/chiennguyen196/go-library/internal/lending/app/command"
)

func (h HttpServer) Checkout(w http.ResponseWriter, r *http.Request, patronId openapi_types.UUID) {
	var req CheckoutJSONRequestBody
	if err := render.Decode(r, &req); err != nil {
		httperr.BadRequest("decode-fail", err, w, r)
		return
	}

	cmd := toCheckoutCommand(patronId, req)

	if err := h.app.Commands.CheckOut.Handle(r.Context(), cmd); err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	render.NoContent(w, r)
}

func toCheckoutCommand(patronID openapi_types.UUID, req CheckoutJSONRequestBody) command.CheckoutCommand {
	return command.CheckoutCommand{
		RequestAt: time.Now(),
		PatronID:  patronID,
		BookID:    req.BookId,
	}
}
