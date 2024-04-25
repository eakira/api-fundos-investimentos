package services

import (
	"api-fundos-investimentos/application/constants"
	"api-fundos-investimentos/application/domain"
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

	repository, service, queue, externo := InitServiceTest(t)

	t.Run("when_sending_a_valid_domain_returns_success", func(t *testing.T) {

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
			nil,
		)
		queue.EXPECT().Produce(gomock.Any()).Return(
			nil,
		)

		err := service.DownloadArquivosCVMService(arquivosDomain)
		assert.Nil(t, err)
	})

}
