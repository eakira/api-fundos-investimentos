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

func TestArquivosRepository_CreateArquivos(t *testing.T) {
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

		arquivoDomain := domain.ArquivosDomain{}
		err := faker.FakeData(&arquivoDomain)
		assert.Nil(t, err)

		domainCriado, err := repo.CreateArquivosRepository(arquivoDomain)
		assert.Nil(t, err)

		assert.EqualValues(t, arquivoDomain.Endereco, domainCriado.Endereco)
		assert.EqualValues(t, arquivoDomain.TipoArquivo, domainCriado.TipoArquivo)
		assert.EqualValues(t, arquivoDomain.Referencia, domainCriado.Referencia)
		assert.EqualValues(t, arquivoDomain.Status, domainCriado.Status)
		assert.EqualValues(t, arquivoDomain.Baixar, domainCriado.Baixar)
		assert.EqualValues(t, arquivoDomain.Download, domainCriado.Download)
		assert.EqualValues(t, arquivoDomain.Processado, domainCriado.Processado)
		assert.EqualValues(t, arquivoDomain.CreatedAt, domainCriado.CreatedAt)
		assert.EqualValues(t, arquivoDomain.UpdatedAt, domainCriado.UpdatedAt)
	})

	mtestDb.Run("return_error_from_database", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(databaseName)
		repo := NewFundosRepository(databaseMock)

		arquivoDomain := domain.ArquivosDomain{}
		err := faker.FakeData(&arquivoDomain)
		assert.Nil(t, err)

		arquivosDomain, err := repo.CreateArquivosRepository(arquivoDomain)

		assert.NotNil(t, err)
		assert.Nil(t, arquivosDomain)
	})
}
