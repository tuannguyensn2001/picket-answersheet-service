package answersheet_usecase

import (
	"context"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"picket/src/entities"
	answersheet_struct "picket/src/features/answersheet/struct"
)

func (u *usecase) UserAnswer(ctx context.Context, input answersheet_struct.UserAnswerInput) error {
	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		return err
	}
	event, err := u.repository.GetLatestStartEvent(ctx, input.Payload.UserId, input.Payload.TestId)
	if err != nil {
		return err
	}
	sessionId := event.Session

	e := entities.Event{
		Id:             primitive.NewObjectID(),
		UserId:         input.Payload.UserId,
		TestId:         input.Payload.TestId,
		Event:          input.Payload.Event,
		Session:        sessionId,
		Answer:         input.Payload.Answer,
		CreatedAt:      input.Payload.CreatedAt,
		UpdatedAt:      input.Payload.UpdatedAt,
		PreviousAnswer: input.Payload.PreviousAnswer,
	}

	if err := u.repository.Create(ctx, &e); err != nil {
		return err
	}

	return nil
}
