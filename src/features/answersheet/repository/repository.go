package answersheet_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"picket/src/entities"
)

type repository struct {
	mongo *mongo.Client
}

func New(mongo *mongo.Client) *repository {
	return &repository{mongo: mongo}
}

func (r *repository) Create(ctx context.Context, event *entities.Event) error {
	collection := r.mongo.Database("picket").Collection("events")

	_, err := collection.InsertOne(ctx, event)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetLatestEvent(ctx context.Context, userId int, testId int) ([]entities.Event, error) {
	filter := bson.D{
		{
			"$and", bson.A{
				bson.D{{"user_id", userId}},
				bson.D{{"test_id", testId}},
			},
		},
	}
	opts := options.Find().SetLimit(2).SetSort(bson.D{{"_id", -1}})
	cursor, err := r.mongo.Database("picket").Collection("events").Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	result := make([]entities.Event, 0)

	err = cursor.All(ctx, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
