package repository

import (
	"api-fundos-investimentos/adapter/output/model/entity"
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/env"
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/configuration/resterrors"
	"context"

	"github.com/jinzhu/copier"
)

func (ur *fundosRepository) CreateManyFundosRepository(
	fundosDomain []domain.FundosDomain,
) *resterrors.RestErr {

	logger.Info("Init CreateManyFundosRepository", "createFundos")

	collection := ur.databaseConnection.Collection(env.GetCollectionFundos())

	entity := []entity.FundosEntity{}
	copier.Copy(&entity, &fundosDomain)

	dados := make([]interface{}, len(entity))
	copier.Copy(&dados, &entity)

	_, err := collection.InsertMany(context.Background(), dados)

	if err != nil {
		logger.Error("Error CreateManyFundosRepository", err, "createFundos")
		return resterrors.NewInternalServerError(err.Error())
	}
	logger.Info("CreateFundosRepository executed successfully", "createFundos")

	return nil
}
