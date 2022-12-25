package answersheet_transport

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	answersheet_struct "picket/src/features/answersheet/struct"
	errpkg "picket/src/packages/err"
	retrypkg "picket/src/packages/retry"
	answersheetpb "picket/src/pb/answer_sheet"
)

type IUsecase interface {
	StartTest(ctx context.Context, input answersheet_struct.StartTestInput) error
	CheckUserDoing(ctx context.Context, userId int, testId int) (bool, error)
	NotifyJobFail(ctx context.Context, jobId int, errFail error) error
	PushToDeadLetterQueue(ctx context.Context, value []byte) error
	NotifyJobSuccess(ctx context.Context, jobId int) error
}

type transport struct {
	usecase IUsecase
	answersheetpb.UnimplementedAnswerSheetServiceServer
}

func New(ctx context.Context, usecase IUsecase) *transport {
	t := &transport{usecase: usecase}
	go t.StartTest(ctx)
	return t
}

func (t *transport) StartTest(ctx context.Context) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "start-test",
		GroupID: "consumer-answersheet-service-1",
	})

	zap.S().Info("listen kafka")

	for {
		m, err := r.FetchMessage(ctx)
		if err != nil {
			zap.S().Error(err)
			continue
		}

		var input answersheet_struct.StartTestInput
		err = json.NewDecoder(bytes.NewBuffer(m.Value)).Decode(&input)
		if err != nil || input.JobId == 0 {
			err := t.usecase.PushToDeadLetterQueue(ctx, m.Value)
			if err == nil {
				r.CommitMessages(ctx, m)
			}
			zap.S().Error(err)
			continue
		}
		err = retrypkg.Do(func() error {
			return t.usecase.StartTest(ctx, input)
		}, retrypkg.Options{
			Attempt: 5,
		})
		if err != nil {
			err := t.usecase.NotifyJobFail(ctx, input.JobId, err)
			if err == nil {
				r.CommitMessages(ctx, m)
			}
			zap.S().Error(err)
			continue
		}

		go t.usecase.NotifyJobSuccess(ctx, input.JobId)

		if err := r.CommitMessages(ctx, m); err != nil {
			zap.S().Error(err)
		}
	}
	if err := r.Close(); err != nil {
		zap.S().Error(err)
	}
}

func (t *transport) CheckUserDoingTest(ctx context.Context, request *answersheetpb.CheckUserDoingTestRequest) (*answersheetpb.CheckUserDoingTestResponse, error) {

	check, err := t.usecase.CheckUserDoing(ctx, int(request.UserId), int(request.TestId))
	if err != nil {
		return nil, status.Error(codes.Internal, errpkg.Answersheet.UserDoingTest.Message)
	}

	return &answersheetpb.CheckUserDoingTestResponse{
		Check:   check,
		Message: "success",
	}, nil
}
