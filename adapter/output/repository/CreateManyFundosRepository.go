package repository

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/env"
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/configuration/resterrors"
	"context"
	"fmt"

	"github.com/jinzhu/copier"
)

func (ur *fundosRepository) CreateManyFundosRepository(
	fundosDomain []domain.FundosDomain,
) *resterrors.RestErr {

	logger.Info("Init CreateManyFundosRepository", "createFundos")

	collection := ur.databaseConnection.Collection(env.GetCollectionFundos())

	entity := make([]interface{}, len(fundosDomain))

	copier.Copy(&entity, &fundosDomain)
	fmt.Println(fundosDomain)
	logger.Info("Init 2 CreateManyFundosRepository", "createFundos")

	_, err := collection.InsertMany(context.Background(), entity)
	logger.Info("Init 3 CreateManyFundosRepository", "createFundos")

	if err != nil {
		logger.Error("Error CreateManyFundosRepository", err, "createFundos")
		return resterrors.NewInternalServerError(err.Error())
	}
	logger.Info("Init 4 CreateManyFundosRepository", "createFundos")

	if err != nil {
		logger.Error("Error trying to CreateFundosRepository", err, "createFundos")
		return resterrors.NewInternalServerError(err.Error())
	}
	logger.Info("Init 5 CreateManyFundosRepository", "createFundos")

	logger.Info("CreateFundosRepository executed successfully", "createFundos")

	return nil
}
