package ports

import (
	"github.com/chiennguyen196/go-library/internal/lending/app"
)

type HttpServer struct {
	app app.Application
}

func NewHttpServer(app app.Application) HttpServer {
	return HttpServer{app: app}
}
