package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/resterrors"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateLaminaCarteiraService(t *testing.T) {

	repository, service, _, _ := InitServiceTest(t)

	t.Run("when_sending_a_valid_domain_returns_success", func(t *testing.T) {

		laminaCarteiraDomain := []domain.LaminaCarteiraDomain{}
		err := faker.FakeData(&laminaCarteiraDomain)
		assert.Nil(t, err)

		repository.EXPECT().CreateManyLaminaCarteiraRepository(laminaCarteiraDomain).Return(
			nil,
		)
		err = service.CreateLaminaCarteiraService(laminaCarteiraDomain)
		assert.Nil(t, err)
	})

	t.Run("return_error_from_repository", func(t *testing.T) {

		laminaCarteiraDomain := []domain.LaminaCarteiraDomain{}

		repository.EXPECT().CreateManyLaminaCarteiraRepository(laminaCarteiraDomain).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		err := service.CreateLaminaCarteiraService(laminaCarteiraDomain)
		assert.NotNil(t, err)

	})
}
