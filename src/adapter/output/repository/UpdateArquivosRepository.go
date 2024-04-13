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
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ur *fundosRepository) UpdateArquivosRepository(
	arquivosDomain domain.ArquivosDomain,
) *resterrors.RestErr {
	logger.Info("Init UpdateArquivosRepository", "sincronizarFundos")

	collection := ur.databaseConnection.Collection(env.GetCollectionArquivos())

	arquivosEntity := &entity.ArquivosEntity{}
	copier.Copy(arquivosEntity, arquivosDomain)

	idHex, _ := primitive.ObjectIDFromHex(arquivosDomain.Id)
	filter := bson.D{{Key: "_id", Value: idHex}}
	update := bson.D{{Key: "$set", Value: arquivosEntity}}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	logger.Info(
		fmt.Sprintf("Domain: %v", arquivosDomain),
		"sincronizarFundos")

	if err != nil {
		logger.Error("Error trying to create user", err, "sincronizarFundos")
		return resterrors.NewInternalServerError(err.Error())
	}

	logger.Info("UpdateArquivosRepository executed successfully", "sincronizarFundos")

	return nil
}
