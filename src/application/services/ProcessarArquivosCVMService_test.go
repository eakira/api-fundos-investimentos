package services

import (
	"api-fundos-investimentos/application/constants"
	"api-fundos-investimentos/application/domain"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func createArquivosDomainParaProcessamento() domain.ArquivosDomain {
	return domain.ArquivosDomain{
		Endereco:    "../../storage/cvm/mock/cad_fi.csv",
		TipoArquivo: "cadastros",
		Referencia:  "2023",
		Status:      "PROCESSANDO",
		Baixar:      true,
		Download:    true,
		Processado:  false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func TestProcessarArquivosCVMService(t *testing.T) {

	t.Run("when_sending_a_valid_domain_returns_success", func(t *testing.T) {
		repository, service, queue, externo := InitServiceTest(t)
		err := os.Setenv("DATABASE_LIMIT_INSERT", "6")
		assert.Nil(t, err)

		arquivosDomain := createArquivosDomainParaProcessamento()
		externo.EXPECT().DownloadArquivosCVMPort(arquivosDomain.Endereco, arquivosDomain.Baixar).Return(
			[]string{arquivosDomain.Endereco},
			nil,
		)

		updateDomain := arquivosDomain
		updateDomain.Processado = true
		updateDomain.Status = constants.PROCESSANDO

		repository.EXPECT().UpdateArquivosRepository(gomock.Any()).DoAndReturn(func(arg domain.ArquivosDomain) {
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

		queue.EXPECT().ProduceLote(gomock.Any()).Return(nil)

		err = service.ProcessarArquivosCVMService(arquivosDomain)
		assert.Nil(t, err)
	})

}
