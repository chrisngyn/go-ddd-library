package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	watermillSQL "github.com/ThreeDotsLabs/watermill-sql/pkg/sql"
	"github.com/ThreeDotsLabs/watermill/components/forwarder"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"

	"github.com/chiennguyen196/go-library/internal/common/database"
	"github.com/chiennguyen196/go-library/internal/common/logs"
	"github.com/chiennguyen196/go-library/internal/common/server"
)

const (
	kafkaTopic = "catalogue-events"
)

func main() {
	logs.Init()

	severToRun := strings.ToLower(os.Args[1])

	db := database.NewSqlDB()

	switch severToRun {
	case "http":
		httpServer := HttpServer{db: NewDB(db)}
		server.RunHTTPServer("/api/v1", func(router chi.Router) http.Handler {
			return HandlerFromMux(httpServer, router)
		})
	case "forwarder":
		RunForwarder(db)
	default:
		panic(fmt.Sprintf("Not support serverToRun=%s", severToRun))
	}

}

func RunForwarder(db *sql.DB) {
	sqlSubscriber, err := watermillSQL.NewSubscriber(
		db,
		watermillSQL.SubscriberConfig{
			SchemaAdapter:    watermillSQL.DefaultPostgreSQLSchema{},
			OffsetsAdapter:   watermillSQL.DefaultPostgreSQLOffsetsAdapter{},
			InitializeSchema: true,
		},
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		log.Fatal().Err(err).Msg("Fail to create sql subscriber")
	}

	kafkaPublisher, err := kafka.NewPublisher(
		kafka.PublisherConfig{
			Brokers:   []string{"localhost:9092"},
			Marshaler: kafka.DefaultMarshaler{},
		},
		watermill.NewStdLogger(false, false),
	)

	if err != nil {
		log.Fatal().Err(err).Msg("Fail to create kafka publisher")
	}

	fwd, err := forwarder.NewForwarder(sqlSubscriber, kafkaPublisher, watermill.NewStdLogger(false, false), forwarder.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("Fail to create forwarder")
	}

	if err := fwd.Run(context.Background()); err != nil {
		log.Fatal().Err(err).Msg("Run forwarder fail")
	}

	return
}
