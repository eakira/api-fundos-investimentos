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

func (ur *fundosRepository) CreateManyCdaBlc1Repository(
	cdaDomain []domain.CdaBlc1Domain,
) *resterrors.RestErr {

	logger.Info("Init CreateManyBalecenteRepository", "createBalacente")

	collection := ur.databaseConnection.Collection(env.GetCollectionCdaBcl1())

	entity := []entity.BalacenteEntity{}
	copier.Copy(&entity, &cdaDomain)

	dados := make([]interface{}, len(entity))
	copier.Copy(&dados, &entity)

	_, err := collection.InsertMany(context.Background(), dados)
	if err != nil {
		logger.Error("Error trying to CreateManyBalecenteRepository", err, "createBalacente")
		return resterrors.NewInternalServerError(err.Error())
	}

	logger.Info("CreateManyBalecenteRepository executed successfully", "createBalacente")

	return nil
}

func (ur *fundosRepository) CreateManyCdaBlc2Repository(
	cdaDomain []domain.CdaBlc2Domain,
) *resterrors.RestErr {

	logger.Info("Init CreateManyBalecenteRepository", "createBalacente")

	collection := ur.databaseConnection.Collection(env.GetCollectionCdaBcl2())

	entity := []entity.BalacenteEntity{}
	copier.Copy(&entity, &cdaDomain)

	dados := make([]interface{}, len(entity))
	copier.Copy(&dados, &entity)

	_, err := collection.InsertMany(context.Background(), dados)
	if err != nil {
		logger.Error("Error trying to CreateManyBalecenteRepository", err, "createBalacente")
		return resterrors.NewInternalServerError(err.Error())
	}

	logger.Info("CreateManyBalecenteRepository executed successfully", "createBalacente")

	return nil
}

func (ur *fundosRepository) CreateManyCdaBlc3Repository(
	cdaDomain []domain.CdaBlc3Domain,
) *resterrors.RestErr {

	logger.Info("Init CreateManyBalecenteRepository", "createBalacente")

	collection := ur.databaseConnection.Collection(env.GetCollectionCdaBcl3())

	entity := []entity.BalacenteEntity{}
	copier.Copy(&entity, &cdaDomain)

	dados := make([]interface{}, len(entity))
	copier.Copy(&dados, &entity)

	_, err := collection.InsertMany(context.Background(), dados)
	if err != nil {
		logger.Error("Error trying to CreateManyBalecenteRepository", err, "createBalacente")
		return resterrors.NewInternalServerError(err.Error())
	}

	logger.Info("CreateManyBalecenteRepository executed successfully", "createBalacente")

	return nil
}

func (ur *fundosRepository) CreateManyCdaBlc4Repository(
	cdaDomain []domain.CdaBlc4Domain,
) *resterrors.RestErr {

	logger.Info("Init CreateManyBalecenteRepository", "createBalacente")

	collection := ur.databaseConnection.Collection(env.GetCollectionCdaBcl4())

	entity := []entity.BalacenteEntity{}
	copier.Copy(&entity, &cdaDomain)

	dados := make([]interface{}, len(entity))
	copier.Copy(&dados, &entity)

	_, err := collection.InsertMany(context.Background(), dados)
	if err != nil {
		logger.Error("Error trying to CreateManyBalecenteRepository", err, "createBalacente")
		return resterrors.NewInternalServerError(err.Error())
	}

	logger.Info("CreateManyBalecenteRepository executed successfully", "createBalacente")

	return nil
}

func (ur *fundosRepository) CreateManyCdaBlc5Repository(
	cdaDomain []domain.CdaBlc5Domain,
) *resterrors.RestErr {

	logger.Info("Init CreateManyBalecenteRepository", "createBalacente")

	collection := ur.databaseConnection.Collection(env.GetCollectionCdaBcl5())

	entity := []entity.BalacenteEntity{}
	copier.Copy(&entity, &cdaDomain)

	dados := make([]interface{}, len(entity))
	copier.Copy(&dados, &entity)

	_, err := collection.InsertMany(context.Background(), dados)
	if err != nil {
		logger.Error("Error trying to CreateManyBalecenteRepository", err, "createBalacente")
		return resterrors.NewInternalServerError(err.Error())
	}

	logger.Info("CreateManyBalecenteRepository executed successfully", "createBalacente")

	return nil
}

func (ur *fundosRepository) CreateManyCdaBlc6Repository(
	cdaDomain []domain.CdaBlc6Domain,
) *resterrors.RestErr {

	logger.Info("Init CreateManyBalecenteRepository", "createBalacente")

	collection := ur.databaseConnection.Collection(env.GetCollectionCdaBcl6())

	entity := []entity.BalacenteEntity{}
	copier.Copy(&entity, &cdaDomain)

	dados := make([]interface{}, len(entity))
	copier.Copy(&dados, &entity)

	_, err := collection.InsertMany(context.Background(), dados)
	if err != nil {
		logger.Error("Error trying to CreateManyBalecenteRepository", err, "createBalacente")
		return resterrors.NewInternalServerError(err.Error())
	}

	logger.Info("CreateManyBalecenteRepository executed successfully", "createBalacente")

	return nil
}

func (ur *fundosRepository) CreateManyCdaBlc7Repository(
	cdaDomain []domain.CdaBlc7Domain,
) *resterrors.RestErr {

	logger.Info("Init CreateManyBalecenteRepository", "createBalacente")

	collection := ur.databaseConnection.Collection(env.GetCollectionCdaBcl7())

	entity := []entity.BalacenteEntity{}
	copier.Copy(&entity, &cdaDomain)

	dados := make([]interface{}, len(entity))
	copier.Copy(&dados, &entity)

	_, err := collection.InsertMany(context.Background(), dados)
	if err != nil {
		logger.Error("Error trying to CreateManyBalecenteRepository", err, "createBalacente")
		return resterrors.NewInternalServerError(err.Error())
	}

	logger.Info("CreateManyBalecenteRepository executed successfully", "createBalacente")

	return nil
}

func (ur *fundosRepository) CreateManyCdaBlc8Repository(
	cdaDomain []domain.CdaBlc8Domain,
) *resterrors.RestErr {

	logger.Info("Init CreateManyBalecenteRepository", "createBalacente")

	collection := ur.databaseConnection.Collection(env.GetCollectionCdaBcl8())

	entity := []entity.BalacenteEntity{}
	copier.Copy(&entity, &cdaDomain)

	dados := make([]interface{}, len(entity))
	copier.Copy(&dados, &entity)

	_, err := collection.InsertMany(context.Background(), dados)
	if err != nil {
		logger.Error("Error trying to CreateManyBalecenteRepository", err, "createBalacente")
		return resterrors.NewInternalServerError(err.Error())
	}

	logger.Info("CreateManyBalecenteRepository executed successfully", "createBalacente")

	return nil
}
