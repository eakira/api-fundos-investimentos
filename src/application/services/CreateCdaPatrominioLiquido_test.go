package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/resterrors"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateCdaPatrominioLiquido(t *testing.T) {

	repository, service := InitServiceTest(t)

	t.Run("when_sending_a_valid_domain_returns_success", func(t *testing.T) {

		patrimonioLiquidoDomain := []domain.CdaPatrimonioLiquidoDomain{}
		err := faker.FakeData(&patrimonioLiquidoDomain)
		assert.Nil(t, err)

		repository.EXPECT().CreateManyCdaPatrimonioLiquidoRepository(patrimonioLiquidoDomain).Return(
			nil,
		)
		err = service.CreateCdaPatrominioLiquido(patrimonioLiquidoDomain)
		assert.Nil(t, err)
	})

	t.Run("return_error_from_repository", func(t *testing.T) {

		patrimonioLiquidoDomain := []domain.CdaPatrimonioLiquidoDomain{}

		repository.EXPECT().CreateManyCdaPatrimonioLiquidoRepository(patrimonioLiquidoDomain).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		err := service.CreateCdaPatrominioLiquido(patrimonioLiquidoDomain)
		assert.NotNil(t, err)

	})
}
