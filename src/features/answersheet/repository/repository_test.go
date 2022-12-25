package answersheet_repository

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
)

func TestRepository_GetLatestEvent(t *testing.T) {
	m, _ := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017/picket"))
	repo := New(m)

	result, err := repo.GetLatestEvent(context.TODO(), 1, 1)

	assert.Nil(t, err)
	assert.NotNil(t, result)
}
