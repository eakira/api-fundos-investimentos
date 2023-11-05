package repository

import (
	"api-fundos-investimentos/application/port/output"

	"go.mongodb.org/mongo-driver/mongo"
)

func NewFundosRepository(
	database *mongo.Database,
) output.FundosPort {
	return &fundosRepository{
		database,
	}
}

type fundosRepository struct {
	databaseConnection *mongo.Database
}
