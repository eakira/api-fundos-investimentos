package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/resterrors"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateInformacaoComplementarFundoService(t *testing.T) {

	repository, service := InitServiceTest(t)

	t.Run("when_sending_a_valid_domain_returns_success", func(t *testing.T) {

		informacacoesFundoDomain := []domain.InformacoesFundoDomain{}
		err := faker.FakeData(&informacacoesFundoDomain)
		assert.Nil(t, err)

		repository.EXPECT().CreateManyInformacaoComplementarFundoRepository(informacacoesFundoDomain).Return(
			nil,
		)
		err = service.CreateInformacaoComplementarFundoService(informacacoesFundoDomain)
		assert.Nil(t, err)
	})

	t.Run("return_error_from_repository", func(t *testing.T) {

		informacacoesFundoDomain := []domain.InformacoesFundoDomain{}

		repository.EXPECT().CreateManyInformacaoComplementarFundoRepository(informacacoesFundoDomain).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		err := service.CreateInformacaoComplementarFundoService(informacacoesFundoDomain)
		assert.NotNil(t, err)

	})
}

func TestCreateInformacaoComplementarDivulgacaoService(t *testing.T) {

	repository, service := InitServiceTest(t)

	t.Run("when_sending_a_valid_domain_returns_success", func(t *testing.T) {

		informacoesDivulgacaoDomain := []domain.InformacoesDivulgacaoDomain{}
		err := faker.FakeData(&informacoesDivulgacaoDomain)
		assert.Nil(t, err)

		repository.EXPECT().CreateManyInformacaoComplementarDivulgacaoRepository(informacoesDivulgacaoDomain).Return(
			nil,
		)
		err = service.CreateInformacaoComplementarDivulgacaoService(informacoesDivulgacaoDomain)
		assert.Nil(t, err)
	})

	t.Run("return_error_from_repository", func(t *testing.T) {

		informacoesDivulgacaoDomain := []domain.InformacoesDivulgacaoDomain{}

		repository.EXPECT().CreateManyInformacaoComplementarDivulgacaoRepository(informacoesDivulgacaoDomain).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		err := service.CreateInformacaoComplementarDivulgacaoService(informacoesDivulgacaoDomain)
		assert.NotNil(t, err)

	})
}

func TestCreateInformacaoComplementarCotistaService(t *testing.T) {

	repository, service := InitServiceTest(t)

	t.Run("when_sending_a_valid_domain_returns_success", func(t *testing.T) {

		informacoesCotistaDomain := []domain.InformacoesCotistaDomain{}
		err := faker.FakeData(&informacoesCotistaDomain)
		assert.Nil(t, err)

		repository.EXPECT().CreateManyInformacaoComplementarCotistaRepository(informacoesCotistaDomain).Return(
			nil,
		)
		err = service.CreateInformacaoComplementarCotistaService(informacoesCotistaDomain)
		assert.Nil(t, err)
	})

	t.Run("return_error_from_repository", func(t *testing.T) {

		informacoesCotistaDomain := []domain.InformacoesCotistaDomain{}

		repository.EXPECT().CreateManyInformacaoComplementarCotistaRepository(informacoesCotistaDomain).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		err := service.CreateInformacaoComplementarCotistaService(informacoesCotistaDomain)
		assert.NotNil(t, err)

	})
}

func TestCreateInformacaoComplementarServicoPrestadoService(t *testing.T) {

	repository, service := InitServiceTest(t)

	t.Run("when_sending_a_valid_domain_returns_success", func(t *testing.T) {

		servicoPrestadoDomain := []domain.ServicoPrestadoDomain{}
		err := faker.FakeData(&servicoPrestadoDomain)
		assert.Nil(t, err)

		repository.EXPECT().CreateManyInformacaoComplementarServicoPrestadoRepository(servicoPrestadoDomain).Return(
			nil,
		)
		err = service.CreateInformacaoComplementarServicoPrestadoService(servicoPrestadoDomain)
		assert.Nil(t, err)
	})

	t.Run("return_error_from_repository", func(t *testing.T) {

		servicoPrestadoDomain := []domain.ServicoPrestadoDomain{}

		repository.EXPECT().CreateManyInformacaoComplementarServicoPrestadoRepository(servicoPrestadoDomain).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		err := service.CreateInformacaoComplementarServicoPrestadoService(servicoPrestadoDomain)
		assert.NotNil(t, err)

	})
}
