package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/resterrors"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateCdaSwapService(t *testing.T) {

	repository, service := InitServiceTest(t)

	t.Run("when_sending_a_valid_domain_returns_success", func(t *testing.T) {

		swapDomain := []domain.CdaSwapDomain{}
		err := faker.FakeData(&swapDomain)
		assert.Nil(t, err)

		repository.EXPECT().CreateManyCdaSwapRepository(swapDomain).Return(
			nil,
		)
		err = service.CreateCdaSwapService(swapDomain)
		assert.Nil(t, err)
	})

	t.Run("return_error_from_repository", func(t *testing.T) {

		swapDomain := []domain.CdaSwapDomain{}

		repository.EXPECT().CreateManyCdaSwapRepository(swapDomain).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		err := service.CreateCdaSwapService(swapDomain)
		assert.NotNil(t, err)

	})
}
