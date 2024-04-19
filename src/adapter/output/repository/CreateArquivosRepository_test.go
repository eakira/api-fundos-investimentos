package repository

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/logger"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestArquivosRepository_CreateArquivos(t *testing.T) {
	databaseName := "user_database_test"

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

		arquivoDomain := domain.ArquivosDomain{
			Endereco:    "FI/CAD/DADOS/cad_fi.csv",
			TipoArquivo: "cadastros",
			Referencia:  "2023",
			Status:      "PROCESSANDO",
			Baixar:      true,
			Download:    false,
			Processado:  false,
			CreatedAt:   time.Now(),
			UpdateAt:    time.Now(),
		}

		_, err := repo.CreateArquivosRepository(arquivoDomain)

		assert.Nil(t, err)
	})

	mtestDb.Run("return_error_from_database", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(databaseName)
		repo := NewFundosRepository(databaseMock)

		arquivoDomain := domain.ArquivosDomain{
			Endereco:    "FI/CAD/DADOS/cad_fi.csv",
			TipoArquivo: "cadastros",
			Referencia:  "2023",
			Status:      "PROCESSANDO",
			Baixar:      true,
			Download:    false,
			Processado:  false,
			CreatedAt:   time.Now(),
			UpdateAt:    time.Now(),
		}

		arquivosDomain, err := repo.CreateArquivosRepository(arquivoDomain)

		assert.NotNil(t, err)
		assert.Nil(t, arquivosDomain)
	})
}
