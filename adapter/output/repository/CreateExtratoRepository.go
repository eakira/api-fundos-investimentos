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

func (ur *fundosRepository) CreateExtratoRepository(
	domain domain.ExtratoDomain,
) (*domain.ExtratoDomain, *resterrors.RestErr) {

	logger.Info("Init CreateExtratoRepository repository", "createExtrato")

	collection := ur.databaseConnection.Collection(env.GetCollectionExtrato())

	entity := &entity.ExtratoEntity{}
	copier.Copy(entity, domain)

	_, err := collection.InsertOne(context.Background(), entity)
	if err != nil {
		return nil, resterrors.NewInternalServerError(err.Error())
	}

	if err != nil {
		logger.Error("Error trying to create CreateExtratoRepository", err, "createExtrato")
		return nil, resterrors.NewInternalServerError(err.Error())
	}

	logger.Info("CreateExtratoRepository executed successfully", "createExtrato")

	return &domain, nil
}
