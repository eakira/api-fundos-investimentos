package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/resterrors"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateLaminaRentabilidadeAnoService(t *testing.T) {

	repository, service, _, _ := InitServiceTest(t)

	t.Run("when_sending_a_valid_domain_returns_success", func(t *testing.T) {

		rentabilidadeAnoDomain := []domain.LaminaRentabilidadeAnoDomain{}
		err := faker.FakeData(&rentabilidadeAnoDomain)
		assert.Nil(t, err)

		repository.EXPECT().CreateManyLaminaRentabilidadeAnoRepository(rentabilidadeAnoDomain).Return(
			nil,
		)
		err = service.CreateLaminaRentabilidadeAnoService(rentabilidadeAnoDomain)
		assert.Nil(t, err)
	})

	t.Run("return_error_from_repository", func(t *testing.T) {

		rentabilidadeAnoDomain := []domain.LaminaRentabilidadeAnoDomain{}

		repository.EXPECT().CreateManyLaminaRentabilidadeAnoRepository(rentabilidadeAnoDomain).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		err := service.CreateLaminaRentabilidadeAnoService(rentabilidadeAnoDomain)
		assert.NotNil(t, err)

	})
}
