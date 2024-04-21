package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/resterrors"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateCdaFiimService(t *testing.T) {

	repository, service := InitServiceTest(t)

	t.Run("when_sending_a_valid_domain_returns_success", func(t *testing.T) {

		fiimDomain := []domain.CdaFiimDomain{}
		err := faker.FakeData(&fiimDomain)
		assert.Nil(t, err)

		repository.EXPECT().CreateManyCdaFiimRepository(fiimDomain).Return(
			nil,
		)
		err = service.CreateCdaFiimService(fiimDomain)
		assert.Nil(t, err)
	})

	t.Run("return_error_from_repository", func(t *testing.T) {

		fiimDomain := []domain.CdaFiimDomain{}

		repository.EXPECT().CreateManyCdaFiimRepository(fiimDomain).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		err := service.CreateCdaFiimService(fiimDomain)
		assert.NotNil(t, err)

	})
}
