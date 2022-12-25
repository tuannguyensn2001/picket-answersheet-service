package answersheet_usecase

import (
	"context"
	"github.com/segmentio/kafka-go"
)

func (u *usecase) PushToDeadLetterQueue(ctx context.Context, value []byte) error {
	w := &kafka.Writer{
		Addr:                   kafka.TCP("localhost:9092"),
		Topic:                  "dead-letter-queue",
		AllowAutoTopicCreation: true,
		Balancer:               &kafka.LeastBytes{},
		BatchSize:              1,
	}

	err := w.WriteMessages(ctx, kafka.Message{
		Value: value,
	})
	if err != nil {
		return err
	}

	return nil

}
