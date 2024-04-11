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

func (ur *fundosRepository) CreateManyInformaComplementarFundoRepository(
	domain []domain.InformacoesFundoDomain,
) *resterrors.RestErr {

	logger.Info("Init CreateManyInformaComplementarFundoRepository", "createBalacente")

	collection := ur.databaseConnection.Collection(env.GetCollectionBalancete())

	entity := []entity.BalacenteEntity{}
	copier.Copy(&entity, &domain)

	dados := make([]interface{}, len(entity))
	copier.Copy(&dados, &entity)

	_, err := collection.InsertMany(context.Background(), dados)
	if err != nil {
		logger.Error("Error trying to CreateManyInformaComplementarFundoRepository", err, "createBalacente")
		return resterrors.NewInternalServerError(err.Error())
	}

	logger.Info("CreateManyInformaComplementarFundoRepository executed successfully", "createBalacente")

	return nil
}

func (ur *fundosRepository) CreateManyInformaComplementarDivulgacaoRepository(
	domain []domain.InformacoesDivulgacaoDomain,
) *resterrors.RestErr {

	logger.Info("Init CreateManyInformaComplementarDivulgacaoRepository", "createBalacente")

	collection := ur.databaseConnection.Collection(env.GetCollectionBalancete())

	entity := []entity.BalacenteEntity{}
	copier.Copy(&entity, &domain)

	dados := make([]interface{}, len(entity))
	copier.Copy(&dados, &entity)

	_, err := collection.InsertMany(context.Background(), dados)
	if err != nil {
		logger.Error("Error trying to CreateManyInformaComplementarDivulgacaoRepository", err, "createBalacente")
		return resterrors.NewInternalServerError(err.Error())
	}

	logger.Info("CreateManyInformaComplementarDivulgacaoRepository executed successfully", "createBalacente")

	return nil
}

func (ur *fundosRepository) CreateManyInformaComplementarCotistaRepository(
	domain []domain.InformacoesCotistaDomain,
) *resterrors.RestErr {

	logger.Info("Init CreateManyInformaComplementarCotistaRepository", "createBalacente")

	collection := ur.databaseConnection.Collection(env.GetCollectionBalancete())

	entity := []entity.BalacenteEntity{}
	copier.Copy(&entity, &domain)

	dados := make([]interface{}, len(entity))
	copier.Copy(&dados, &entity)

	_, err := collection.InsertMany(context.Background(), dados)
	if err != nil {
		logger.Error("Error trying to CreateManyInformaComplementarCotistaRepository", err, "createBalacente")
		return resterrors.NewInternalServerError(err.Error())
	}

	logger.Info("CreateManyInformaComplementarCotistaRepository executed successfully", "createBalacente")

	return nil
}

func (ur *fundosRepository) CreateManyInformaComplementarServicoPrestadoRepository(
	domain []domain.ServicoPrestadoDomain,
) *resterrors.RestErr {

	logger.Info("Init CreateManyInformaComplementarServicoPrestadoRepository", "createBalacente")

	collection := ur.databaseConnection.Collection(env.GetCollectionBalancete())

	entity := []entity.BalacenteEntity{}
	copier.Copy(&entity, &domain)

	dados := make([]interface{}, len(entity))
	copier.Copy(&dados, &entity)

	_, err := collection.InsertMany(context.Background(), dados)
	if err != nil {
		logger.Error("Error trying to CreateManyInformaComplementarServicoPrestadoRepository", err, "createBalacente")
		return resterrors.NewInternalServerError(err.Error())
	}

	logger.Info("CreateManyInformaComplementarServicoPrestadoRepository executed successfully", "createBalacente")

	return nil
}
