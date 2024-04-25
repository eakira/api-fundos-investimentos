package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/resterrors"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateFundosService(t *testing.T) {

	repository, service, _, _ := InitServiceTest(t)

	t.Run("when_sending_a_valid_domain_returns_success", func(t *testing.T) {

		fundosDomain := []domain.FundosDomain{}
		err := faker.FakeData(&fundosDomain)
		assert.Nil(t, err)

		repository.EXPECT().CreateManyFundosRepository(fundosDomain).Return(
			nil,
		)
		err = service.CreateFundosService(fundosDomain)
		assert.Nil(t, err)
	})

	t.Run("return_error_from_repository", func(t *testing.T) {

		fundosDomain := []domain.FundosDomain{}

		repository.EXPECT().CreateManyFundosRepository(fundosDomain).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		err := service.CreateFundosService(fundosDomain)
		assert.NotNil(t, err)

	})
}
