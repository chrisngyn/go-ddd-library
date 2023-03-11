package ports

import (
	"net/http"

	"github.com/go-chi/render"
)

func respondWithData(w http.ResponseWriter, r *http.Request, data map[string]any) {
	render.Respond(w, r, dataResponse{
		Code:    "success",
		Message: "Success",
		Data:    data,
	})
}

type dataResponse struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Data    map[string]any `json:"data"`
}

func respondSuccess(w http.ResponseWriter, r *http.Request) {
	render.Respond(w, r, BaseResponse{
		Code:    "success",
		Message: "Success",
	})
}
