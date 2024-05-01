package services

import (
	"api-fundos-investimentos/application/constants"
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/resterrors"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func createArquivosDomainParaProcessamento() domain.ArquivosDomain {
	return domain.ArquivosDomain{
		Endereco:    "../../storage/mock/cad_fi.csv",
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

	// Erro ao salvar o arquivo
	t.Run("processing_a_invalid_domain_returning_error", func(t *testing.T) {
		repository, service, queue, _ := InitServiceTest(t)
		err := os.Setenv("DATABASE_LIMIT_INSERT", "6")
		assert.Nil(t, err)
		err = os.Setenv("PERSISTENCIA", "fila")
		assert.Nil(t, err)

		arquivosDomain := createArquivosDomainParaProcessamento()

		updateDomain := arquivosDomain
		updateDomain.Processado = true
		updateDomain.Status = constants.PROCESSANDO

		repository.EXPECT().UpdateArquivosRepository(gomock.Any()).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		queue.EXPECT().ProduceLote(gomock.Any()).Return(nil)

		err = service.ProcessarArquivosCVMService(arquivosDomain)
		assert.NotNil(t, err)
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
		arquivosDomain.Endereco = "../../storage//mock/blank.csv"

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

	// Processando o arquivo de cadastro localmente
	t.Run("processing_a_funds_locally_returning_sucess", func(t *testing.T) {
		repository, service, _, _ := InitServiceTest(t)
		err := os.Setenv("DATABASE_LIMIT_INSERT", "6")
		assert.Nil(t, err)
		err = os.Setenv("PERSISTENCIA", "local")
		assert.Nil(t, err)

		arquivosDomain := createArquivosDomainParaProcessamento()
		arquivosDomain.Endereco = "../../storage//mock/cad_fi.csv"

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

		repository.EXPECT().CreateManyFundosRepository(gomock.Any()).Return(
			nil,
		)

		repository.EXPECT().CreateManyFundosRepository(gomock.Any()).Return(
			nil,
		)

		err = service.ProcessarArquivosCVMService(arquivosDomain)
		assert.Nil(t, err)
	})

	// Processando o arquivo de cadastro localmente
	t.Run("processing_a_funds_locally_returning_error", func(t *testing.T) {
		repository, service, _, _ := InitServiceTest(t)
		err := os.Setenv("DATABASE_LIMIT_INSERT", "100")
		assert.Nil(t, err)
		err = os.Setenv("PERSISTENCIA", "local")
		assert.Nil(t, err)

		arquivosDomain := createArquivosDomainParaProcessamento()
		arquivosDomain.Endereco = "../../storage//mock/cad_fi.csv"

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

		repository.EXPECT().CreateManyFundosRepository(gomock.Any()).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		err = service.ProcessarArquivosCVMService(arquivosDomain)
		assert.NotNil(t, err)
	})

	// Processando o arquivo de Balancente localmente
	t.Run("processing_a_balancente_locally_returning_sucess", func(t *testing.T) {
		repository, service, _, _ := InitServiceTest(t)
		err := os.Setenv("DATABASE_LIMIT_INSERT", "100")
		assert.Nil(t, err)
		err = os.Setenv("PERSISTENCIA", "local")
		assert.Nil(t, err)

		arquivosDomain := createArquivosDomainParaProcessamento()
		arquivosDomain.Endereco = "../../storage//mock/balancete_fi.csv"
		arquivosDomain.TipoArquivo = "balancete"

		repository.EXPECT().UpdateArquivosRepository(gomock.Any()).Return(nil)

		repository.EXPECT().CreateManyBalecenteRepository(gomock.Any()).Return(nil)

		err = service.ProcessarArquivosCVMService(arquivosDomain)
		assert.Nil(t, err)
	})

	// Processando o arquivo de Balancente localmente
	t.Run("processing_a_balancente_locally_returning_error", func(t *testing.T) {
		repository, service, _, _ := InitServiceTest(t)
		err := os.Setenv("DATABASE_LIMIT_INSERT", "100")
		assert.Nil(t, err)
		err = os.Setenv("PERSISTENCIA", "local")
		assert.Nil(t, err)

		arquivosDomain := createArquivosDomainParaProcessamento()
		arquivosDomain.Endereco = "../../storage//mock/balancete_fi.csv"
		arquivosDomain.TipoArquivo = "balancete"

		repository.EXPECT().UpdateArquivosRepository(gomock.Any()).Return(nil)

		repository.EXPECT().CreateManyBalecenteRepository(gomock.Any()).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		err = service.ProcessarArquivosCVMService(arquivosDomain)
		assert.NotNil(t, err)
	})

	// Processando o arquivo de selic localmente
	t.Run("processing_a_selic_locally_returning_sucess", func(t *testing.T) {
		repository, service, _, _ := InitServiceTest(t)
		err := os.Setenv("DATABASE_LIMIT_INSERT", "100")
		assert.Nil(t, err)
		err = os.Setenv("PERSISTENCIA", "local")
		assert.Nil(t, err)

		arquivosDomain := createArquivosDomainParaProcessamento()
		arquivosDomain.Endereco = "../../storage//mock/cda_fi_BLC_1_202402.csv"
		arquivosDomain.TipoArquivo = "cda"

		repository.EXPECT().UpdateArquivosRepository(gomock.Any()).Return(nil)

		repository.EXPECT().CreateManyCdaSelicRepository(gomock.Any()).Return(nil)

		err = service.ProcessarArquivosCVMService(arquivosDomain)
		assert.Nil(t, err)
	})

	// Processando o arquivo de selic localmente
	t.Run("processing_a_selic_locally_returning_error", func(t *testing.T) {
		repository, service, _, _ := InitServiceTest(t)
		err := os.Setenv("DATABASE_LIMIT_INSERT", "100")
		assert.Nil(t, err)
		err = os.Setenv("PERSISTENCIA", "local")
		assert.Nil(t, err)

		arquivosDomain := createArquivosDomainParaProcessamento()
		arquivosDomain.Endereco = "../../storage//mock/cda_fi_BLC_1_202402.csv"
		arquivosDomain.TipoArquivo = "cda"

		repository.EXPECT().UpdateArquivosRepository(gomock.Any()).Return(nil)

		repository.EXPECT().CreateManyCdaSelicRepository(gomock.Any()).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		err = service.ProcessarArquivosCVMService(arquivosDomain)
		assert.NotNil(t, err)
	})

	// Processando o arquivo de cotas localmente
	t.Run("processing_a_cotas_locally_returning_sucess", func(t *testing.T) {
		repository, service, _, _ := InitServiceTest(t)
		err := os.Setenv("DATABASE_LIMIT_INSERT", "100")
		assert.Nil(t, err)
		err = os.Setenv("PERSISTENCIA", "local")
		assert.Nil(t, err)

		arquivosDomain := createArquivosDomainParaProcessamento()
		arquivosDomain.Endereco = "../../storage//mock/cda_fi_BLC_2_202402.csv"
		arquivosDomain.TipoArquivo = "cda"

		repository.EXPECT().UpdateArquivosRepository(gomock.Any()).Return(nil)

		repository.EXPECT().CreateManyCdaFundosInvestimentosRepository(gomock.Any()).Return(nil)

		err = service.ProcessarArquivosCVMService(arquivosDomain)
		assert.Nil(t, err)
	})

	// Processando o arquivo de cotas localmente
	t.Run("processing_a_cotas_locally_returning_error", func(t *testing.T) {
		repository, service, _, _ := InitServiceTest(t)
		err := os.Setenv("DATABASE_LIMIT_INSERT", "100")
		assert.Nil(t, err)
		err = os.Setenv("PERSISTENCIA", "local")
		assert.Nil(t, err)

		arquivosDomain := createArquivosDomainParaProcessamento()
		arquivosDomain.Endereco = "../../storage//mock/cda_fi_BLC_2_202402.csv"
		arquivosDomain.TipoArquivo = "cda"

		repository.EXPECT().UpdateArquivosRepository(gomock.Any()).Return(nil)

		repository.EXPECT().CreateManyCdaFundosInvestimentosRepository(gomock.Any()).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		err = service.ProcessarArquivosCVMService(arquivosDomain)
		assert.NotNil(t, err)
	})
}
