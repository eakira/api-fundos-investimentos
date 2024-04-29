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

	// Enviando um caso de sucesso como fila
	t.Run("processing_a_queued_file_returning_success", func(t *testing.T) {
		repository, service, queue, _ := InitServiceTest(t)
		err := os.Setenv("DATABASE_LIMIT_INSERT", "6")
		assert.Nil(t, err)
		err = os.Setenv("PERSISTENCIA", "fila")
		assert.Nil(t, err)

		arquivosDomain := createArquivosDomainParaProcessamento()

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

	// Erro ao Abrir um arquivo
	t.Run("processing_a_invalid_file_returning_error", func(t *testing.T) {
		repository, service, queue, _ := InitServiceTest(t)
		err := os.Setenv("DATABASE_LIMIT_INSERT", "6")
		assert.Nil(t, err)
		err = os.Setenv("PERSISTENCIA", "fila")
		assert.Nil(t, err)

		arquivosDomain := createArquivosDomainParaProcessamento()
		arquivosDomain.Endereco = "ArquivoQueNãoExiste.csv"

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
		assert.NotNil(t, err)
	})

	// Arquivo vazio
	t.Run("processing_a_blank_file_returning_error", func(t *testing.T) {
		repository, service, queue, _ := InitServiceTest(t)
		err := os.Setenv("DATABASE_LIMIT_INSERT", "6")
		assert.Nil(t, err)
		err = os.Setenv("PERSISTENCIA", "fila")
		assert.Nil(t, err)

		arquivosDomain := createArquivosDomainParaProcessamento()
		arquivosDomain.Endereco = "../../storage/cvm/mock/blank.csv"

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
		assert.NotNil(t, err)
	})

	// Arquivo apenas com cabeçalho
	t.Run("processing_file_with_header_only_returning_error", func(t *testing.T) {
		repository, service, queue, _ := InitServiceTest(t)
		err := os.Setenv("DATABASE_LIMIT_INSERT", "6")
		assert.Nil(t, err)
		err = os.Setenv("PERSISTENCIA", "fila")
		assert.Nil(t, err)

		arquivosDomain := createArquivosDomainParaProcessamento()
		arquivosDomain.Endereco = "../../storage/cvm/mock/apenas_cabecalho.csv"

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
		assert.NotNil(t, err)
	})

}
