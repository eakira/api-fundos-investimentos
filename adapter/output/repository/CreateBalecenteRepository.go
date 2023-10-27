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

	logger.Info("Init createUser repository", "createArquivos")

	collection := ur.databaseConnection.Collection(env.GetCollectionBalancete())

	entity := &entity.BalacenteEntity{}
	copier.Copy(entity, balanceteDomain)

	_, err := collection.InsertOne(context.Background(), entity)
	if err != nil {
		return nil, resterrors.NewInternalServerError(err.Error())
	}

	if err != nil {
		logger.Error("Error trying to create user", err, "createArquivos")
		return nil, resterrors.NewInternalServerError(err.Error())
	}

	logger.Info("CreateUser repository executed successfully", "createArquivos")

	return &balanceteDomain, nil
}
