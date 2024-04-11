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

func (ur *fundosRepository) CreateManyPerfilMensalRepository(
	domain []domain.PerfilMensalDomain,
) *resterrors.RestErr {

	logger.Info("Init CreateManyPerfilMensalRepository", "createPerfilMensal")

	collection := ur.databaseConnection.Collection(env.GetCollectionPerfilMensal())

	entity := []entity.PerfilMensalEntity{}
	copier.Copy(&entity, &domain)

	dados := make([]interface{}, len(entity))
	copier.Copy(&dados, &entity)

	_, err := collection.InsertMany(context.Background(), dados)
	if err != nil {
		logger.Error("Error trying to CreateManyPerfilMensalRepository", err, "createPerfilMensal")
		return resterrors.NewInternalServerError(err.Error())
	}

	logger.Info("CreateManyPerfilMensalRepository executed successfully", "createPerfilMensal")

	return nil
}
