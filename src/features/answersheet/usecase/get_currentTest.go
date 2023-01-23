package answersheet_usecase

import (
	"context"
	"picket/src/entities"
)

func (u *usecase) GetCurrentTest(ctx context.Context, testId int, userId int) ([]entities.Event, error) {
	event, err := u.repository.GetLatestStartEvent(ctx, userId, testId)
	if err != nil {
		return nil, err
	}
	return u.repository.FindAnswerByUserIdAndTestId(ctx, userId, testId, event.Session)
}
