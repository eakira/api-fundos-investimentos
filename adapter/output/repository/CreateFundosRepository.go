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

func (ur *fundosRepository) CreateFundosRepository(
	fundosDomain domain.FundosDomain,
) (*domain.FundosDomain, *resterrors.RestErr) {

	logger.Info("Init createUser repository", "createArquivos")

	collection := ur.databaseConnection.Collection(env.GetCollectionFundos())

	entity := &entity.FundosEntity{}
	copier.Copy(entity, fundosDomain)

	_, err := collection.InsertOne(context.Background(), entity)
	if err != nil {
		return nil, resterrors.NewInternalServerError(err.Error())
	}

	if err != nil {
		logger.Error("Error trying to create user", err, "createArquivos")
		return nil, resterrors.NewInternalServerError(err.Error())
	}

	logger.Info("CreateUser repository executed successfully", "createArquivos")

	return &fundosDomain, nil
}
