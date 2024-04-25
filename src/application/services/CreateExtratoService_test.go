package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/resterrors"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateExtratoServiceService(t *testing.T) {

	repository, service, _, _ := InitServiceTest(t)

	t.Run("when_sending_a_valid_domain_returns_success", func(t *testing.T) {

		extratoDomain := []domain.ExtratoDomain{}
		err := faker.FakeData(&extratoDomain)
		assert.Nil(t, err)

		repository.EXPECT().CreateManyExtratoRepository(extratoDomain).Return(
			nil,
		)
		err = service.CreateExtratoService(extratoDomain)
		assert.Nil(t, err)
	})

	t.Run("return_error_from_repository", func(t *testing.T) {

		extratoDomain := []domain.ExtratoDomain{}

		repository.EXPECT().CreateManyExtratoRepository(extratoDomain).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		err := service.CreateExtratoService(extratoDomain)
		assert.NotNil(t, err)

	})
}
