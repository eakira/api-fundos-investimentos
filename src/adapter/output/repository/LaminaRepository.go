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

func (ur *fundosRepository) CreateManyLaminaRepository(
	domain []domain.LaminaDomain,
) *resterrors.RestErr {

	logger.Info("Init CreateManyLaminaRepository", "sincronizarFundos")

	collection := ur.databaseConnection.Collection(env.GetCollectionLamina())

	entity := []entity.LaminaEntity{}
	copier.Copy(&entity, &domain)

	dados := make([]interface{}, len(entity))
	copier.Copy(&dados, &entity)

	_, err := collection.InsertMany(context.Background(), dados)
	if err != nil {
		logger.Error("Error trying to CreateManyLaminaRepository", err, "sincronizarFundos")
		return resterrors.NewInternalServerError(err.Error())
	}

	logger.Info("CreateManyLaminaRepository executed successfully", "sincronizarFundos")

	return nil
}

func (ur *fundosRepository) CreateManyLaminaCarteiraRepository(
	domain []domain.LaminaCarteiraDomain,
) *resterrors.RestErr {

	logger.Info("Init CreateManyLaminaCarteiraRepository", "sincronizarFundos")

	collection := ur.databaseConnection.Collection(env.GetCollectionLaminaCarteira())

	entity := []entity.LaminaCarteiraEntity{}
	copier.Copy(&entity, &domain)

	dados := make([]interface{}, len(entity))
	copier.Copy(&dados, &entity)

	_, err := collection.InsertMany(context.Background(), dados)
	if err != nil {
		logger.Error("Error trying to CreateManyLaminaCarteiraRepository", err, "sincronizarFundos")
		return resterrors.NewInternalServerError(err.Error())
	}

	logger.Info("CreateManyLaminaCarteiraRepository executed successfully", "sincronizarFundos")

	return nil
}

func (ur *fundosRepository) CreateManyLaminaRentabilidadeAnoRepository(
	domain []domain.LaminaRentabilidadeAnoDomain,
) *resterrors.RestErr {

	logger.Info("Init CreateManyLaminaRentabilidadeAnoRepository", "sincronizarFundos")

	collection := ur.databaseConnection.Collection(env.GetCollectionLaminaRentabilidadeAno())

	entity := []entity.LaminaRentabilidadeAnoEntity{}
	copier.Copy(&entity, &domain)

	dados := make([]interface{}, len(entity))
	copier.Copy(&dados, &entity)

	_, err := collection.InsertMany(context.Background(), dados)
	if err != nil {
		logger.Error("Error trying to CreateManyLaminaRentabilidadeAnoRepository", err, "sincronizarFundos")
		return resterrors.NewInternalServerError(err.Error())
	}

	logger.Info("CreateManyLaminaRentabilidadeAnoRepository executed successfully", "sincronizarFundos")

	return nil
}

func (ur *fundosRepository) CreateManyLaminaRentabilidadeMesRepository(
	domain []domain.LaminaRentabilidadeMesDomain,
) *resterrors.RestErr {

	logger.Info("Init CreateManyLaminaRentabilidadeMesRepository", "sincronizarFundos")

	collection := ur.databaseConnection.Collection(env.GetCollectionLaminaRentabilidadeMes())

	entity := []entity.LaminaRentabilidadeMesEntity{}
	copier.Copy(&entity, &domain)

	dados := make([]interface{}, len(entity))
	copier.Copy(&dados, &entity)

	_, err := collection.InsertMany(context.Background(), dados)
	if err != nil {
		logger.Error("Error trying to CreateManyLaminaRentabilidadeMesRepository", err, "sincronizarFundos")
		return resterrors.NewInternalServerError(err.Error())
	}

	logger.Info("CreateManyLaminaRentabilidadeMesRepository executed successfully", "sincronizarFundos")

	return nil
}
