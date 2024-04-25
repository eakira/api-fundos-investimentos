package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/resterrors"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateCdaFiimConfidencialidade(t *testing.T) {

	repository, service, _, _ := InitServiceTest(t)

	t.Run("when_sending_a_valid_domain_returns_success", func(t *testing.T) {

		confidencialidadeDomain := []domain.CdaFiimConfidencialDomain{}
		err := faker.FakeData(&confidencialidadeDomain)
		assert.Nil(t, err)

		repository.EXPECT().CreateManyCdaFiimConfidencialRepository(confidencialidadeDomain).Return(
			nil,
		)
		err = service.CreateCdaFiimConfidencialidade(confidencialidadeDomain)
		assert.Nil(t, err)
	})

	t.Run("return_error_from_repository", func(t *testing.T) {

		confidencialidadeDomain := []domain.CdaFiimConfidencialDomain{}

		repository.EXPECT().CreateManyCdaFiimConfidencialRepository(confidencialidadeDomain).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		err := service.CreateCdaFiimConfidencialidade(confidencialidadeDomain)
		assert.NotNil(t, err)

	})
}
