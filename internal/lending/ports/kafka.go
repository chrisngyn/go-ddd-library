package ports

import (
	"context"
	"encoding/json"
	"os"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/alexdrl/zerowater"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	commonErrors "github.com/chiennguyen196/go-library/internal/common/errors"
	"github.com/chiennguyen196/go-library/internal/lending/app"
	"github.com/chiennguyen196/go-library/internal/lending/app/command"
	"github.com/chiennguyen196/go-library/internal/lending/domain/book"
)

type KafkaConsumer struct {
	app        app.Application
	subscriber message.Subscriber
	kafkaTopic string
}

func NewKafkaConsumer(anApp app.Application) KafkaConsumer {
	logger := zerowater.NewZerologLoggerAdapter(log.Logger.With().Str("component", "watermill").Logger())
	if anApp == (app.Application{}) {
		panic("missing app")
	}

	saramaSubscriberConfig := kafka.DefaultSaramaSubscriberConfig()
	// equivalent of auto.offset.reset: earliest
	saramaSubscriberConfig.Consumer.Offsets.Initial = sarama.OffsetOldest

	brokers := strings.Split(os.Getenv("KAFKA_BROKERS"), ",")
	if len(brokers) == 0 {
		panic("missing KAFKA_BROKERS")
	}

	subscriber, err := kafka.NewSubscriber(
		kafka.SubscriberConfig{
			Brokers:               brokers,
			Unmarshaler:           kafka.DefaultMarshaler{},
			OverwriteSaramaConfig: saramaSubscriberConfig,
			ConsumerGroup:         "lending_consumer_group",
		},
		logger,
	)
	if err != nil {
		panic(err)
	}

	kafkaTopic := os.Getenv("KAFKA_CATALOGUE_EVENTS_TOPIC")
	if kafkaTopic == "" {
		panic("missing KAFKA_CATALOGUE_EVENTS_TOPIC")
	}

	return KafkaConsumer{
		app:        anApp,
		subscriber: subscriber,
		kafkaTopic: kafkaTopic,
	}

}

func (c KafkaConsumer) ConsumeCatalogueEvents() error {
	messages, err := c.subscriber.Subscribe(context.Background(), c.kafkaTopic)
	if err != nil {
		return errors.Wrap(err, "subscribe")
	}

	c.processMessage(messages)
	return nil
}

func (c KafkaConsumer) processMessage(messages <-chan *message.Message) {
	for msg := range messages {
		logger := log.With().Str("event_id", msg.UUID).Logger()
		msg.SetContext(logger.WithContext(msg.Context()))
		eventType := getEventType(msg)
		logger.Info().
			Str("event_type", eventType).
			Str("payload", string(msg.Payload)).
			Msg("Received event")

		var err error
		switch eventType {
		case "BookInstanceAdded":
			err = c.handleBookInstanceAddedEvent(msg)
		default:
			logger.Info().Msg("Have no handler, ignore")
		}

		if err != nil {
			logger.Error().Err(err).Msg("Process message fail")
		}

		if shouldNack(err) {
			msg.Ack()
		}

		msg.Ack()
	}
}

func (c KafkaConsumer) handleBookInstanceAddedEvent(msg *message.Message) error {
	var event BookInstanceAdded
	if err := json.Unmarshal(msg.Payload, &event); err != nil {
		return errors.Wrap(err, "unmarshal event")
	}

	cmd, err := toAddNewBookCommand(event.BookID, event.BookType, event.LibraryBranchID)
	if err != nil {
		return errors.Wrap(err, "to command")
	}

	if err := c.app.Commands.AddNewBook.Handle(msg.Context(), cmd); err != nil {
		return errors.Wrap(err, "handle command")
	}
	return nil
}

func getEventType(msg *message.Message) string {
	return msg.Metadata.Get("eventType")
}

type BookInstanceAdded struct {
	ISBN            string    `json:"isbn"`
	BookID          uuid.UUID `json:"bookID"`
	BookType        string    `json:"bookType"`
	LibraryBranchID uuid.UUID `json:"libraryBranchID"`
	When            time.Time `json:"when"`
}

func toAddNewBookCommand(bookID uuid.UUID, bookType string, libraryBranchID uuid.UUID) (command.AddNewBookCommand, error) {
	domainBookType, err := toDomainBookType(bookType)
	if err != nil {
		return command.AddNewBookCommand{}, err
	}
	return command.AddNewBookCommand{
		BookID:          bookID,
		BookType:        domainBookType,
		LibraryBranchID: libraryBranchID,
	}, nil
}

func toDomainBookType(bookType string) (book.Type, error) {
	switch bookType {
	case "Restricted":
		return book.TypeRestricted, nil
	case "Circulating":
		return book.TypeCirculating, nil
	default:
		return book.Type{}, commonErrors.NewIncorrectInputError("invalid-book-type", "invalid book type")
	}
}

func shouldNack(err error) bool {
	if err == nil {
		return false
	}
	var slugErr commonErrors.SlugError
	if errors.As(err, &slugErr) {
		if slugErr.ErrorType() == commonErrors.ErrorTypeIncorrectInput {
			return false
		}
	}
	return true
}
