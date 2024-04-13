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

func (ur *fundosRepository) CreateManyBalecenteRepository(
	balanceteDomain []domain.BalanceteDomain,
) *resterrors.RestErr {

	logger.Info("Init CreateManyBalecenteRepository", "sincronizarFundos")

	collection := ur.databaseConnection.Collection(env.GetCollectionBalancete())

	entity := []entity.BalacenteEntity{}
	copier.Copy(&entity, &balanceteDomain)

	dados := make([]interface{}, len(entity))
	copier.Copy(&dados, &entity)

	_, err := collection.InsertMany(context.Background(), dados)
	if err != nil {
		logger.Error("Error trying to CreateManyBalecenteRepository", err, "sincronizarFundos")
		return resterrors.NewInternalServerError(err.Error())
	}

	logger.Info("CreateManyBalecenteRepository executed successfully", "sincronizarFundos")

	return nil
}
