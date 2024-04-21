package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/resterrors"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateCdaSelicService(t *testing.T) {

	repository, service := InitServiceTest(t)

	t.Run("when_sending_a_valid_domain_returns_success", func(t *testing.T) {

		selicDomain := []domain.CdaSelicDomain{}
		err := faker.FakeData(&selicDomain)
		assert.Nil(t, err)

		repository.EXPECT().CreateManyCdaSelicRepository(selicDomain).Return(
			nil,
		)
		err = service.CreateCdaSelicService(selicDomain)
		assert.Nil(t, err)
	})

	t.Run("return_error_from_repository", func(t *testing.T) {

		selicDomain := []domain.CdaSelicDomain{}

		repository.EXPECT().CreateManyCdaSelicRepository(selicDomain).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		err := service.CreateCdaSelicService(selicDomain)
		assert.NotNil(t, err)

	})
}
