package answersheet_usecase

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func (u *usecase) GetLatestStartTime(ctx context.Context, testId int, userId int) (*time.Time, error) {

	event, err := u.repository.GetLatestStartEvent(ctx, userId, testId)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	}
	if err != nil && errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}

	result := event.CreatedAt

	//t, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	//convert := result.In(t)

	//zap.S().Info(convert.Format("15:04:05 02/01/2006"))
	return result, nil
}
