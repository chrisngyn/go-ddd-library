package server

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	"github.com/rs/zerolog/log"

	"github.com/chiennguyen196/go-library/internal/common/logs"
)

func RunHTTPServer(apiParentRoute string, createHandler func(router chi.Router) http.Handler) {
	RunHTTPServerOnAddr(":"+os.Getenv("HTTP_PORT"), apiParentRoute, createHandler)
}

func RunHTTPServerOnAddr(addr string, apiParentRoute string, createHandler func(router chi.Router) http.Handler) {
	apiRouter := chi.NewRouter()
	setMiddlewares(apiRouter)

	rootRouter := chi.NewRouter()
	// we are mounting all APIs under api parent route
	rootRouter.Mount(apiParentRoute, createHandler(apiRouter))
	// those api for check health on K8s
	setHealthApi(rootRouter)

	log.Info().Msgf("Starting HTTP server at address=%s", addr)

	if err := http.ListenAndServe(addr, rootRouter); err != nil {
		log.Fatal().Err(err).Msg("Stopped")
	}
}

func setHealthApi(router *chi.Mux) {
	router.Get("/health", writeOK)
	router.Get("/info", writeOK)
}

func writeOK(writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(http.StatusOK)
	_, _ = writer.Write([]byte("OK"))
}

func setMiddlewares(router *chi.Mux) {
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(logs.NewHTTPStructuredLogger(log.Logger))
	router.Use(middleware.Recoverer)

	router.Use(cors.AllowAll().Handler)

	router.Use(
		middleware.SetHeader("X-Content-Type-Options", "nosniff"),
		middleware.SetHeader("X-Frame-Options", "deny"),
	)
	router.Use(middleware.NoCache)
}
