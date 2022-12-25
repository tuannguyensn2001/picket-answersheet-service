package answersheet_usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
)

func (u *usecase) NotifyJobFail(ctx context.Context, jobId int, errFail error) error {
	jobFail := &kafka.Writer{
		Addr:                   kafka.TCP("localhost:9092"),
		Topic:                  "job-fail",
		AllowAutoTopicCreation: true,
		Balancer:               &kafka.LeastBytes{},
		BatchSize:              1,
	}

	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(map[string]interface{}{
		"job_id":        jobId,
		"error_message": errFail.Error(),
	})
	if err != nil {
		return err
	}

	err = jobFail.WriteMessages(ctx, kafka.Message{
		Value: b.Bytes(),
	})
	if err != nil {
		return err
	}

	return nil
}
