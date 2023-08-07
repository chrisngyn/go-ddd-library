package main

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"github.com/google/uuid"

	"github.com/chiennguyen196/go-library/internal/common/server/httperr"
)

type HttpServer struct {
	db DB
}

func (h HttpServer) CreateABook(w http.ResponseWriter, r *http.Request) {
	var req CreateABookJSONRequestBody
	if err := render.Decode(r, &req); err != nil {
		httperr.BadRequest("decode-fail", err, w, r)
		return
	}

	book, err := NewBook(req.Isbn, req.Title, req.Author)
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	if err := h.db.AddABook(r.Context(), book); err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	render.NoContent(w, r)
}

func (h HttpServer) CreateABookInstance(w http.ResponseWriter, r *http.Request, isbn string) {
	var req CreateABookInstanceJSONRequestBody
	if err := render.Decode(r, &req); err != nil {
		httperr.BadRequest("decode-fail", err, w, r)
		return
	}

	exist, err := h.db.Exists(r.Context(), isbn)
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	if !exist {
		httperr.BadRequest("isbn-not-existed", errors.New("isbn not exist"), w, r)
		return
	}

	bookInstance, event, err := NewBookInstance(uuid.NewString(), isbn, req.PlacedAt, req.BookType)
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	if err := h.db.AddABookInstance(r.Context(), bookInstance, event); err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	render.NoContent(w, r)
}
