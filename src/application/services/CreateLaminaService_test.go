package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/resterrors"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateLaminaService(t *testing.T) {

	repository, service := InitServiceTest(t)

	t.Run("when_sending_a_valid_domain_returns_success", func(t *testing.T) {

		laminaDomain := []domain.LaminaDomain{}
		err := faker.FakeData(&laminaDomain)
		assert.Nil(t, err)

		repository.EXPECT().CreateManyLaminaRepository(laminaDomain).Return(
			nil,
		)
		err = service.CreateLaminaService(laminaDomain)
		assert.Nil(t, err)
	})

	t.Run("return_error_from_repository", func(t *testing.T) {

		laminaDomain := []domain.LaminaDomain{}

		repository.EXPECT().CreateManyLaminaRepository(laminaDomain).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		err := service.CreateLaminaService(laminaDomain)
		assert.NotNil(t, err)

	})
}
