package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/resterrors"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreatePerfilMensalService(t *testing.T) {

	repository, service, _, _ := InitServiceTest(t)

	t.Run("when_sending_a_valid_domain_returns_success", func(t *testing.T) {

		perfilMensalDomain := []domain.PerfilMensalDomain{}
		err := faker.FakeData(&perfilMensalDomain)
		assert.Nil(t, err)

		repository.EXPECT().CreateManyPerfilMensalRepository(perfilMensalDomain).Return(
			nil,
		)
		err = service.CreatePerfilMensalService(perfilMensalDomain)
		assert.Nil(t, err)
	})

	t.Run("return_error_from_repository", func(t *testing.T) {

		perfilMensalDomain := []domain.PerfilMensalDomain{}

		repository.EXPECT().CreateManyPerfilMensalRepository(perfilMensalDomain).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		err := service.CreatePerfilMensalService(perfilMensalDomain)
		assert.NotNil(t, err)

	})
}
