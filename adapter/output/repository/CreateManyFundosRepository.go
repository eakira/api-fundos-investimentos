package repository

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/env"
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/configuration/resterrors"
	"context"
	"fmt"

	"github.com/jinzhu/copier"
)

func (ur *fundosRepository) CreateManyFundosRepository(
	fundosDomain []domain.FundosDomain,
) *resterrors.RestErr {

	logger.Info("Init CreateManyFundosRepository", "createFundos")

	collection := ur.databaseConnection.Collection(env.GetCollectionFundos())

	entity := make([]interface{}, len(fundosDomain))

	copier.Copy(&entity, &fundosDomain)
	fmt.Println(fundosDomain)

	_, err := collection.InsertMany(context.Background(), entity)

	if err != nil {
		logger.Error("Error CreateManyFundosRepository", err, "createFundos")
		return resterrors.NewInternalServerError(err.Error())
	}
	logger.Info("CreateFundosRepository executed successfully", "createFundos")

	return nil
}
