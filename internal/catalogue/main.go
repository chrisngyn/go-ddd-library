package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/chiennguyen196/go-library/internal/common/database"
	"github.com/chiennguyen196/go-library/internal/common/logs"
	"github.com/chiennguyen196/go-library/internal/common/server"
)

func main() {
	logs.Init()

	db := database.NewSqlDB()
	httpServer := HttpServer{db: NewDB(db)}

	server.RunHTTPServer("/api/v1", func(router chi.Router) http.Handler {
		return HandlerFromMux(httpServer, router)
	})
}
