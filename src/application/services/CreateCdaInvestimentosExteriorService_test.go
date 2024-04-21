package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/resterrors"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateCdaInvestimentosExteriorService(t *testing.T) {

	repository, service := InitServiceTest(t)

	t.Run("when_sending_a_valid_domain_returns_success", func(t *testing.T) {

		investimentosExteriorDomain := []domain.CdaInvestimentosExteriorDomain{}
		err := faker.FakeData(&investimentosExteriorDomain)
		assert.Nil(t, err)

		repository.EXPECT().CreateManyCdaInvestimentosExteriorRepository(investimentosExteriorDomain).Return(
			nil,
		)
		err = service.CreateCdaInvestimentosExteriorService(investimentosExteriorDomain)
		assert.Nil(t, err)
	})

	t.Run("return_error_from_repository", func(t *testing.T) {

		investimentosExteriorDomain := []domain.CdaInvestimentosExteriorDomain{}

		repository.EXPECT().CreateManyCdaInvestimentosExteriorRepository(investimentosExteriorDomain).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		err := service.CreateCdaInvestimentosExteriorService(investimentosExteriorDomain)
		assert.NotNil(t, err)

	})
}
