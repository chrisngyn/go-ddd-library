package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	"github.com/rs/zerolog/log"

	"github.com/chiennguyen196/go-library/internal/common/logs"
	"github.com/chiennguyen196/go-library/internal/lending/app"
	"github.com/chiennguyen196/go-library/internal/lending/ports"
	"github.com/chiennguyen196/go-library/internal/lending/service"
)

func main() {
	logs.Init()

	severToRun := strings.ToLower(os.Args[1])

	anApp := service.NewApplication()

	switch severToRun {
	case "http":
		RunHTTPServer(func(router chi.Router) http.Handler {
			return ports.HandlerFromMux(ports.NewHttpServer(anApp), router)
		})
	case "job":
		jobName := strings.ToLower(os.Args[2])
		RunJob(jobName, anApp)
	default:
		panic(fmt.Sprintf("Not support serverToRun=%s", severToRun))
	}
}

func RunHTTPServer(createHandler func(router chi.Router) http.Handler) {
	RunHTTPServerOnAddr(":"+os.Getenv("HTTP_PORT"), createHandler)
}

func RunHTTPServerOnAddr(addr string, createHandler func(router chi.Router) http.Handler) {
	apiRouter := chi.NewRouter()
	setMiddlewares(apiRouter)

	rootRouter := chi.NewRouter()
	// we are mounting all APIs under /api path
	rootRouter.Mount("/api/v1", createHandler(apiRouter))
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
	router.Use(logs.NewStructuredLogger(log.Logger))
	router.Use(middleware.Recoverer)

	router.Use(cors.AllowAll().Handler)

	router.Use(
		middleware.SetHeader("X-Content-Type-Options", "nosniff"),
		middleware.SetHeader("X-Frame-Options", "deny"),
	)
	router.Use(middleware.NoCache)
}

func RunJob(jobName string, anApp app.Application) {
	job := ports.NewJob(anApp)
	switch jobName {
	case "daily-cancel-expired-holds":
		job.CancelExpiredHolds(time.Now())
	case "daily-mark-overdue-checkouts":
		job.MarkOverdueCheckouts(time.Now())
	default:
		panic(fmt.Sprintf("Not support job=%s", jobName))
	}
}
