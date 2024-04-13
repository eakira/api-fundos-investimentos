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

func (ur *fundosRepository) CreateManyExtratoRepository(
	domain []domain.ExtratoDomain,
) *resterrors.RestErr {

	logger.Info("Init CreateExtratoRepository repository", "sincronizarFundos")

	collection := ur.databaseConnection.Collection(env.GetCollectionExtrato())

	entity := []entity.ExtratoEntity{}
	copier.Copy(&entity, &domain)

	dados := make([]interface{}, len(entity))
	copier.Copy(&dados, &entity)

	_, err := collection.InsertMany(context.Background(), dados)
	if err != nil {
		logger.Error("Error trying to create CreateExtratoRepository", err, "sincronizarFundos")
		return resterrors.NewInternalServerError(err.Error())
	}

	logger.Info("CreateExtratoRepository executed successfully", "sincronizarFundos")

	return nil
}
