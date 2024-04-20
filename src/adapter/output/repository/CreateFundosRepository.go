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

	logger.Info("Init CreateFundosRepository", "sincronizarFundos")

	collection := ur.databaseConnection.Collection(env.GetCollectionFundos())

	entity := &entity.FundosEntity{}
	copier.Copy(entity, fundosDomain)

	_, err := collection.InsertOne(context.Background(), entity)

	if err != nil {
		logger.Error("Error trying to CreateFundosRepository", err, "sincronizarFundos")
		return nil, resterrors.NewInternalServerError(err.Error())
	}

	fundosDomainReturn := &domain.FundosDomain{}
	copier.Copy(fundosDomainReturn, entity)
	fundosDomainReturn.Id = entity.ID.Hex()

	logger.Info("CreateFundosRepository executed successfully", "sincronizarFundos")

	return fundosDomainReturn, nil
}
