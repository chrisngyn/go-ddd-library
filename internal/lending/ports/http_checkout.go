package ports

import (
	"net/http"
	"time"

	"github.com/go-chi/render"

	"github.com/chiennguyen196/go-library/internal/common/server/httperr"
	"github.com/chiennguyen196/go-library/internal/lending/app/command"
	"github.com/chiennguyen196/go-library/internal/lending/domain"
)

func (h HttpServer) Checkout(w http.ResponseWriter, r *http.Request, patronId string) {
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

	respondSuccess(w, r)
}

func toCheckoutCommand(patronID string, req CheckoutJSONRequestBody) command.CheckoutCommand {
	return command.CheckoutCommand{
		RequestAt: time.Now(),
		PatronID:  domain.PatronID(patronID),
		BookID:    domain.BookID(req.BookId),
	}
}
