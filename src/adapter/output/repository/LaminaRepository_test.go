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

func TestCreateManyLaminaRepository(t *testing.T) {
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

		laminaDomain := []domain.LaminaDomain{}
		err := faker.FakeData(&laminaDomain)
		assert.Nil(t, err)

		err = repo.CreateManyLaminaRepository(laminaDomain)
		assert.Nil(t, err)

	})

	mtestDb.Run("return_error_from_database", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(databaseName)
		repo := NewFundosRepository(databaseMock)

		laminaDomain := []domain.LaminaDomain{}
		err := faker.FakeData(&laminaDomain)
		assert.Nil(t, err)

		err = repo.CreateManyLaminaRepository(laminaDomain)
		assert.NotNil(t, err)
	})
}

func TestCreateManyLaminaCarteiraRepository(t *testing.T) {
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

		laminaCarteiraDomain := []domain.LaminaCarteiraDomain{}
		err := faker.FakeData(&laminaCarteiraDomain)
		assert.Nil(t, err)

		err = repo.CreateManyLaminaCarteiraRepository(laminaCarteiraDomain)
		assert.Nil(t, err)

	})

	mtestDb.Run("return_error_from_database", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(databaseName)
		repo := NewFundosRepository(databaseMock)

		laminaCarteiraDomain := []domain.LaminaCarteiraDomain{}
		err := faker.FakeData(&laminaCarteiraDomain)
		assert.Nil(t, err)

		err = repo.CreateManyLaminaCarteiraRepository(laminaCarteiraDomain)
		assert.NotNil(t, err)
	})
}

func TestCreateManyLaminaRentabilidadeAnoRepository(t *testing.T) {
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

		laminaRentabilidadeAnoDomain := []domain.LaminaRentabilidadeAnoDomain{}
		err := faker.FakeData(&laminaRentabilidadeAnoDomain)
		assert.Nil(t, err)

		err = repo.CreateManyLaminaRentabilidadeAnoRepository(laminaRentabilidadeAnoDomain)
		assert.Nil(t, err)

	})

	mtestDb.Run("return_error_from_database", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(databaseName)
		repo := NewFundosRepository(databaseMock)

		laminaRentabilidadeAnoDomain := []domain.LaminaRentabilidadeAnoDomain{}
		err := faker.FakeData(&laminaRentabilidadeAnoDomain)
		assert.Nil(t, err)

		err = repo.CreateManyLaminaRentabilidadeAnoRepository(laminaRentabilidadeAnoDomain)
		assert.NotNil(t, err)
	})
}

func TestCreateManyLaminaRentabilidadeMesRepository(t *testing.T) {
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

		laminaRentabilidadeMesDomain := []domain.LaminaRentabilidadeMesDomain{}
		err := faker.FakeData(&laminaRentabilidadeMesDomain)
		assert.Nil(t, err)

		err = repo.CreateManyLaminaRentabilidadeMesRepository(laminaRentabilidadeMesDomain)
		assert.Nil(t, err)

	})

	mtestDb.Run("return_error_from_database", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(databaseName)
		repo := NewFundosRepository(databaseMock)

		laminaRentabilidadeMesDomain := []domain.LaminaRentabilidadeMesDomain{}
		err := faker.FakeData(&laminaRentabilidadeMesDomain)
		assert.Nil(t, err)

		err = repo.CreateManyLaminaRentabilidadeMesRepository(laminaRentabilidadeMesDomain)
		assert.NotNil(t, err)
	})
}
