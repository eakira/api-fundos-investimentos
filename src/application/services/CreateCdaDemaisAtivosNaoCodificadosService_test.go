package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/resterrors"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateCdaDemaisAtivosNaoCodificadosService(t *testing.T) {

	repository, service, _, _ := InitServiceTest(t)

	t.Run("when_sending_a_valid_domain_returns_success", func(t *testing.T) {

		naoCondificadosDomain := []domain.CdaDemaisAtivosNaoCodificadosDomain{}
		err := faker.FakeData(&naoCondificadosDomain)
		assert.Nil(t, err)

		repository.EXPECT().CreateManyCdaDemaisAtivosNaoCodificadosRepository(naoCondificadosDomain).Return(
			nil,
		)
		err = service.CreateCdaDemaisAtivosNaoCodificadosService(naoCondificadosDomain)
		assert.Nil(t, err)
	})

	t.Run("return_error_from_repository", func(t *testing.T) {

		naoCondificadosDomain := []domain.CdaDemaisAtivosNaoCodificadosDomain{}

		repository.EXPECT().CreateManyCdaDemaisAtivosNaoCodificadosRepository(naoCondificadosDomain).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		err := service.CreateCdaDemaisAtivosNaoCodificadosService(naoCondificadosDomain)
		assert.NotNil(t, err)

	})
}
