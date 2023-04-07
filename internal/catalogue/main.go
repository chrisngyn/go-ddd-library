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
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
	"github.com/alexdrl/zerowater"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"

	"github.com/chiennguyen196/go-library/internal/common/database"
	"github.com/chiennguyen196/go-library/internal/common/logs"
	"github.com/chiennguyen196/go-library/internal/common/server"
)

// Catalogue context is quite simple. It's just a bunch of CRUD APIs.
// So we don't need to apply hexagon architecture into it because it seemed so over-engineering.
// And apply simple architecture instead. If this context become more complicated, we can apply hexagon architecture into it later.

var (
	logger     watermill.LoggerAdapter = watermill.NopLogger{}
	mysqlTable                         = "events"
)

func main() {
	logs.Init()

	logger = zerowater.NewZerologLoggerAdapter(log.Logger.With().Str("component", "watermill").Logger())

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
	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		log.Fatal().Err(err).Msg("Fail to create router")
	}

	router.AddPlugin(plugin.SignalsHandler)
	router.AddMiddleware(middleware.Recoverer)

	subscriber, err := watermillSQL.NewSubscriber(
		db,
		watermillSQL.SubscriberConfig{
			SchemaAdapter:    watermillSQL.DefaultPostgreSQLSchema{},
			OffsetsAdapter:   watermillSQL.DefaultPostgreSQLOffsetsAdapter{},
			InitializeSchema: true,
		},
		logger,
	)
	if err != nil {
		log.Fatal().Err(err).Msg("Fail to create subscriber")
	}

	brokers := strings.Split(os.Getenv("KAFKA_BROKERS"), ",")
	if len(brokers) == 0 {
		panic("missing KAFKA_BROKERS")
	}

	publisher, err := kafka.NewPublisher(
		kafka.PublisherConfig{
			Brokers:   brokers,
			Marshaler: kafka.DefaultMarshaler{},
		},
		logger,
	)
	if err != nil {
		log.Fatal().Err(err).Msg("Fail to create publisher")
	}

	kafkaTopic := os.Getenv("KAFKA_CATALOGUE_EVENTS_TOPIC")
	if kafkaTopic == "" {
		panic("missing KAFKA_CATALOGUE_EVENTS_TOPIC")
	}

	router.AddHandler(
		"mysql-to-kafka",
		mysqlTable,
		subscriber,
		kafkaTopic,
		publisher,
		func(msg *message.Message) ([]*message.Message, error) {
			log.Info().Msgf("Forward event with uuid: %s", msg.UUID)
			return []*message.Message{msg}, nil
		},
	)

	if err := router.Run(context.Background()); err != nil {
		log.Fatal().Err(err).Msg("Error run router")
	}

}
