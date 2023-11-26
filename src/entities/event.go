package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Event struct {
	Id             primitive.ObjectID `bson:"_id,omitempty"`
	UserId         int                `bson:"user_id,omitempty"`
	TestId         int                `bson:"test_id,omitempty"`
	Event          string             `bson:"event,omitempty"`
	Session        string             `bson:"session,omitempty"`
	Answer         string             `bson:"answer,omitempty"`
	PreviousAnswer string             `bson:"previous_answer,omitempty"`
	CreatedAt      *time.Time         `bson:"created_at,omitempty"`
	UpdatedAt      *time.Time         `bson:"updated_at,omitempty"`

	QuestionId int `bson:"question_id,omitempty"`
}

const (
	START  = "START"
	DOING  = "DOING"
	END    = "END"
	ANSWER = "ANSWER"
)
