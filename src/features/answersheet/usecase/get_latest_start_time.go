package answersheet_usecase

import (
	"context"
	"time"
)

func (u *usecase) GetLatestStartTime(ctx context.Context, testId int, userId int) (*time.Time, error) {

	event, err := u.repository.GetLatestStartEvent(ctx, userId, testId)
	if err != nil {
		return nil, err
	}

	result := event.CreatedAt

	return result, nil
}
