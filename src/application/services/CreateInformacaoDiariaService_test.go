package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/resterrors"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateInformacaoDiariaService(t *testing.T) {

	repository, service, _, _ := InitServiceTest(t)

	t.Run("when_sending_a_valid_domain_returns_success", func(t *testing.T) {

		informacaoDiariaDomain := []domain.InformacaoDiariaDomain{}
		err := faker.FakeData(&informacaoDiariaDomain)
		assert.Nil(t, err)

		repository.EXPECT().CreateManyInformacaoDiariaRepository(informacaoDiariaDomain).Return(
			nil,
		)
		err = service.CreateInformacaoDiariaService(informacaoDiariaDomain)
		assert.Nil(t, err)
	})

	t.Run("return_error_from_repository", func(t *testing.T) {

		informacaoDiariaDomain := []domain.InformacaoDiariaDomain{}

		repository.EXPECT().CreateManyInformacaoDiariaRepository(informacaoDiariaDomain).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		err := service.CreateInformacaoDiariaService(informacaoDiariaDomain)
		assert.NotNil(t, err)

	})
}
