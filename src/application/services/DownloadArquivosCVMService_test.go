package services

import (
	"api-fundos-investimentos/adapter/output/model/response"
	"api-fundos-investimentos/application/constants"
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/env"
	"api-fundos-investimentos/configuration/resterrors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func createArquivosDomainForComparison() domain.ArquivosDomain {
	return domain.ArquivosDomain{
		Endereco:    "FI/CAD/DADOS/cad_fi.csv",
		TipoArquivo: "cadastros",
		Referencia:  "2023",
		Status:      "PROCESSANDO",
		Baixar:      true,
		Download:    false,
		Processado:  false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func TestDownloadArquivosCVMService(t *testing.T) {

	t.Run("when_sending_a_valid_domain_returns_success", func(t *testing.T) {
		repository, service, queue, externo := InitServiceTest(t)

		arquivosDomain := createArquivosDomainForComparison()
		externo.EXPECT().DownloadArquivosCVMPort(arquivosDomain.Endereco, arquivosDomain.Baixar).Return(
			[]string{arquivosDomain.Endereco},
			nil,
		)

		updateDomain := arquivosDomain
		updateDomain.Download = true
		updateDomain.Status = constants.FINALIZADO

		repository.EXPECT().UpdateArquivosRepository(updateDomain).Do(func(arg domain.ArquivosDomain) {
			if arg.Endereco != updateDomain.Endereco ||
				arg.TipoArquivo != updateDomain.TipoArquivo ||
				arg.Referencia != updateDomain.Referencia ||
				arg.Status != updateDomain.Status ||
				arg.Baixar != updateDomain.Baixar ||
				arg.Download != updateDomain.Download ||
				arg.Processado != updateDomain.Processado ||
				!arg.CreatedAt.Equal(updateDomain.CreatedAt) {
				t.Errorf("Os campos não correspondem")
			}
		}).Return(nil)

		arquivosDomain.Download = true
		arquivosDomain.Status = constants.DOWNLOAD

		repository.EXPECT().CreateArquivosRepository(gomock.Any()).Do(func(arg domain.ArquivosDomain) {
			if arg.Endereco != arquivosDomain.Endereco ||
				arg.TipoArquivo != arquivosDomain.TipoArquivo ||
				arg.Referencia != arquivosDomain.Referencia ||
				arg.Status != arquivosDomain.Status ||
				arg.Baixar != arquivosDomain.Baixar ||
				arg.Download != arquivosDomain.Download ||
				arg.Processado != arquivosDomain.Processado ||
				!arquivosDomain.CreatedAt.Equal(arquivosDomain.CreatedAt) {
				t.Errorf("Os campos não correspondem")
			}
		}).Return(&arquivosDomain, nil)

		queue.EXPECT().Produce(gomock.Any()).Do(func(arg response.FundosQueueResponse) {
			if arg.Topic != env.GetTopicProcessarArquivos() ||
				arg.Queue != "update-all" {
				t.Errorf("Os campos não correspondem")
			}

		}).Return(nil)

		err := service.DownloadArquivosCVMService(arquivosDomain)
		assert.Nil(t, err)
	})

	t.Run("when_sending_a_invalid_file_to_download_returns_error", func(t *testing.T) {
		repository, service, queue, externo := InitServiceTest(t)

		arquivosDomain := createArquivosDomainForComparison()

		externo.EXPECT().DownloadArquivosCVMPort(arquivosDomain.Endereco, arquivosDomain.Baixar).Return(
			nil,
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		updateDomain := arquivosDomain
		updateDomain.Download = true
		updateDomain.Status = constants.FINALIZADO

		repository.EXPECT().UpdateArquivosRepository(updateDomain).Return(nil)

		repository.EXPECT().CreateArquivosRepository(gomock.Any()).Return(&arquivosDomain, nil)

		queue.EXPECT().Produce(gomock.Any()).Return(nil)

		err := service.DownloadArquivosCVMService(arquivosDomain)
		assert.NotNil(t, err)
	})

	t.Run("when_sending_a_invalid_domain_returns_error", func(t *testing.T) {
		repository, service, queue, externo := InitServiceTest(t)

		arquivosDomain := createArquivosDomainForComparison()

		externo.EXPECT().DownloadArquivosCVMPort(arquivosDomain.Endereco, arquivosDomain.Baixar).Return(
			[]string{arquivosDomain.Endereco},
			nil,
		)

		updateDomain := arquivosDomain
		updateDomain.Download = true
		updateDomain.Status = constants.FINALIZADO

		repository.EXPECT().UpdateArquivosRepository(gomock.Any()).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		repository.EXPECT().CreateArquivosRepository(gomock.Any()).Return(&arquivosDomain, nil)

		queue.EXPECT().Produce(gomock.Any()).Return(nil)

		err2 := service.DownloadArquivosCVMService(arquivosDomain)
		assert.NotNil(t, err2)
	})

	t.Run("when_sending_a_invalid_repository_returns_error", func(t *testing.T) {
		repository, service, queue, externo := InitServiceTest(t)

		arquivosDomain := createArquivosDomainForComparison()

		externo.EXPECT().DownloadArquivosCVMPort(arquivosDomain.Endereco, arquivosDomain.Baixar).Return(
			[]string{arquivosDomain.Endereco},
			nil,
		)

		updateDomain := arquivosDomain
		updateDomain.Download = true
		updateDomain.Status = constants.FINALIZADO

		repository.EXPECT().UpdateArquivosRepository(gomock.Any()).Return(nil)

		repository.EXPECT().CreateArquivosRepository(gomock.Any()).Return(
			&arquivosDomain,
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		queue.EXPECT().Produce(gomock.Any()).Return(nil)

		err := service.DownloadArquivosCVMService(arquivosDomain)
		assert.NotNil(t, err)
	})

	t.Run("when_sending_a_invalid_produce_returns_error", func(t *testing.T) {
		repository, service, queue, externo := InitServiceTest(t)

		arquivosDomain := createArquivosDomainForComparison()

		externo.EXPECT().DownloadArquivosCVMPort(arquivosDomain.Endereco, arquivosDomain.Baixar).Return(
			[]string{arquivosDomain.Endereco},
			nil,
		)

		updateDomain := arquivosDomain
		updateDomain.Download = true
		updateDomain.Status = constants.FINALIZADO

		repository.EXPECT().UpdateArquivosRepository(gomock.Any()).Return(nil)

		repository.EXPECT().CreateArquivosRepository(gomock.Any()).Return(&arquivosDomain, nil)

		queue.EXPECT().Produce(gomock.Any()).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		err := service.DownloadArquivosCVMService(arquivosDomain)
		assert.NotNil(t, err)
	})

}
