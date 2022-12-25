package config

import (
	"github.com/go-redis/redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type IConfig interface {
	GetEnv() string
	GetGrpcAddress() string
	GetHttpAddress() string
	GetGoogleClientId() string
	GetGoogleClientSecret() string
	GetClientUrl() string
	GetDB() *gorm.DB
	GetSecretKey() string
	GetRedis() *redis.Client
	GetMongo() *mongo.Client
}

func (c config) GetGoogleClientId() string {
	return c.GoogleClientId
}

func (c config) GetGoogleClientSecret() string {
	return c.GoogleClientSecret
}

func (c config) GetClientUrl() string {
	return c.ClientUrl
}

func (c config) GetSecretKey() string {
	return c.secretKey
}

func (c config) GetRedis() *redis.Client {
	return c.redis
}

func (c config) GetMongo() *mongo.Client {
	return c.mongo
}
