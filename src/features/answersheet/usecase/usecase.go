package answersheet_usecase

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
	"picket/src/entities"
	answersheet_struct "picket/src/features/answersheet/struct"
)

type IRepository interface {
	Create(ctx context.Context, event *entities.Event) error
	GetLatestEvent(ctx context.Context, userId int, testId int) ([]entities.Event, error)
}

type usecase struct {
	repository IRepository
}

func New(repository IRepository) *usecase {
	return &usecase{repository: repository}
}

var tracer = otel.Tracer("answersheet_usecase")

func (u *usecase) StartTest(ctx context.Context, input answersheet_struct.StartTestInput) error {
	ctx, span := tracer.Start(ctx, "start test")
	defer span.End()
	e := entities.Event{
		UserId:    input.Payload.UserId,
		TestId:    input.Payload.TestId,
		CreatedAt: input.Payload.CreatedAt,
		UpdatedAt: input.Payload.UpdatedAt,
		Event:     input.Payload.Event,
		Id:        primitive.NewObjectID(),
	}
	err := u.repository.Create(ctx, &e)
	if err != nil {
		return err
	}
	zap.S().Info(e.Id)

	return nil
}
