package ports

import (
	"net/http"

	"github.com/chiennguyen196/go-library/internal/common/server/httperr"
	"github.com/chiennguyen196/go-library/internal/lending/app/command"
	"github.com/chiennguyen196/go-library/internal/lending/domain"
)

func (h HttpServer) ReturnBook(w http.ResponseWriter, r *http.Request, bookId string) {
	if err := h.app.Commands.ReturnBook.Handle(r.Context(), command.ReturnBookCommand{
		BookID: domain.BookID(bookId),
	}); err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}
	respondSuccess(w, r)
}
