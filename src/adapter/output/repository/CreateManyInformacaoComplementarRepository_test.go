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

func TestCreateManyInformacaoComplementarFundoRepository(t *testing.T) {
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

		informacoesFundoDomain := []domain.InformacoesFundoDomain{}
		err := faker.FakeData(&informacoesFundoDomain)
		assert.Nil(t, err)

		err = repo.CreateManyInformacaoComplementarFundoRepository(informacoesFundoDomain)
		assert.Nil(t, err)

	})

	mtestDb.Run("return_error_from_database", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(databaseName)
		repo := NewFundosRepository(databaseMock)

		informacoesFundoDomain := []domain.InformacoesFundoDomain{}
		err := faker.FakeData(&informacoesFundoDomain)
		assert.Nil(t, err)

		err = repo.CreateManyInformacaoComplementarFundoRepository(informacoesFundoDomain)
		assert.NotNil(t, err)
	})
}

func TestCreateManyInformacaoComplementarDivulgacaoRepository(t *testing.T) {
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

		informacoesDivulgacaoDomain := []domain.InformacoesDivulgacaoDomain{}
		err := faker.FakeData(&informacoesDivulgacaoDomain)
		assert.Nil(t, err)

		err = repo.CreateManyInformacaoComplementarDivulgacaoRepository(informacoesDivulgacaoDomain)
		assert.Nil(t, err)

	})

	mtestDb.Run("return_error_from_database", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(databaseName)
		repo := NewFundosRepository(databaseMock)

		informacoesDivulgacaoDomain := []domain.InformacoesDivulgacaoDomain{}
		err := faker.FakeData(&informacoesDivulgacaoDomain)
		assert.Nil(t, err)

		err = repo.CreateManyInformacaoComplementarDivulgacaoRepository(informacoesDivulgacaoDomain)
		assert.NotNil(t, err)
	})
}

func TestCreateManyInformacaoComplementarCotistaRepository(t *testing.T) {
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

		informacoesCotistaDomain := []domain.InformacoesCotistaDomain{}
		err := faker.FakeData(&informacoesCotistaDomain)
		assert.Nil(t, err)

		err = repo.CreateManyInformacaoComplementarCotistaRepository(informacoesCotistaDomain)
		assert.Nil(t, err)

	})

	mtestDb.Run("return_error_from_database", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(databaseName)
		repo := NewFundosRepository(databaseMock)

		informacoesCotistaDomain := []domain.InformacoesCotistaDomain{}
		err := faker.FakeData(&informacoesCotistaDomain)
		assert.Nil(t, err)

		err = repo.CreateManyInformacaoComplementarCotistaRepository(informacoesCotistaDomain)
		assert.NotNil(t, err)
	})
}

func TestCreateManyInformacaoComplementarServicoPrestadoRepository(t *testing.T) {
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

		servicoPrestadoDomain := []domain.ServicoPrestadoDomain{}
		err := faker.FakeData(&servicoPrestadoDomain)
		assert.Nil(t, err)

		err = repo.CreateManyInformacaoComplementarServicoPrestadoRepository(servicoPrestadoDomain)
		assert.Nil(t, err)

	})

	mtestDb.Run("return_error_from_database", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(databaseName)
		repo := NewFundosRepository(databaseMock)

		servicoPrestadoDomain := []domain.ServicoPrestadoDomain{}
		err := faker.FakeData(&servicoPrestadoDomain)
		assert.Nil(t, err)

		err = repo.CreateManyInformacaoComplementarServicoPrestadoRepository(servicoPrestadoDomain)
		assert.NotNil(t, err)
	})
}
