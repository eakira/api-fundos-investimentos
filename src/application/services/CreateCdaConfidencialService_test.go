package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/resterrors"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateCdaConfidencialService(t *testing.T) {

	repository, service := InitServiceTest(t)

	t.Run("when_sending_a_valid_domain_returns_success", func(t *testing.T) {

		confidencialDomain := []domain.CdaConfidencialDomain{}
		err := faker.FakeData(&confidencialDomain)
		assert.Nil(t, err)

		repository.EXPECT().CreateManyCdaConfidencialRepository(confidencialDomain).Return(
			nil,
		)
		err = service.CreateCdaConfidencialService(confidencialDomain)
		assert.Nil(t, err)
	})

	t.Run("return_error_from_repository", func(t *testing.T) {

		confidencialDomain := []domain.CdaConfidencialDomain{}

		repository.EXPECT().CreateManyCdaConfidencialRepository(confidencialDomain).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		err := service.CreateCdaConfidencialService(confidencialDomain)
		assert.NotNil(t, err)

	})
}
