package services

import (
	"api-fundos-investimentos/application/constants"
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/resterrors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func createArquivosDomainForSync() domain.ArquivosDomain {
	return domain.ArquivosDomain{
		Endereco:    "FI/CAD/DADOS/cad_fi.csv",
		TipoArquivo: "cadastros",
		Referencia:  "2023",
		Status:      constants.ENVIADO,
		Baixar:      true,
		Download:    false,
		Processado:  false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func TestQueueFundosSincronizarService(t *testing.T) {

	t.Run("when_sending_a_valid_domains_returns_success", func(t *testing.T) {
		repository, service, queue, _ := InitServiceTest(t)

		arquivosDomain := createArquivosDomainForSync()

		queue.EXPECT().Produce(gomock.Any()).Return(nil).AnyTimes()
		repository.EXPECT().CreateArquivosRepository(gomock.Any()).Return(&arquivosDomain, nil).AnyTimes()

		tipo := []string{
			"cadastros",
			"balancete",
			"cda",
			"extrato",
			"informacao-diaria",
			"informacoes-complementares",
			"lamina",
			"perfil-mensal",
		}
		for _, value := range tipo {
			err := service.QueueFundosSincronizarService(value, false)
			assert.Nil(t, err)
		}
	})

	t.Run("when_sending_a_valid_domain_returns_success", func(t *testing.T) {
		repository, service, queue, _ := InitServiceTest(t)

		arquivosDomain := createArquivosDomainForSync()

		queue.EXPECT().Produce(gomock.Any()).Return(nil).AnyTimes()
		repository.EXPECT().CreateArquivosRepository(gomock.Any()).DoAndReturn(func(arg domain.ArquivosDomain) {
			if arg.Endereco != arquivosDomain.Endereco ||
				arg.TipoArquivo != arquivosDomain.TipoArquivo ||
				arg.Referencia != arquivosDomain.Referencia ||
				arg.Status != arquivosDomain.Status ||
				arg.Baixar != arquivosDomain.Baixar ||
				arg.Download != arquivosDomain.Download ||
				arg.Processado != arquivosDomain.Processado {
				t.Errorf("Os campos não correspondem")
			}
		}).Return(&arquivosDomain, nil)

		err := service.QueueFundosSincronizarService("cadastros", true)
		assert.Nil(t, err)
	})

	t.Run("when_sending_a_invalid_domain_returns_error", func(t *testing.T) {
		repository, service, queue, _ := InitServiceTest(t)

		arquivosDomain := createArquivosDomainForSync()

		queue.EXPECT().Produce(gomock.Any()).Return(nil).AnyTimes()
		repository.EXPECT().CreateArquivosRepository(gomock.Any()).DoAndReturn(func(arg domain.ArquivosDomain) {
			if arg.Endereco != arquivosDomain.Endereco ||
				arg.TipoArquivo != arquivosDomain.TipoArquivo ||
				arg.Referencia != arquivosDomain.Referencia ||
				arg.Status != arquivosDomain.Status ||
				arg.Baixar == false ||
				arg.Download != arquivosDomain.Download ||
				arg.Processado != arquivosDomain.Processado {
				t.Errorf("Os campos não correspondem")
			}
		}).Return(&arquivosDomain, nil)

		err := service.QueueFundosSincronizarService("cadastros", true)
		assert.Nil(t, err)
	})

	// Enviando cadastro de vez cadastros retorna um error
	t.Run("when_sending_a_invalid_type_return_error", func(t *testing.T) {
		repository, service, queue, _ := InitServiceTest(t)

		arquivosDomain := createArquivosDomainForSync()

		queue.EXPECT().Produce(gomock.Any()).Return(nil).AnyTimes()
		repository.EXPECT().CreateArquivosRepository(gomock.Any()).Return(&arquivosDomain, nil).AnyTimes()

		err := service.QueueFundosSincronizarService("cadastro", false)
		assert.NotNil(t, err)

	})

	t.Run("when_sending_a_invalid_repository_return_error", func(t *testing.T) {
		repository, service, queue, _ := InitServiceTest(t)

		queue.EXPECT().Produce(gomock.Any()).Return(nil).AnyTimes()

		repository.EXPECT().CreateArquivosRepository(gomock.Any()).Return(
			nil,
			resterrors.NewInternalServerError("Erro pra teste"),
		).AnyTimes()

		err := service.QueueFundosSincronizarService("cadastros", false)
		assert.NotNil(t, err)

	})

	t.Run("when_sending_a_invalid_produce_return_error", func(t *testing.T) {
		repository, service, queue, _ := InitServiceTest(t)

		arquivosDomain := createArquivosDomainForSync()

		queue.EXPECT().Produce(gomock.Any()).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		).AnyTimes()

		repository.EXPECT().CreateArquivosRepository(gomock.Any()).Return(&arquivosDomain, nil).AnyTimes()

		err := service.QueueFundosSincronizarService("cadastros", false)
		assert.NotNil(t, err)

	})

}
