package httperr

import (
	"errors"
	"net/http"
	"os"

	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"

	commonErrors "github.com/chiennguyen196/go-library/internal/common/errors"
)

func InternalError(slug string, err error, w http.ResponseWriter, r *http.Request) {
	httpRespondWithError(slug, err, w, r, "Internal server error", http.StatusInternalServerError)
}

func Unauthorised(slug string, err error, w http.ResponseWriter, r *http.Request) {
	httpRespondWithError(slug, err, w, r, "Unauthorised", http.StatusUnauthorized)
}

func BadRequest(slug string, err error, w http.ResponseWriter, r *http.Request) {
	httpRespondWithError(slug, err, w, r, "Bad request", http.StatusBadRequest)
}

func RespondWithSlugError(err error, w http.ResponseWriter, r *http.Request) {
	var slugError commonErrors.SlugError

	if ok := errors.As(err, &slugError); !ok {
		InternalError("internal-server-error", err, w, r)
		return
	}

	switch slugError.ErrorType() {
	case commonErrors.ErrorTypeIncorrectInput:
		BadRequest(slugError.Slug(), slugError, w, r)
	default:
		InternalError(slugError.Slug(), slugError, w, r)
	}
}

func httpRespondWithError(slug string, err error, w http.ResponseWriter, r *http.Request, logMSg string, status int) {
	log.Ctx(r.Context()).Warn().
		Err(err).
		Str("error-slug", slug).
		Msg(logMSg)

	message := err.Error()
	if env := os.Getenv("ENV"); env == "production" {
		message = "Internal error! Check log to see the detail"
	}

	resp := ErrorResponse{slug, message, status}

	if err := render.Render(w, r, resp); err != nil {
		panic(err)
	}
}

type ErrorResponse struct {
	Slug       string `json:"slug"`
	Message    string `json:"message"`
	httpStatus int
}

func (e ErrorResponse) Render(w http.ResponseWriter, _ *http.Request) error {
	w.WriteHeader(e.httpStatus)
	return nil
}
