package repository

import (
	"api-fundos-investimentos/adapter/output/model/entity"
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/env"
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/configuration/resterrors"
	"context"
	"fmt"

	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ur *fundosRepository) CreateArquivosRepository(
	arquivosDomain domain.ArquivosDomain,
) (*domain.ArquivosDomain, *resterrors.RestErr) {
	logger.Info("Init CreateArquivosRepository", "createArquivos")

	collection := ur.databaseConnection.Collection(env.GetCollectionArquivos())

	arquivosEntity := &entity.ArquivosEntity{}
	copier.Copy(arquivosEntity, arquivosDomain)

	result, err := collection.InsertOne(context.Background(), arquivosEntity)
	if err != nil {
		return nil, resterrors.NewInternalServerError(err.Error())
	}

	logger.Info(
		fmt.Sprintf("Domain: %v", arquivosDomain),
		"createArquivos")

	if err != nil {
		logger.Error("Error trying to CreateArquivosRepository", err, "createArquivos")
		return nil, resterrors.NewInternalServerError(err.Error())
	}

	arquivosEntity.ID = result.InsertedID.(primitive.ObjectID)

	logger.Info("CreateArquivosRepository executed successfully", "createArquivos")

	arquivosDomainReturn := &domain.ArquivosDomain{}
	copier.Copy(arquivosDomainReturn, arquivosEntity)
	arquivosDomainReturn.Id = arquivosEntity.ID.Hex()

	return arquivosDomainReturn, nil
}
