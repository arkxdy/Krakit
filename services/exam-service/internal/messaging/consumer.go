package messaging

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

type HandlerFunc func(ctx context.Context, msg kafka.Message) error

type Consumer struct {
	reader  *kafka.Reader
	handler HandlerFunc
}

func NewConsumer(broker, topic, groupID string, handler HandlerFunc) *Consumer {
	return &Consumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:  []string{broker},
			Topic:    topic,
			GroupID:  groupID,
			MinBytes: 1,
			MaxBytes: 10e6, // 10MB
		}),
		handler: handler,
	}
}

// Call this in a goroutine
func (c *Consumer) Start(ctx context.Context) error {
	for {
		msg, err := c.reader.FetchMessage(ctx)
		if err != nil {
			if errors.Is(err, context.Canceled) {
				return nil // clean shutdown
			}
			return fmt.Errorf("fetch: %w", err)
		}

		if err := c.handler(ctx, msg); err != nil {
			log.Printf("handler error: %v", err) // dead letter queue logic goes here
			continue
		}

		// Only commit after successful handling
		if err := c.reader.CommitMessages(ctx, msg); err != nil {
			return fmt.Errorf("commit: %w", err)
		}
	}
}

func (c *Consumer) Close() error {
	return c.reader.Close()
}
