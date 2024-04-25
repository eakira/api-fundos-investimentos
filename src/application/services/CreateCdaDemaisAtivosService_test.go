package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/resterrors"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateCdaDemaisAtivosService(t *testing.T) {

	repository, service, _, _ := InitServiceTest(t)

	t.Run("when_sending_a_valid_domain_returns_success", func(t *testing.T) {

		demaisAtivosDomain := []domain.CdaDemaisAtivosDomain{}
		err := faker.FakeData(&demaisAtivosDomain)
		assert.Nil(t, err)

		repository.EXPECT().CreateManyCdaDemaisAtivosRepository(demaisAtivosDomain).Return(
			nil,
		)
		err = service.CreateCdaDemaisAtivosService(demaisAtivosDomain)
		assert.Nil(t, err)
	})

	t.Run("return_error_from_repository", func(t *testing.T) {

		demaisAtivosDomain := []domain.CdaDemaisAtivosDomain{}

		repository.EXPECT().CreateManyCdaDemaisAtivosRepository(demaisAtivosDomain).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		err := service.CreateCdaDemaisAtivosService(demaisAtivosDomain)
		assert.NotNil(t, err)

	})
}
