package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/resterrors"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateLaminaRentabilidadeMesService(t *testing.T) {

	repository, service := InitServiceTest(t)

	t.Run("when_sending_a_valid_domain_returns_success", func(t *testing.T) {

		rentabilidadeMesDomain := []domain.LaminaRentabilidadeMesDomain{}
		err := faker.FakeData(&rentabilidadeMesDomain)
		assert.Nil(t, err)

		repository.EXPECT().CreateManyLaminaRentabilidadeMesRepository(rentabilidadeMesDomain).Return(
			nil,
		)
		err = service.CreateLaminaRentabilidadeMesService(rentabilidadeMesDomain)
		assert.Nil(t, err)
	})

	t.Run("return_error_from_repository", func(t *testing.T) {

		rentabilidadeMesDomain := []domain.LaminaRentabilidadeMesDomain{}

		repository.EXPECT().CreateManyLaminaRentabilidadeMesRepository(rentabilidadeMesDomain).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		err := service.CreateLaminaRentabilidadeMesService(rentabilidadeMesDomain)
		assert.NotNil(t, err)

	})
}
