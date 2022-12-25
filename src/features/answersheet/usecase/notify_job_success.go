package answersheet_usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
)

func (u *usecase) NotifyJobSuccess(ctx context.Context, jobId int) error {
	jobSuccess := &kafka.Writer{
		Addr:                   kafka.TCP("localhost:9092"),
		Topic:                  "job-success",
		AllowAutoTopicCreation: true,
		Balancer:               &kafka.LeastBytes{},
		BatchSize:              1,
	}

	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(map[string]interface{}{
		"job_id": jobId,
	})
	if err != nil {
		return err
	}

	err = jobSuccess.WriteMessages(ctx, kafka.Message{
		Value: b.Bytes(),
	})
	if err != nil {
		return err
	}

	return nil
}
