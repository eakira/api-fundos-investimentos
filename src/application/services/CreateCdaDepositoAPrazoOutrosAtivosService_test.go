package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/resterrors"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateCdaDepositoAPrazoOutrosAtivosService(t *testing.T) {

	repository, service := InitServiceTest(t)

	t.Run("when_sending_a_valid_domain_returns_success", func(t *testing.T) {

		depositoPrazoDomain := []domain.CdaDepositoAPrazoOutrosAtivosDomain{}
		err := faker.FakeData(&depositoPrazoDomain)
		assert.Nil(t, err)

		repository.EXPECT().CreateManyCdaDepositoAPrazoOutrosAtivosRepository(depositoPrazoDomain).Return(
			nil,
		)
		err = service.CreateCdaDepositoAPrazoOutrosAtivosService(depositoPrazoDomain)
		assert.Nil(t, err)
	})

	t.Run("return_error_from_repository", func(t *testing.T) {

		depositoPrazoDomain := []domain.CdaDepositoAPrazoOutrosAtivosDomain{}

		repository.EXPECT().CreateManyCdaDepositoAPrazoOutrosAtivosRepository(depositoPrazoDomain).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		err := service.CreateCdaDepositoAPrazoOutrosAtivosService(depositoPrazoDomain)
		assert.NotNil(t, err)

	})
}
