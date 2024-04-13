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

func (ur *fundosRepository) CreateManyCdaSelicRepository(
	cdaDomain []domain.CdaSelicDomain,
) *resterrors.RestErr {

	logger.Info("Init CreateManyBalecenteRepository", "sincronizarFundos")

	collection := ur.databaseConnection.Collection(env.GetCollectionCdaSelic())

	entity := []entity.CdaSelicEntity{}
	copier.Copy(&entity, &cdaDomain)

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

func (ur *fundosRepository) CreateManyCdaFundosInvestimentosRepository(
	cdaDomain []domain.CdaFundosInvestimentosDomain,
) *resterrors.RestErr {

	logger.Info("Init CreateManyBalecenteRepository", "sincronizarFundos")

	collection := ur.databaseConnection.Collection(env.GetCollectionCdaFundosInvestimentos())

	entity := []entity.CdaFundosInvestimentosEntity{}
	copier.Copy(&entity, &cdaDomain)

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

func (ur *fundosRepository) CreateManyCdaSwapRepository(
	cdaDomain []domain.CdaSwapDomain,
) *resterrors.RestErr {

	logger.Info("Init CreateManyBalecenteRepository", "sincronizarFundos")

	collection := ur.databaseConnection.Collection(env.GetCollectionCdaSwap())

	entity := []entity.CdaSwapEntity{}
	copier.Copy(&entity, &cdaDomain)

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

func (ur *fundosRepository) CreateManyCdaDemaisAtivosRepository(
	cdaDomain []domain.CdaDemaisAtivosDomain,
) *resterrors.RestErr {

	logger.Info("Init CreateManyBalecenteRepository", "sincronizarFundos")

	collection := ur.databaseConnection.Collection(env.GetCollectionCdaDemaisAtivos())

	entity := []entity.CdaDemaisAtivosEntity{}
	copier.Copy(&entity, &cdaDomain)

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

func (ur *fundosRepository) CreateManyCdaDepositoAPrazoOutrosAtivosRepository(
	cdaDomain []domain.CdaDepositoAPrazoOutrosAtivosDomain,
) *resterrors.RestErr {

	logger.Info("Init CreateManyBalecenteRepository", "sincronizarFundos")

	collection := ur.databaseConnection.Collection(env.GetCollectionCdaDepositoAPrazoOutrosAtivos())

	entity := []entity.CdaDepositoAPrazoOutrosAtivosEntity{}
	copier.Copy(&entity, &cdaDomain)

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

func (ur *fundosRepository) CreateManyCdaAgroCreditoPrivadoRepository(
	cdaDomain []domain.CdaAgroCreditoPrivadoDomain,
) *resterrors.RestErr {

	logger.Info("Init CreateManyBalecenteRepository", "sincronizarFundos")

	collection := ur.databaseConnection.Collection(env.GetCollectionCdaAgroCreditoPrivado())

	entity := []entity.CdaAgroCreditoPrivadoEntity{}
	copier.Copy(&entity, &cdaDomain)

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

func (ur *fundosRepository) CreateManyCdaInvestimentosExteriorRepository(
	cdaDomain []domain.CdaInvestimentosExteriorDomain,
) *resterrors.RestErr {

	logger.Info("Init CreateManyBalecenteRepository", "sincronizarFundos")

	collection := ur.databaseConnection.Collection(env.GetCollectionCdaInvestimentosExterior())

	entity := []entity.CdaInvestimentosExteriorEntity{}
	copier.Copy(&entity, &cdaDomain)

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

func (ur *fundosRepository) CreateManyCdaDemaisAtivosNaoCodificadosRepository(
	cdaDomain []domain.CdaDemaisAtivosNaoCodificadosDomain,
) *resterrors.RestErr {

	logger.Info("Init CreateManyBalecenteRepository", "sincronizarFundos")

	collection := ur.databaseConnection.Collection(env.GetCollectionCdaDemaisAtivosNaoCodificados())

	entity := []entity.CdaDemaisAtivosNaoCodificadosEntity{}
	copier.Copy(&entity, &cdaDomain)

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

func (ur *fundosRepository) CreateManyCdaConfidencialRepository(
	cdaDomain []domain.CdaConfidencialDomain,
) *resterrors.RestErr {

	logger.Info("Init CreateManyCdaConfidencialRepository", "sincronizarFundos")

	collection := ur.databaseConnection.Collection(env.GetCollectionCdaConfidencialidade())

	entity := []entity.CdaConfidencialEntity{}
	copier.Copy(&entity, &cdaDomain)

	dados := make([]interface{}, len(entity))
	copier.Copy(&dados, &entity)

	_, err := collection.InsertMany(context.Background(), dados)
	if err != nil {
		logger.Error("Error trying to CreateManyCdaConfidencialRepository", err, "sincronizarFundos")
		return resterrors.NewInternalServerError(err.Error())
	}

	logger.Info("CreateManyCdaConfidencialRepository executed successfully", "sincronizarFundos")

	return nil
}

func (ur *fundosRepository) CreateManyCdaFiimRepository(
	cdaDomain []domain.CdaFiimDomain,
) *resterrors.RestErr {

	logger.Info("Init CreateManyBalecenteRepository", "sincronizarFundos")

	collection := ur.databaseConnection.Collection(env.GetCollectionCdaFiim())

	entity := []entity.CdaFiimEntity{}
	copier.Copy(&entity, &cdaDomain)

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

func (ur *fundosRepository) CreateManyCdaFiimConfidencialRepository(
	cdaDomain []domain.CdaFiimConfidencialDomain,
) *resterrors.RestErr {

	logger.Info("Init CreateManyBalecenteRepository", "sincronizarFundos")

	collection := ur.databaseConnection.Collection(env.GetCollectionCdaFiimConfidencialidade())

	entity := []entity.CdaFiimConfidencialEntity{}
	copier.Copy(&entity, &cdaDomain)

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

func (ur *fundosRepository) CreateManyCdaPatrimonioLiquidoRepository(
	cdaDomain []domain.CdaPatrimonioLiquidoDomain,
) *resterrors.RestErr {

	logger.Info("Init CreateManyCdaPatrimonioLiquidoRepository", "sincronizarFundos")

	collection := ur.databaseConnection.Collection(env.GetCollectionCdaPatrominioLiquido())

	entity := []entity.CdaPatrimonioLiquidoEntity{}
	copier.Copy(&entity, &cdaDomain)

	dados := make([]interface{}, len(entity))
	copier.Copy(&dados, &entity)

	_, err := collection.InsertMany(context.Background(), dados)
	if err != nil {
		logger.Error("Error trying to CreateManyCdaPatrimonioLiquidoRepository", err, "sincronizarFundos")
		return resterrors.NewInternalServerError(err.Error())
	}

	logger.Info("CreateManyCdaPatrimonioLiquidoRepository executed successfully", "sincronizarFundos")

	return nil
}
