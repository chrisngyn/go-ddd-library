package ports

import (
	"net/http"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/go-chi/render"

	"github.com/chiennguyen196/go-library/internal/common/server/httperr"
	"github.com/chiennguyen196/go-library/internal/lending/app/command"
)

func (h HttpServer) ReturnBook(w http.ResponseWriter, r *http.Request, bookId openapi_types.UUID) {
	if err := h.app.Commands.ReturnBook.Handle(r.Context(), command.ReturnBookCommand{
		BookID: bookId,
	}); err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}
	render.NoContent(w, r)
}
