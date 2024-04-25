package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/resterrors"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateCdaFundosInvestimentosService(t *testing.T) {

	repository, service, _, _ := InitServiceTest(t)

	t.Run("when_sending_a_valid_domain_returns_success", func(t *testing.T) {

		fundosInvestimentosDomain := []domain.CdaFundosInvestimentosDomain{}
		err := faker.FakeData(&fundosInvestimentosDomain)
		assert.Nil(t, err)

		repository.EXPECT().CreateManyCdaFundosInvestimentosRepository(fundosInvestimentosDomain).Return(
			nil,
		)
		err = service.CreateCdaFundosInvestimentosService(fundosInvestimentosDomain)
		assert.Nil(t, err)
	})

	t.Run("return_error_from_repository", func(t *testing.T) {

		fundosInvestimentosDomain := []domain.CdaFundosInvestimentosDomain{}

		repository.EXPECT().CreateManyCdaFundosInvestimentosRepository(fundosInvestimentosDomain).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		err := service.CreateCdaFundosInvestimentosService(fundosInvestimentosDomain)
		assert.NotNil(t, err)

	})
}
