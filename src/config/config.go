package config

import (
	"context"
	"github.com/go-redis/redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/gorm"
	"log"
)

type config struct {
	Env                string
	GrpcAddress        string
	HttpAddress        string
	GoogleClientId     string
	GoogleClientSecret string
	ClientUrl          string
	db                 *gorm.DB
	secretKey          string
	redis              *redis.Client
	mongo              *mongo.Client
}

func GetConfig() config {
	structure := bootstrap()
	result := config{
		Env:                structure.App.Env,
		GrpcAddress:        structure.App.GrpcAddress,
		HttpAddress:        structure.App.HttpAddress,
		GoogleClientId:     structure.OAuth2.Google.ClientId,
		GoogleClientSecret: structure.OAuth2.Google.ClientSecret,
		ClientUrl:          structure.Client.Url,
		secretKey:          structure.App.SecretKey,
	}

	rd := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	status := rd.Ping(context.TODO())
	if status.Err() != nil {
		log.Println("redis ping error", status.Err())
	}

	m, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(structure.Database.Mongo))
	if err != nil {
		log.Fatalln(err)
	}

	result.redis = rd
	result.mongo = m

	return result
}

func (c config) GetEnv() string {
	return c.Env
}

func (c config) GetGrpcAddress() string {
	return c.GrpcAddress
}

func (c config) GetHttpAddress() string {
	return c.HttpAddress
}

func (c config) GetDB() *gorm.DB {
	return c.db
}
