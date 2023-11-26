package answersheet_usecase

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
	"picket/src/entities"
	answersheet_struct "picket/src/features/answersheet/struct"
)

type IRepository interface {
	Create(ctx context.Context, event *entities.Event) error
	GetLatestEvent(ctx context.Context, userId int, testId int) ([]entities.Event, error)
	GetLatestStartEvent(ctx context.Context, userId int, testId int) (*entities.Event, error)
	GetLatestEventWithLimit(ctx context.Context, userId int, testId int, limit int) ([]entities.Event, error)
	FindAnswerByUserIdAndTestId(ctx context.Context, userId int, testId int, sessionId string) ([]entities.Event, error)
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
	//latest, err := u.repository.GetLatestEventWithLimit(ctx, input.Payload.UserId, input.Payload.TestId, 1)
	//if err != nil {
	//	return err
	//}
	//if len(latest) == 1 && latest[0].Event == entities.START {
	//	return errpkg.Answersheet.UserStarted
	//}
	session := uuid.New()
	e := entities.Event{
		UserId:    input.Payload.UserId,
		TestId:    input.Payload.TestId,
		CreatedAt: input.Payload.CreatedAt,
		UpdatedAt: input.Payload.UpdatedAt,
		Event:     input.Payload.Event,
		Id:        primitive.NewObjectID(),
		Session:   session.String(),
	}
	err := u.repository.Create(ctx, &e)
	zap.S().Info(input.Payload.CreatedAt.Format("15:04:05 02/01/2006"))
	if err != nil {
		return err
	}
	zap.S().Info(e.Id)

	return nil
}
