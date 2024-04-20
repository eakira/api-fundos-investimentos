package repository

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/logger"
	"os"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestCreateFundosRepository(t *testing.T) {
	databaseName := "database_test"

	err := os.Setenv("MONGODB_DATABASE", databaseName)
	if err != nil {
		t.FailNow()
		return
	}
	defer os.Clearenv()

	logger.InitTestLogger()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mtestDb.Run("when_sending_a_valid_domain_returns_success", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "n", Value: 1},
			{Key: "acknowledged", Value: true},
		})

		databaseMock := mt.Client.Database(databaseName)
		repo := NewFundosRepository(databaseMock)

		fundosDomain := domain.FundosDomain{}
		err := faker.FakeData(&fundosDomain)
		assert.Nil(t, err)

		domainCriado, err := repo.CreateFundosRepository(fundosDomain)

		assert.Nil(t, err)
		assert.EqualValues(t, fundosDomain.AdministradorNome, domainCriado.AdministradorNome)
		assert.EqualValues(t, fundosDomain.AuditorNome, domainCriado.AuditorNome)
		assert.EqualValues(t, fundosDomain.CodigoCVM, domainCriado.CodigoCVM)
		assert.EqualValues(t, fundosDomain.Classe, domainCriado.Classe)
		assert.EqualValues(t, fundosDomain.ClasseAnbima, domainCriado.ClasseAnbima)
		assert.EqualValues(t, fundosDomain.AdministradorCNPJ, domainCriado.AdministradorCNPJ)
		assert.EqualValues(t, fundosDomain.AuditorCNPJ, domainCriado.AuditorCNPJ)
		assert.EqualValues(t, fundosDomain.ControladorCNPJ, domainCriado.ControladorCNPJ)
		assert.EqualValues(t, fundosDomain.CustodianteCNPJ, domainCriado.CustodianteCNPJ)
		assert.EqualValues(t, fundosDomain.FundoCNPJ, domainCriado.FundoCNPJ)
		assert.EqualValues(t, fundosDomain.Condominio, domainCriado.Condominio)
		assert.EqualValues(t, fundosDomain.Controlador, domainCriado.Controlador)
		assert.EqualValues(t, fundosDomain.GestorCPFCNPJ, domainCriado.GestorCPFCNPJ)
		assert.EqualValues(t, fundosDomain.Custodiante, domainCriado.Custodiante)
		assert.EqualValues(t, fundosDomain.FundoNome, domainCriado.FundoNome)
		assert.EqualValues(t, fundosDomain.Diretor, domainCriado.Diretor)
		assert.EqualValues(t, fundosDomain.DataCancelamento, domainCriado.DataCancelamento)
		assert.EqualValues(t, fundosDomain.DataConstituicao, domainCriado.DataConstituicao)
		assert.EqualValues(t, fundosDomain.DataFim, domainCriado.DataFim)
		assert.EqualValues(t, fundosDomain.DataInicio, domainCriado.DataInicio)
		assert.EqualValues(t, fundosDomain.DataInicioClasse, domainCriado.DataInicioClasse)
		assert.EqualValues(t, fundosDomain.DataInicioSocial, domainCriado.DataInicioSocial)
		assert.EqualValues(t, fundosDomain.DataInicioSituacao, domainCriado.DataInicioSituacao)
		assert.EqualValues(t, fundosDomain.DataPatrimonioLiquido, domainCriado.DataPatrimonioLiquido)
		assert.EqualValues(t, fundosDomain.DataRegistro, domainCriado.DataRegistro)
		assert.EqualValues(t, fundosDomain.Entidade, domainCriado.Entidade)
		assert.EqualValues(t, fundosDomain.CotasPossui, domainCriado.CotasPossui)
		assert.EqualValues(t, fundosDomain.Exclusivo, domainCriado.Exclusivo)
		assert.EqualValues(t, fundosDomain.GestorNome, domainCriado.GestorNome)
		assert.EqualValues(t, fundosDomain.InformacoesAdicionais, domainCriado.InformacoesAdicionais)
		assert.EqualValues(t, fundosDomain.InfoTaxaPerformance, domainCriado.InfoTaxaPerformance)
		assert.EqualValues(t, fundosDomain.Exterior100, domainCriado.Exterior100)
		assert.EqualValues(t, fundosDomain.PessoaFouPJ, domainCriado.PessoaFouPJ)
		assert.EqualValues(t, fundosDomain.PublicoAlvo, domainCriado.PublicoAlvo)
		assert.EqualValues(t, fundosDomain.IndicadorDesempenho, domainCriado.IndicadorDesempenho)
		assert.EqualValues(t, fundosDomain.Situacao, domainCriado.Situacao)
		assert.EqualValues(t, fundosDomain.TaxaAdm, domainCriado.TaxaAdm)
		assert.EqualValues(t, fundosDomain.TaxaPerformance, domainCriado.TaxaPerformance)
		assert.EqualValues(t, fundosDomain.TipoFundo, domainCriado.TipoFundo)
		assert.EqualValues(t, fundosDomain.TributacaoLongoPrazo, domainCriado.TributacaoLongoPrazo)
		assert.EqualValues(t, fundosDomain.ValorPatriminioLiq, domainCriado.ValorPatriminioLiq)
	})

	mtestDb.Run("return_error_from_database", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(databaseName)
		repo := NewFundosRepository(databaseMock)

		fundosDomain := domain.FundosDomain{}
		err := faker.FakeData(&fundosDomain)
		assert.Nil(t, err)

		arquivosDomain, err := repo.CreateFundosRepository(fundosDomain)

		assert.NotNil(t, err)
		assert.Nil(t, arquivosDomain)
	})
}
