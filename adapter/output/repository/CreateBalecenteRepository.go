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

func (ur *fundosRepository) CreateBalecenteRepository(
	balanceteDomain domain.BalanceteDomain,
) (*domain.BalanceteDomain, *resterrors.RestErr) {

	logger.Info("Init CreateBalecenteRepository", "createBalacente")

	collection := ur.databaseConnection.Collection(env.GetCollectionBalancete())

	entity := &entity.BalacenteEntity{}
	copier.Copy(entity, balanceteDomain)

	_, err := collection.InsertOne(context.Background(), entity)
	if err != nil {
		return nil, resterrors.NewInternalServerError(err.Error())
	}

	if err != nil {
		logger.Error("Error trying to CreateBalecenteRepository", err, "createBalacente")
		return nil, resterrors.NewInternalServerError(err.Error())
	}

	logger.Info("CreateBalecenteRepository executed successfully", "createBalacente")

	return &balanceteDomain, nil
}
