package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/resterrors"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateCdaAgroCreditoPrivadoService(t *testing.T) {

	repository, service, _, _ := InitServiceTest(t)

	t.Run("when_sending_a_valid_domain_returns_success", func(t *testing.T) {

		agroDomain := []domain.CdaAgroCreditoPrivadoDomain{}
		err := faker.FakeData(&agroDomain)
		assert.Nil(t, err)

		repository.EXPECT().CreateManyCdaAgroCreditoPrivadoRepository(agroDomain).Return(
			nil,
		)
		err = service.CreateCdaAgroCreditoPrivadoService(agroDomain)
		assert.Nil(t, err)
	})

	t.Run("return_error_from_repository", func(t *testing.T) {

		agroDomain := []domain.CdaAgroCreditoPrivadoDomain{}

		repository.EXPECT().CreateManyCdaAgroCreditoPrivadoRepository(agroDomain).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		err := service.CreateCdaAgroCreditoPrivadoService(agroDomain)
		assert.NotNil(t, err)

	})
}
