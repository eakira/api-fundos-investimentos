package mongodb

import (
	"api-fundos-investimentos/configuration/env"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDBConnection(
	ctx context.Context,
) (*mongo.Database, error) {
	mongodb_uri := env.GetMongoUrl()
	mongodb_database := env.GetMongoDataBase()

	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(mongodb_uri))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client.Database(mongodb_database), nil
}
