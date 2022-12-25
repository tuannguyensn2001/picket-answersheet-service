package answersheet_usecase

import (
	"context"
	"errors"
	"go.opentelemetry.io/otel/trace"
	"picket/src/entities"
)

func (u *usecase) CheckUserDoing(ctx context.Context, userId int, testId int) (bool, error) {
	span := trace.SpanFromContext(ctx)
	defer span.End()
	ctx, span = tracer.Start(ctx, "query mongo")
	list, err := u.repository.GetLatestEvent(ctx, userId, testId)
	span.End()
	if err != nil {
		span.RecordError(err)
		return false, err
	}
	span.AddEvent("get list success")
	if len(list) == 0 {
		span.RecordError(errors.New("list length == 0"))
		return false, nil
	}
	if len(list) == 1 {
		if list[0].Event == entities.START || list[0].Event == entities.DOING {
			span.RecordError(errors.New("error len list == 1"))
			return true, nil
		}
		return false, nil
	}

	first, second := list[0], list[1]

	if first.Event == entities.END && second.Event == entities.START {
		return false, nil
	}

	span.AddEvent("false")
	return true, nil
}
