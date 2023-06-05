package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
	"time"
)

func NewMongoDatabase(configuration Config) *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoPoolMin, _ := strconv.Atoi(configuration.Get("MONGO_POOL_MIN"))
	mongoPoolMax, _ := strconv.Atoi(configuration.Get("MONGO_POOL_MAX"))
	mongoMaxIdleTime, _ := strconv.Atoi(configuration.Get("MONGO_MAX_IDLE_TIME_SECOND"))
	option := options.Client().
		ApplyURI(configuration.Get("MONGO_URI")).
		SetMinPoolSize(uint64(mongoPoolMin)).
		SetMaxPoolSize(uint64(mongoPoolMax)).
		SetMaxConnIdleTime(time.Duration(mongoMaxIdleTime) * time.Second)

	client, _ := mongo.NewClient(option)
	err := client.Connect(ctx)
	if err != nil {
		panic("cannot connect to database")
	}

	database := client.Database(configuration.Get("MONGO_DATABASE"))
	return database
}
