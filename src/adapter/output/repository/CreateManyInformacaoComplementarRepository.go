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

func (ur *fundosRepository) CreateManyInformacaoComplementarFundoRepository(
	domain []domain.InformacoesFundoDomain,
) *resterrors.RestErr {

	logger.Info("Init CreateManyInformacaoComplementarFundoRepository", "sincronizarFundos")

	collection := ur.databaseConnection.Collection(env.GetCollectionInformacaooComplementarFundo())

	entity := []entity.InformacoesFundoEntity{}
	copier.Copy(&entity, &domain)

	dados := make([]interface{}, len(entity))
	copier.Copy(&dados, &entity)

	_, err := collection.InsertMany(context.Background(), dados)
	if err != nil {
		logger.Error("Error trying to CreateManyInformacaoComplementarFundoRepository", err, "sincronizarFundos")
		return resterrors.NewInternalServerError(err.Error())
	}

	logger.Info("CreateManyInformacaoComplementarFundoRepository executed successfully", "sincronizarFundos")

	return nil
}

func (ur *fundosRepository) CreateManyInformacaoComplementarDivulgacaoRepository(
	domain []domain.InformacoesDivulgacaoDomain,
) *resterrors.RestErr {

	logger.Info("Init CreateManyInformacaoComplementarDivulgacaoRepository", "sincronizarFundos")

	collection := ur.databaseConnection.Collection(env.GetCollectionInformacaoComplementarDivulgacao())

	entity := []entity.InformacoesDivulgacaoEntity{}
	copier.Copy(&entity, &domain)

	dados := make([]interface{}, len(entity))
	copier.Copy(&dados, &entity)

	_, err := collection.InsertMany(context.Background(), dados)
	if err != nil {
		logger.Error("Error trying to CreateManyInformacaoComplementarDivulgacaoRepository", err, "sincronizarFundos")
		return resterrors.NewInternalServerError(err.Error())
	}

	logger.Info("CreateManyInformacaoComplementarDivulgacaoRepository executed successfully", "sincronizarFundos")

	return nil
}

func (ur *fundosRepository) CreateManyInformacaoComplementarCotistaRepository(
	domain []domain.InformacoesCotistaDomain,
) *resterrors.RestErr {

	logger.Info("Init CreateManyInformacaoComplementarCotistaRepository", "sincronizarFundos")

	collection := ur.databaseConnection.Collection(env.GetCollectionInformacaoComplementarCotista())

	entity := []entity.InformacoesCotistaEntity{}
	copier.Copy(&entity, &domain)

	dados := make([]interface{}, len(entity))
	copier.Copy(&dados, &entity)

	_, err := collection.InsertMany(context.Background(), dados)
	if err != nil {
		logger.Error("Error trying to CreateManyInformacaoComplementarCotistaRepository", err, "sincronizarFundos")
		return resterrors.NewInternalServerError(err.Error())
	}

	logger.Info("CreateManyInformacaoComplementarCotistaRepository executed successfully", "sincronizarFundos")

	return nil
}

func (ur *fundosRepository) CreateManyInformacaoComplementarServicoPrestadoRepository(
	domain []domain.ServicoPrestadoDomain,
) *resterrors.RestErr {

	logger.Info("Init CreateManyInformacaoComplementarServicoPrestadoRepository", "sincronizarFundos")

	collection := ur.databaseConnection.Collection(env.GetCollectionInformacaoComplementarServicoPrestado())

	entity := []entity.ServicoPrestadoEntity{}
	copier.Copy(&entity, &domain)

	dados := make([]interface{}, len(entity))
	copier.Copy(&dados, &entity)

	_, err := collection.InsertMany(context.Background(), dados)
	if err != nil {
		logger.Error("Error trying to CreateManyInformacaoComplementarServicoPrestadoRepository", err, "sincronizarFundos")
		return resterrors.NewInternalServerError(err.Error())
	}

	logger.Info("CreateManyInformacaoComplementarServicoPrestadoRepository executed successfully", "sincronizarFundos")

	return nil
}
