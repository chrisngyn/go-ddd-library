package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"

	"github.com/chiennguyen196/go-library/internal/common/logs"
	"github.com/chiennguyen196/go-library/internal/common/server"
	"github.com/chiennguyen196/go-library/internal/lending/app"
	"github.com/chiennguyen196/go-library/internal/lending/ports"
	"github.com/chiennguyen196/go-library/internal/lending/service"
)

func main() {
	logs.Init()

	severToRun := strings.ToLower(os.Args[1])

	anApp, cleanFn := service.NewApplication()
	defer cleanFn()

	switch severToRun {
	case "http":
		server.RunHTTPServer("/api/v1/", func(router chi.Router) http.Handler {
			return ports.HandlerFromMux(ports.NewHttpServer(anApp), router)
		})
	case "job":
		jobName := strings.ToLower(os.Args[2])
		RunJob(jobName, anApp)
	case "consumer":
		RunConsumer(anApp)
	default:
		panic(fmt.Sprintf("Not support serverToRun=%s", severToRun))
	}
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

func RunConsumer(anApp app.Application) {
	consumer := ports.NewKafkaConsumer(anApp)
	if err := consumer.ConsumeCatalogueEvents(); err != nil {
		log.Fatal().Err(err).Msg("Fail to consume message")
	}
}
