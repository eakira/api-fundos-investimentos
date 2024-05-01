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
		arquivosDomain.Endereco = "../../storage/mock/blank.csv"

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
		arquivosDomain.Endereco = "../../storage/mock/cad_fi.csv"

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
		arquivosDomain.Endereco = "../../storage/mock/cad_fi.csv"

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
		arquivosDomain.Endereco = "../../storage/mock/balancete_fi.csv"
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
		arquivosDomain.Endereco = "../../storage/mock/balancete_fi.csv"
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
		arquivosDomain.Endereco = "../../storage/mock/cda_fi_BLC_1_202402.csv"
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
		arquivosDomain.Endereco = "../../storage/mock/cda_fi_BLC_1_202402.csv"
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
		arquivosDomain.Endereco = "../../storage/mock/cda_fi_BLC_2_202402.csv"
		arquivosDomain.TipoArquivo = "cda"

		repository.EXPECT().UpdateArquivosRepository(gomock.Any()).Return(nil)

		repository.EXPECT().CreateManyCdaFundosInvestimentosRepository(gomock.Any()).Return(nil)

		err = service.ProcessarArquivosCVMService(arquivosDomain)
		assert.Nil(t, err)
	})

	// Processando o arquivo de cotas localmente com erro
	t.Run("processing_a_cotas_locally_returning_error", func(t *testing.T) {
		repository, service, _, _ := InitServiceTest(t)
		err := os.Setenv("DATABASE_LIMIT_INSERT", "100")
		assert.Nil(t, err)
		err = os.Setenv("PERSISTENCIA", "local")
		assert.Nil(t, err)

		arquivosDomain := createArquivosDomainParaProcessamento()
		arquivosDomain.Endereco = "../../storage/mock/cda_fi_BLC_2_202402.csv"
		arquivosDomain.TipoArquivo = "cda"

		repository.EXPECT().UpdateArquivosRepository(gomock.Any()).Return(nil)

		repository.EXPECT().CreateManyCdaFundosInvestimentosRepository(gomock.Any()).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		err = service.ProcessarArquivosCVMService(arquivosDomain)
		assert.NotNil(t, err)
	})

	// Processando o arquivo de swap localmente
	t.Run("processing_a_swap_locally_returning_sucess", func(t *testing.T) {
		repository, service, _, _ := InitServiceTest(t)
		err := os.Setenv("DATABASE_LIMIT_INSERT", "100")
		assert.Nil(t, err)
		err = os.Setenv("PERSISTENCIA", "local")
		assert.Nil(t, err)

		arquivosDomain := createArquivosDomainParaProcessamento()
		arquivosDomain.Endereco = "../../storage/mock/cda_fi_BLC_3_202402.csv"
		arquivosDomain.TipoArquivo = "cda"

		repository.EXPECT().UpdateArquivosRepository(gomock.Any()).Return(nil)

		repository.EXPECT().CreateManyCdaSwapRepository(gomock.Any()).Return(nil)

		err = service.ProcessarArquivosCVMService(arquivosDomain)
		assert.Nil(t, err)
	})

	// Processando o arquivo de swap localmente com erro
	t.Run("processing_a_swap_locally_returning_error", func(t *testing.T) {
		repository, service, _, _ := InitServiceTest(t)
		err := os.Setenv("DATABASE_LIMIT_INSERT", "100")
		assert.Nil(t, err)
		err = os.Setenv("PERSISTENCIA", "local")
		assert.Nil(t, err)

		arquivosDomain := createArquivosDomainParaProcessamento()
		arquivosDomain.Endereco = "../../storage/mock/cda_fi_BLC_3_202402.csv"
		arquivosDomain.TipoArquivo = "cda"

		repository.EXPECT().UpdateArquivosRepository(gomock.Any()).Return(nil)

		repository.EXPECT().CreateManyCdaSwapRepository(gomock.Any()).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		err = service.ProcessarArquivosCVMService(arquivosDomain)
		assert.NotNil(t, err)
	})

	// Processando o arquivo de demais ativos localmente
	t.Run("processing_a_demais_ativos_locally_returning_sucess", func(t *testing.T) {
		repository, service, _, _ := InitServiceTest(t)
		err := os.Setenv("DATABASE_LIMIT_INSERT", "100")
		assert.Nil(t, err)
		err = os.Setenv("PERSISTENCIA", "local")
		assert.Nil(t, err)

		arquivosDomain := createArquivosDomainParaProcessamento()
		arquivosDomain.Endereco = "../../storage/mock/cda_fi_BLC_4_202402.csv"
		arquivosDomain.TipoArquivo = "cda"

		repository.EXPECT().UpdateArquivosRepository(gomock.Any()).Return(nil)

		repository.EXPECT().CreateManyCdaDemaisAtivosRepository(gomock.Any()).Return(nil)

		err = service.ProcessarArquivosCVMService(arquivosDomain)
		assert.Nil(t, err)
	})

	// Processando o arquivo de demais ativos localmente
	t.Run("processing_a_demais_ativos_locally_returning_error", func(t *testing.T) {
		repository, service, _, _ := InitServiceTest(t)
		err := os.Setenv("DATABASE_LIMIT_INSERT", "100")
		assert.Nil(t, err)
		err = os.Setenv("PERSISTENCIA", "local")
		assert.Nil(t, err)

		arquivosDomain := createArquivosDomainParaProcessamento()
		arquivosDomain.Endereco = "../../storage/mock/cda_fi_BLC_4_202402.csv"
		arquivosDomain.TipoArquivo = "cda"

		repository.EXPECT().UpdateArquivosRepository(gomock.Any()).Return(nil)

		repository.EXPECT().CreateManyCdaDemaisAtivosRepository(gomock.Any()).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		err = service.ProcessarArquivosCVMService(arquivosDomain)
		assert.NotNil(t, err)
	})

	// Processando o arquivo de deposito a prazo localmente
	t.Run("processing_a_deposito_a_prazo_locally_returning_sucess", func(t *testing.T) {
		repository, service, _, _ := InitServiceTest(t)
		err := os.Setenv("DATABASE_LIMIT_INSERT", "100")
		assert.Nil(t, err)
		err = os.Setenv("PERSISTENCIA", "local")
		assert.Nil(t, err)

		arquivosDomain := createArquivosDomainParaProcessamento()
		arquivosDomain.Endereco = "../../storage/mock/cda_fi_BLC_5_202402.csv"
		arquivosDomain.TipoArquivo = "cda"

		repository.EXPECT().UpdateArquivosRepository(gomock.Any()).Return(nil)

		repository.EXPECT().CreateManyCdaDepositoAPrazoOutrosAtivosRepository(gomock.Any()).Return(nil)

		err = service.ProcessarArquivosCVMService(arquivosDomain)
		assert.Nil(t, err)
	})

	// Processando o arquivo de demais ativos localmente
	t.Run("processing_a_deposito_a_prazo_locally_returning_error", func(t *testing.T) {
		repository, service, _, _ := InitServiceTest(t)
		err := os.Setenv("DATABASE_LIMIT_INSERT", "100")
		assert.Nil(t, err)
		err = os.Setenv("PERSISTENCIA", "local")
		assert.Nil(t, err)

		arquivosDomain := createArquivosDomainParaProcessamento()
		arquivosDomain.Endereco = "../../storage/mock/cda_fi_BLC_5_202402.csv"
		arquivosDomain.TipoArquivo = "cda"

		repository.EXPECT().UpdateArquivosRepository(gomock.Any()).Return(nil)

		repository.EXPECT().CreateManyCdaDepositoAPrazoOutrosAtivosRepository(gomock.Any()).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		err = service.ProcessarArquivosCVMService(arquivosDomain)
		assert.NotNil(t, err)
	})

	// Processando o arquivo de credito privado localmente
	t.Run("processing_a_credito_privado_locally_returning_sucess", func(t *testing.T) {
		repository, service, _, _ := InitServiceTest(t)
		err := os.Setenv("DATABASE_LIMIT_INSERT", "100")
		assert.Nil(t, err)
		err = os.Setenv("PERSISTENCIA", "local")
		assert.Nil(t, err)

		arquivosDomain := createArquivosDomainParaProcessamento()
		arquivosDomain.Endereco = "../../storage/mock/cda_fi_BLC_6_202402.csv"
		arquivosDomain.TipoArquivo = "cda"

		repository.EXPECT().UpdateArquivosRepository(gomock.Any()).Return(nil)

		repository.EXPECT().CreateManyCdaAgroCreditoPrivadoRepository(gomock.Any()).Return(nil)

		err = service.ProcessarArquivosCVMService(arquivosDomain)
		assert.Nil(t, err)
	})

	// Processando o arquivo de credito privado localmente
	t.Run("processing_a_credito_privado_locally_returning_error", func(t *testing.T) {
		repository, service, _, _ := InitServiceTest(t)
		err := os.Setenv("DATABASE_LIMIT_INSERT", "100")
		assert.Nil(t, err)
		err = os.Setenv("PERSISTENCIA", "local")
		assert.Nil(t, err)

		arquivosDomain := createArquivosDomainParaProcessamento()
		arquivosDomain.Endereco = "../../storage/mock/cda_fi_BLC_6_202402.csv"
		arquivosDomain.TipoArquivo = "cda"

		repository.EXPECT().UpdateArquivosRepository(gomock.Any()).Return(nil)

		repository.EXPECT().CreateManyCdaAgroCreditoPrivadoRepository(gomock.Any()).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		err = service.ProcessarArquivosCVMService(arquivosDomain)
		assert.NotNil(t, err)
	})

	// Processando o arquivo de investimentos no exterior localmente
	t.Run("processing_a_investimentos_no_exterior_locally_returning_sucess", func(t *testing.T) {
		repository, service, _, _ := InitServiceTest(t)
		err := os.Setenv("DATABASE_LIMIT_INSERT", "100")
		assert.Nil(t, err)
		err = os.Setenv("PERSISTENCIA", "local")
		assert.Nil(t, err)

		arquivosDomain := createArquivosDomainParaProcessamento()
		arquivosDomain.Endereco = "../../storage/mock/cda_fi_BLC_7_202402.csv"
		arquivosDomain.TipoArquivo = "cda"

		repository.EXPECT().UpdateArquivosRepository(gomock.Any()).Return(nil)

		repository.EXPECT().CreateManyCdaInvestimentosExteriorRepository(gomock.Any()).Return(nil)

		err = service.ProcessarArquivosCVMService(arquivosDomain)
		assert.Nil(t, err)
	})

	// Processando o arquivo de investimentos no exterior localmente com erro
	t.Run("processing_a_investimentos_no_exterior_returning_error", func(t *testing.T) {
		repository, service, _, _ := InitServiceTest(t)
		err := os.Setenv("DATABASE_LIMIT_INSERT", "100")
		assert.Nil(t, err)
		err = os.Setenv("PERSISTENCIA", "local")
		assert.Nil(t, err)

		arquivosDomain := createArquivosDomainParaProcessamento()
		arquivosDomain.Endereco = "../../storage/mock/cda_fi_BLC_7_202402.csv"
		arquivosDomain.TipoArquivo = "cda"

		repository.EXPECT().UpdateArquivosRepository(gomock.Any()).Return(nil)

		repository.EXPECT().CreateManyCdaInvestimentosExteriorRepository(gomock.Any()).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		err = service.ProcessarArquivosCVMService(arquivosDomain)
		assert.NotNil(t, err)
	})

	// Processando o arquivo de demais ativos não codificados localmente
	t.Run("processing_a_file_demais_ativos_nao_codificados_locally_returning_sucess", func(t *testing.T) {
		repository, service, _, _ := InitServiceTest(t)
		err := os.Setenv("DATABASE_LIMIT_INSERT", "100")
		assert.Nil(t, err)
		err = os.Setenv("PERSISTENCIA", "local")
		assert.Nil(t, err)

		arquivosDomain := createArquivosDomainParaProcessamento()
		arquivosDomain.Endereco = "../../storage/mock/cda_fi_BLC_8_202402.csv"
		arquivosDomain.TipoArquivo = "cda"

		repository.EXPECT().UpdateArquivosRepository(gomock.Any()).Return(nil)

		repository.EXPECT().CreateManyCdaDemaisAtivosNaoCodificadosRepository(gomock.Any()).Return(nil)

		err = service.ProcessarArquivosCVMService(arquivosDomain)
		assert.Nil(t, err)
	})

	// Processando o arquivo de demais ativos não codificados localmente com erro
	t.Run("processing_a_file_demais_ativos_nao_codificados_returning_error", func(t *testing.T) {
		repository, service, _, _ := InitServiceTest(t)
		err := os.Setenv("DATABASE_LIMIT_INSERT", "100")
		assert.Nil(t, err)
		err = os.Setenv("PERSISTENCIA", "local")
		assert.Nil(t, err)

		arquivosDomain := createArquivosDomainParaProcessamento()
		arquivosDomain.Endereco = "../../storage/mock/cda_fi_BLC_8_202402.csv"
		arquivosDomain.TipoArquivo = "cda"

		repository.EXPECT().UpdateArquivosRepository(gomock.Any()).Return(nil)

		repository.EXPECT().CreateManyCdaDemaisAtivosNaoCodificadosRepository(gomock.Any()).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		err = service.ProcessarArquivosCVMService(arquivosDomain)
		assert.NotNil(t, err)
	})

	// Processando o arquivo confidencial localmente
	t.Run("processing_a_file_confidencial_locally_returning_sucess", func(t *testing.T) {
		repository, service, _, _ := InitServiceTest(t)
		err := os.Setenv("DATABASE_LIMIT_INSERT", "100")
		assert.Nil(t, err)
		err = os.Setenv("PERSISTENCIA", "local")
		assert.Nil(t, err)

		arquivosDomain := createArquivosDomainParaProcessamento()
		arquivosDomain.Endereco = "../../storage/mock/cda_fi_CONFID_202402.csv"
		arquivosDomain.TipoArquivo = "cda"

		repository.EXPECT().UpdateArquivosRepository(gomock.Any()).Return(nil)

		repository.EXPECT().CreateManyCdaConfidencialRepository(gomock.Any()).Return(nil)

		err = service.ProcessarArquivosCVMService(arquivosDomain)
		assert.Nil(t, err)
	})

	// Processando o arquivo confidencial localmente com erro
	t.Run("processing_a_file_confidencial_returning_error", func(t *testing.T) {
		repository, service, _, _ := InitServiceTest(t)
		err := os.Setenv("DATABASE_LIMIT_INSERT", "100")
		assert.Nil(t, err)
		err = os.Setenv("PERSISTENCIA", "local")
		assert.Nil(t, err)

		arquivosDomain := createArquivosDomainParaProcessamento()
		arquivosDomain.Endereco = "../../storage/mock/cda_fi_CONFID_202402.csv"
		arquivosDomain.TipoArquivo = "cda"

		repository.EXPECT().UpdateArquivosRepository(gomock.Any()).Return(nil)

		repository.EXPECT().CreateManyCdaConfidencialRepository(gomock.Any()).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		err = service.ProcessarArquivosCVMService(arquivosDomain)
		assert.NotNil(t, err)
	})

	// Processando o arquivo de fiim localmente
	t.Run("processing_a_file_fiim_locally_returning_sucess", func(t *testing.T) {
		repository, service, _, _ := InitServiceTest(t)
		err := os.Setenv("DATABASE_LIMIT_INSERT", "100")
		assert.Nil(t, err)
		err = os.Setenv("PERSISTENCIA", "local")
		assert.Nil(t, err)

		arquivosDomain := createArquivosDomainParaProcessamento()
		arquivosDomain.Endereco = "../../storage/mock/cda_fiim_202402.csv"
		arquivosDomain.TipoArquivo = "cda"

		repository.EXPECT().UpdateArquivosRepository(gomock.Any()).Return(nil)

		repository.EXPECT().CreateManyCdaFiimRepository(gomock.Any()).Return(nil)

		err = service.ProcessarArquivosCVMService(arquivosDomain)
		assert.Nil(t, err)
	})

	// Processando o arquivo fiim localmente com erro
	t.Run("processing_a_file_fiim_returning_error", func(t *testing.T) {
		repository, service, _, _ := InitServiceTest(t)
		err := os.Setenv("DATABASE_LIMIT_INSERT", "100")
		assert.Nil(t, err)
		err = os.Setenv("PERSISTENCIA", "local")
		assert.Nil(t, err)

		arquivosDomain := createArquivosDomainParaProcessamento()
		arquivosDomain.Endereco = "../../storage/mock/cda_fiim_202402.csv"
		arquivosDomain.TipoArquivo = "cda"

		repository.EXPECT().UpdateArquivosRepository(gomock.Any()).Return(nil)

		repository.EXPECT().CreateManyCdaFiimRepository(gomock.Any()).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		err = service.ProcessarArquivosCVMService(arquivosDomain)
		assert.NotNil(t, err)
	})

	// Processando o arquivo de fiim confidencial localmente
	t.Run("processing_a_file_fiim_confidencial_locally_returning_sucess", func(t *testing.T) {
		repository, service, _, _ := InitServiceTest(t)
		err := os.Setenv("DATABASE_LIMIT_INSERT", "100")
		assert.Nil(t, err)
		err = os.Setenv("PERSISTENCIA", "local")
		assert.Nil(t, err)

		arquivosDomain := createArquivosDomainParaProcessamento()
		arquivosDomain.Endereco = "../../storage/mock/cda_fiim_CONFID_202402.csv"
		arquivosDomain.TipoArquivo = "cda"

		repository.EXPECT().UpdateArquivosRepository(gomock.Any()).Return(nil)

		repository.EXPECT().CreateManyCdaFiimConfidencialRepository(gomock.Any()).Return(nil)

		err = service.ProcessarArquivosCVMService(arquivosDomain)
		assert.Nil(t, err)
	})

	// Processando o arquivo fiim confidencial localmente com erro
	t.Run("processing_a_file_fiim_confidencial_returning_error", func(t *testing.T) {
		repository, service, _, _ := InitServiceTest(t)
		err := os.Setenv("DATABASE_LIMIT_INSERT", "100")
		assert.Nil(t, err)
		err = os.Setenv("PERSISTENCIA", "local")
		assert.Nil(t, err)

		arquivosDomain := createArquivosDomainParaProcessamento()
		arquivosDomain.Endereco = "../../storage/mock/cda_fiim_CONFID_202402.csv"
		arquivosDomain.TipoArquivo = "cda"

		repository.EXPECT().UpdateArquivosRepository(gomock.Any()).Return(nil)

		repository.EXPECT().CreateManyCdaFiimConfidencialRepository(gomock.Any()).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		err = service.ProcessarArquivosCVMService(arquivosDomain)
		assert.NotNil(t, err)
	})

	// Processando o arquivo de patrimonio liquido localmente
	t.Run("processing_a_file_patrimonio_liquido_locally_returning_sucess", func(t *testing.T) {
		repository, service, _, _ := InitServiceTest(t)
		err := os.Setenv("DATABASE_LIMIT_INSERT", "100")
		assert.Nil(t, err)
		err = os.Setenv("PERSISTENCIA", "local")
		assert.Nil(t, err)

		arquivosDomain := createArquivosDomainParaProcessamento()
		arquivosDomain.Endereco = "../../storage/mock/cda_fi_PL_202402.csv"
		arquivosDomain.TipoArquivo = "cda"

		repository.EXPECT().UpdateArquivosRepository(gomock.Any()).Return(nil)

		repository.EXPECT().CreateManyCdaPatrimonioLiquidoRepository(gomock.Any()).Return(nil)

		err = service.ProcessarArquivosCVMService(arquivosDomain)
		assert.Nil(t, err)
	})

	// Processando o arquivo de patrimonio liquido localmente com erro
	t.Run("processing_a_file_patrimonio_liquido_returning_error", func(t *testing.T) {
		repository, service, _, _ := InitServiceTest(t)
		err := os.Setenv("DATABASE_LIMIT_INSERT", "100")
		assert.Nil(t, err)
		err = os.Setenv("PERSISTENCIA", "local")
		assert.Nil(t, err)

		arquivosDomain := createArquivosDomainParaProcessamento()
		arquivosDomain.Endereco = "../../storage/mock/cda_fi_PL_202402.csv"
		arquivosDomain.TipoArquivo = "cda"

		repository.EXPECT().UpdateArquivosRepository(gomock.Any()).Return(nil)

		repository.EXPECT().CreateManyCdaPatrimonioLiquidoRepository(gomock.Any()).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		err = service.ProcessarArquivosCVMService(arquivosDomain)
		assert.NotNil(t, err)
	})

	// Processando o arquivo de informacao complementar localmente
	t.Run("processing_a_file_informacao_complementar_locally_returning_sucess", func(t *testing.T) {
		repository, service, _, _ := InitServiceTest(t)
		err := os.Setenv("DATABASE_LIMIT_INSERT", "100")
		assert.Nil(t, err)
		err = os.Setenv("PERSISTENCIA", "local")
		assert.Nil(t, err)

		arquivosDomain := createArquivosDomainParaProcessamento()
		arquivosDomain.Endereco = "../../storage/mock/compl_fi_201809.csv"
		arquivosDomain.TipoArquivo = "informacoes-complementares"

		repository.EXPECT().UpdateArquivosRepository(gomock.Any()).Return(nil)

		repository.EXPECT().CreateManyInformacaoComplementarFundoRepository(gomock.Any()).Return(nil)

		err = service.ProcessarArquivosCVMService(arquivosDomain)
		assert.Nil(t, err)
	})

	// Processando o arquivo de informacao complementar localmente com erro
	t.Run("processing_a_file_informacao_complementar_returning_error", func(t *testing.T) {
		repository, service, _, _ := InitServiceTest(t)
		err := os.Setenv("DATABASE_LIMIT_INSERT", "100")
		assert.Nil(t, err)
		err = os.Setenv("PERSISTENCIA", "local")
		assert.Nil(t, err)

		arquivosDomain := createArquivosDomainParaProcessamento()
		arquivosDomain.Endereco = "../../storage/mock/compl_fi_201809.csv"
		arquivosDomain.TipoArquivo = "informacoes-complementares"

		repository.EXPECT().UpdateArquivosRepository(gomock.Any()).Return(nil)

		repository.EXPECT().CreateManyInformacaoComplementarFundoRepository(gomock.Any()).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		err = service.ProcessarArquivosCVMService(arquivosDomain)
		assert.NotNil(t, err)
	})

	// Processando o arquivo de informacao complementar divulgação localmente
	t.Run("processing_a_file_informacao_complementar_divulgacao_returning_sucess", func(t *testing.T) {
		repository, service, _, _ := InitServiceTest(t)
		err := os.Setenv("DATABASE_LIMIT_INSERT", "100")
		assert.Nil(t, err)
		err = os.Setenv("PERSISTENCIA", "local")
		assert.Nil(t, err)

		arquivosDomain := createArquivosDomainParaProcessamento()
		arquivosDomain.Endereco = "../../storage/mock/compl_fi_inf_201901.csv"
		arquivosDomain.TipoArquivo = "informacoes-complementares"

		repository.EXPECT().UpdateArquivosRepository(gomock.Any()).Return(nil)

		repository.EXPECT().CreateManyInformacaoComplementarDivulgacaoRepository(gomock.Any()).Return(nil)

		err = service.ProcessarArquivosCVMService(arquivosDomain)
		assert.Nil(t, err)
	})

	// Processando o arquivo de informacao complementar divulgação localmente com erro
	t.Run("processing_a_file_informacao_complementar_divulgacao_returning_error", func(t *testing.T) {
		repository, service, _, _ := InitServiceTest(t)
		err := os.Setenv("DATABASE_LIMIT_INSERT", "100")
		assert.Nil(t, err)
		err = os.Setenv("PERSISTENCIA", "local")
		assert.Nil(t, err)

		arquivosDomain := createArquivosDomainParaProcessamento()
		arquivosDomain.Endereco = "../../storage/mock/compl_fi_inf_201901.csv"
		arquivosDomain.TipoArquivo = "informacoes-complementares"

		repository.EXPECT().UpdateArquivosRepository(gomock.Any()).Return(nil)

		repository.EXPECT().CreateManyInformacaoComplementarDivulgacaoRepository(gomock.Any()).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		err = service.ProcessarArquivosCVMService(arquivosDomain)
		assert.NotNil(t, err)
	})

	// Processando o arquivo de informacao complementar de cotista localmente
	t.Run("processing_a_file_informacao_complementar_cotista_returning_sucess", func(t *testing.T) {
		repository, service, _, _ := InitServiceTest(t)
		err := os.Setenv("DATABASE_LIMIT_INSERT", "100")
		assert.Nil(t, err)
		err = os.Setenv("PERSISTENCIA", "local")
		assert.Nil(t, err)

		arquivosDomain := createArquivosDomainParaProcessamento()
		arquivosDomain.Endereco = "../../storage/mock/compl_fi_inf_cotst_201901.csv"
		arquivosDomain.TipoArquivo = "informacoes-complementares"

		repository.EXPECT().UpdateArquivosRepository(gomock.Any()).Return(nil)

		repository.EXPECT().CreateManyInformacaoComplementarCotistaRepository(gomock.Any()).Return(nil)

		err = service.ProcessarArquivosCVMService(arquivosDomain)
		assert.Nil(t, err)
	})

	// Processando o arquivo de informacao complementar de cotista localmente com erro
	t.Run("processing_a_file_informacao_complementar_cotista_returning_error", func(t *testing.T) {
		repository, service, _, _ := InitServiceTest(t)
		err := os.Setenv("DATABASE_LIMIT_INSERT", "100")
		assert.Nil(t, err)
		err = os.Setenv("PERSISTENCIA", "local")
		assert.Nil(t, err)

		arquivosDomain := createArquivosDomainParaProcessamento()
		arquivosDomain.Endereco = "../../storage/mock/compl_fi_inf_cotst_201901.csv"
		arquivosDomain.TipoArquivo = "informacoes-complementares"

		repository.EXPECT().UpdateArquivosRepository(gomock.Any()).Return(nil)

		repository.EXPECT().CreateManyInformacaoComplementarCotistaRepository(gomock.Any()).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		err = service.ProcessarArquivosCVMService(arquivosDomain)
		assert.NotNil(t, err)
	})

	// Processando o arquivo de informacao complementar de servico prestado localmente
	t.Run("processing_a_file_informacao_complementar_servico_prestado_returning_sucess", func(t *testing.T) {
		repository, service, _, _ := InitServiceTest(t)
		err := os.Setenv("DATABASE_LIMIT_INSERT", "100")
		assert.Nil(t, err)
		err = os.Setenv("PERSISTENCIA", "local")
		assert.Nil(t, err)

		arquivosDomain := createArquivosDomainParaProcessamento()
		arquivosDomain.Endereco = "../../storage/mock/compl_fi_prest_201901.csv"
		arquivosDomain.TipoArquivo = "informacoes-complementares"

		repository.EXPECT().UpdateArquivosRepository(gomock.Any()).Return(nil)

		repository.EXPECT().CreateManyInformacaoComplementarServicoPrestadoRepository(gomock.Any()).Return(nil)

		err = service.ProcessarArquivosCVMService(arquivosDomain)
		assert.Nil(t, err)
	})

	// Processando o arquivo de informacao complementar de servico prestado localmente com erro
	t.Run("processing_a_file_informacao_complementar_servico_prestado_returning_error", func(t *testing.T) {
		repository, service, _, _ := InitServiceTest(t)
		err := os.Setenv("DATABASE_LIMIT_INSERT", "100")
		assert.Nil(t, err)
		err = os.Setenv("PERSISTENCIA", "local")
		assert.Nil(t, err)

		arquivosDomain := createArquivosDomainParaProcessamento()
		arquivosDomain.Endereco = "../../storage/mock/compl_fi_prest_201901.csv"
		arquivosDomain.TipoArquivo = "informacoes-complementares"

		repository.EXPECT().UpdateArquivosRepository(gomock.Any()).Return(nil)

		repository.EXPECT().CreateManyInformacaoComplementarServicoPrestadoRepository(gomock.Any()).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		err = service.ProcessarArquivosCVMService(arquivosDomain)
		assert.NotNil(t, err)
	})

	// Processando o arquivo de lamina localmente
	t.Run("processing_a_file_lamina_returning_sucess", func(t *testing.T) {
		repository, service, _, _ := InitServiceTest(t)
		err := os.Setenv("DATABASE_LIMIT_INSERT", "100")
		assert.Nil(t, err)
		err = os.Setenv("PERSISTENCIA", "local")
		assert.Nil(t, err)

		arquivosDomain := createArquivosDomainParaProcessamento()
		arquivosDomain.Endereco = "../../storage/mock/lamina_fi_202402.csv"
		arquivosDomain.TipoArquivo = "lamina"

		repository.EXPECT().UpdateArquivosRepository(gomock.Any()).Return(nil)

		repository.EXPECT().CreateManyLaminaRepository(gomock.Any()).Return(nil)

		err = service.ProcessarArquivosCVMService(arquivosDomain)
		assert.Nil(t, err)
	})

	// Processando o arquivo de lamina localmente com erro
	t.Run("processing_a_file_lamina_returning_error", func(t *testing.T) {
		repository, service, _, _ := InitServiceTest(t)
		err := os.Setenv("DATABASE_LIMIT_INSERT", "100")
		assert.Nil(t, err)
		err = os.Setenv("PERSISTENCIA", "local")
		assert.Nil(t, err)

		arquivosDomain := createArquivosDomainParaProcessamento()
		arquivosDomain.Endereco = "../../storage/mock/lamina_fi_202402.csv"
		arquivosDomain.TipoArquivo = "lamina"

		repository.EXPECT().UpdateArquivosRepository(gomock.Any()).Return(nil)

		repository.EXPECT().CreateManyLaminaRepository(gomock.Any()).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		err = service.ProcessarArquivosCVMService(arquivosDomain)
		assert.NotNil(t, err)
	})

	// Processando o arquivo de lamina da carteira localmente
	t.Run("processing_a_file_lamina_da_carteira_returning_sucess", func(t *testing.T) {
		repository, service, _, _ := InitServiceTest(t)
		err := os.Setenv("DATABASE_LIMIT_INSERT", "100")
		assert.Nil(t, err)
		err = os.Setenv("PERSISTENCIA", "local")
		assert.Nil(t, err)

		arquivosDomain := createArquivosDomainParaProcessamento()
		arquivosDomain.Endereco = "../../storage/mock/lamina_fi_carteira_202110.csv"
		arquivosDomain.TipoArquivo = "lamina"

		repository.EXPECT().UpdateArquivosRepository(gomock.Any()).Return(nil)

		repository.EXPECT().CreateManyLaminaCarteiraRepository(gomock.Any()).Return(nil)

		err = service.ProcessarArquivosCVMService(arquivosDomain)
		assert.Nil(t, err)
	})

	// Processando o arquivo de lamina da carteira localmente com erro
	t.Run("processing_a_file_lamina_da_carteira_returning_error", func(t *testing.T) {
		repository, service, _, _ := InitServiceTest(t)
		err := os.Setenv("DATABASE_LIMIT_INSERT", "100")
		assert.Nil(t, err)
		err = os.Setenv("PERSISTENCIA", "local")
		assert.Nil(t, err)

		arquivosDomain := createArquivosDomainParaProcessamento()
		arquivosDomain.Endereco = "../../storage/mock/lamina_fi_carteira_202110.csv"
		arquivosDomain.TipoArquivo = "lamina"

		repository.EXPECT().UpdateArquivosRepository(gomock.Any()).Return(nil)

		repository.EXPECT().CreateManyLaminaCarteiraRepository(gomock.Any()).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		err = service.ProcessarArquivosCVMService(arquivosDomain)
		assert.NotNil(t, err)
	})

	// Processando o arquivo de lamina rentabilidade do mes localmente
	t.Run("processing_a_file_lamina_rentabilidade_mes_returning_sucess", func(t *testing.T) {
		repository, service, _, _ := InitServiceTest(t)
		err := os.Setenv("DATABASE_LIMIT_INSERT", "100")
		assert.Nil(t, err)
		err = os.Setenv("PERSISTENCIA", "local")
		assert.Nil(t, err)

		arquivosDomain := createArquivosDomainParaProcessamento()
		arquivosDomain.Endereco = "../../storage/mock/lamina_fi_rentab_mes_202110.csv"
		arquivosDomain.TipoArquivo = "lamina"

		repository.EXPECT().UpdateArquivosRepository(gomock.Any()).Return(nil)

		repository.EXPECT().CreateManyLaminaRentabilidadeMesRepository(gomock.Any()).Return(nil)

		err = service.ProcessarArquivosCVMService(arquivosDomain)
		assert.Nil(t, err)
	})

	// Processando o arquivo de lamina rentabilidade do mes localmente com erro
	t.Run("processing_a_file_lamina_rentabilidade_mes_returning_error", func(t *testing.T) {
		repository, service, _, _ := InitServiceTest(t)
		err := os.Setenv("DATABASE_LIMIT_INSERT", "100")
		assert.Nil(t, err)
		err = os.Setenv("PERSISTENCIA", "local")
		assert.Nil(t, err)

		arquivosDomain := createArquivosDomainParaProcessamento()
		arquivosDomain.Endereco = "../../storage/mock/lamina_fi_rentab_mes_202110.csv"
		arquivosDomain.TipoArquivo = "lamina"

		repository.EXPECT().UpdateArquivosRepository(gomock.Any()).Return(nil)

		repository.EXPECT().CreateManyLaminaRentabilidadeMesRepository(gomock.Any()).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		err = service.ProcessarArquivosCVMService(arquivosDomain)
		assert.NotNil(t, err)
	})

	// Processando o arquivo de lamina rentabilidade do ano localmente
	t.Run("processing_a_file_lamina_rentabilidade_ano_returning_sucess", func(t *testing.T) {
		repository, service, _, _ := InitServiceTest(t)
		err := os.Setenv("DATABASE_LIMIT_INSERT", "100")
		assert.Nil(t, err)
		err = os.Setenv("PERSISTENCIA", "local")
		assert.Nil(t, err)

		arquivosDomain := createArquivosDomainParaProcessamento()
		arquivosDomain.Endereco = "../../storage/mock/lamina_fi_rentab_ano_202110.csv"
		arquivosDomain.TipoArquivo = "lamina"

		repository.EXPECT().UpdateArquivosRepository(gomock.Any()).Return(nil)

		repository.EXPECT().CreateManyLaminaRentabilidadeAnoRepository(gomock.Any()).Return(nil)

		err = service.ProcessarArquivosCVMService(arquivosDomain)
		assert.Nil(t, err)
	})

	// Processando o arquivo de lamina rentabilidade do ano localmente com erro
	t.Run("processing_a_file_lamina_rentabilidade_ano_returning_error", func(t *testing.T) {
		repository, service, _, _ := InitServiceTest(t)
		err := os.Setenv("DATABASE_LIMIT_INSERT", "100")
		assert.Nil(t, err)
		err = os.Setenv("PERSISTENCIA", "local")
		assert.Nil(t, err)

		arquivosDomain := createArquivosDomainParaProcessamento()
		arquivosDomain.Endereco = "../../storage/mock/lamina_fi_rentab_ano_202110.csv"
		arquivosDomain.TipoArquivo = "lamina"

		repository.EXPECT().UpdateArquivosRepository(gomock.Any()).Return(nil)

		repository.EXPECT().CreateManyLaminaRentabilidadeAnoRepository(gomock.Any()).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		err = service.ProcessarArquivosCVMService(arquivosDomain)
		assert.NotNil(t, err)
	})

}
