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

func TestCreateManyCdaSelicRepository(t *testing.T) {
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

		selicDomain := []domain.CdaSelicDomain{}
		err := faker.FakeData(&selicDomain)
		assert.Nil(t, err)

		err = repo.CreateManyCdaSelicRepository(selicDomain)
		assert.Nil(t, err)

	})

	mtestDb.Run("return_error_from_database", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(databaseName)
		repo := NewFundosRepository(databaseMock)

		selicDomain := []domain.CdaSelicDomain{}
		err := faker.FakeData(&selicDomain)
		assert.Nil(t, err)

		err = repo.CreateManyCdaSelicRepository(selicDomain)

		assert.NotNil(t, err)
	})
}

func TestCreateManyCdaFundosInvestimentosRepository(t *testing.T) {
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

		fundosInvestimentosDomain := []domain.CdaFundosInvestimentosDomain{}
		err := faker.FakeData(&fundosInvestimentosDomain)
		assert.Nil(t, err)

		err = repo.CreateManyCdaFundosInvestimentosRepository(fundosInvestimentosDomain)
		assert.Nil(t, err)

	})

	mtestDb.Run("return_error_from_database", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(databaseName)
		repo := NewFundosRepository(databaseMock)

		fundosInvestimentosDomain := []domain.CdaFundosInvestimentosDomain{}
		err := faker.FakeData(&fundosInvestimentosDomain)
		assert.Nil(t, err)

		err = repo.CreateManyCdaFundosInvestimentosRepository(fundosInvestimentosDomain)

		assert.NotNil(t, err)
	})
}

func TestCreateManyCdaSwapRepository(t *testing.T) {
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

		swapDomain := []domain.CdaSwapDomain{}
		err := faker.FakeData(&swapDomain)
		assert.Nil(t, err)

		err = repo.CreateManyCdaSwapRepository(swapDomain)
		assert.Nil(t, err)

	})

	mtestDb.Run("return_error_from_database", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(databaseName)
		repo := NewFundosRepository(databaseMock)

		swapDomain := []domain.CdaSwapDomain{}
		err := faker.FakeData(&swapDomain)
		assert.Nil(t, err)

		err = repo.CreateManyCdaSwapRepository(swapDomain)
		assert.NotNil(t, err)
	})
}

func TestCreateManyCdaDemaisAtivosRepository(t *testing.T) {
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

		demaisAtivosDomain := []domain.CdaDemaisAtivosDomain{}
		err := faker.FakeData(&demaisAtivosDomain)
		assert.Nil(t, err)

		err = repo.CreateManyCdaDemaisAtivosRepository(demaisAtivosDomain)
		assert.Nil(t, err)

	})

	mtestDb.Run("return_error_from_database", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(databaseName)
		repo := NewFundosRepository(databaseMock)

		demaisAtivosDomain := []domain.CdaDemaisAtivosDomain{}
		err := faker.FakeData(&demaisAtivosDomain)
		assert.Nil(t, err)

		err = repo.CreateManyCdaDemaisAtivosRepository(demaisAtivosDomain)
		assert.NotNil(t, err)
	})
}

func TestCreateManyCdaDepositoAPrazoOutrosAtivosRepository(t *testing.T) {
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

		depositoAPrazoDomain := []domain.CdaDepositoAPrazoOutrosAtivosDomain{}
		err := faker.FakeData(&depositoAPrazoDomain)
		assert.Nil(t, err)

		err = repo.CreateManyCdaDepositoAPrazoOutrosAtivosRepository(depositoAPrazoDomain)
		assert.Nil(t, err)

	})

	mtestDb.Run("return_error_from_database", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(databaseName)
		repo := NewFundosRepository(databaseMock)

		depositoAPrazoDomain := []domain.CdaDepositoAPrazoOutrosAtivosDomain{}
		err := faker.FakeData(&depositoAPrazoDomain)
		assert.Nil(t, err)

		err = repo.CreateManyCdaDepositoAPrazoOutrosAtivosRepository(depositoAPrazoDomain)
		assert.NotNil(t, err)
	})
}

func TestCreateManyCdaAgroCreditoPrivadoRepository(t *testing.T) {
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

		agroDomain := []domain.CdaAgroCreditoPrivadoDomain{}
		err := faker.FakeData(&agroDomain)
		assert.Nil(t, err)

		err = repo.CreateManyCdaAgroCreditoPrivadoRepository(agroDomain)
		assert.Nil(t, err)

	})

	mtestDb.Run("return_error_from_database", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(databaseName)
		repo := NewFundosRepository(databaseMock)

		agroDomain := []domain.CdaAgroCreditoPrivadoDomain{}
		err := faker.FakeData(&agroDomain)
		assert.Nil(t, err)

		err = repo.CreateManyCdaAgroCreditoPrivadoRepository(agroDomain)
		assert.NotNil(t, err)
	})
}

func TestCreateManyCdaInvestimentosExteriorRepository(t *testing.T) {
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

		exteriorDomain := []domain.CdaInvestimentosExteriorDomain{}
		err := faker.FakeData(&exteriorDomain)
		assert.Nil(t, err)

		err = repo.CreateManyCdaInvestimentosExteriorRepository(exteriorDomain)
		assert.Nil(t, err)

	})

	mtestDb.Run("return_error_from_database", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(databaseName)
		repo := NewFundosRepository(databaseMock)

		exteriorDomain := []domain.CdaInvestimentosExteriorDomain{}
		err := faker.FakeData(&exteriorDomain)
		assert.Nil(t, err)

		err = repo.CreateManyCdaInvestimentosExteriorRepository(exteriorDomain)
		assert.NotNil(t, err)
	})
}

func TestCreateManyCdaDemaisAtivosNaoCodificadosRepository(t *testing.T) {
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

		naoCodificadosDomain := []domain.CdaDemaisAtivosNaoCodificadosDomain{}
		err := faker.FakeData(&naoCodificadosDomain)
		assert.Nil(t, err)

		err = repo.CreateManyCdaDemaisAtivosNaoCodificadosRepository(naoCodificadosDomain)
		assert.Nil(t, err)

	})

	mtestDb.Run("return_error_from_database", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(databaseName)
		repo := NewFundosRepository(databaseMock)

		naoCodificadosDomain := []domain.CdaDemaisAtivosNaoCodificadosDomain{}
		err := faker.FakeData(&naoCodificadosDomain)
		assert.Nil(t, err)

		err = repo.CreateManyCdaDemaisAtivosNaoCodificadosRepository(naoCodificadosDomain)
		assert.NotNil(t, err)
	})
}

func TestCreateManyCdaConfidencialRepository(t *testing.T) {
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

		confidencialDomain := []domain.CdaConfidencialDomain{}
		err := faker.FakeData(&confidencialDomain)
		assert.Nil(t, err)

		err = repo.CreateManyCdaConfidencialRepository(confidencialDomain)
		assert.Nil(t, err)

	})

	mtestDb.Run("return_error_from_database", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(databaseName)
		repo := NewFundosRepository(databaseMock)

		confidencialDomain := []domain.CdaConfidencialDomain{}
		err := faker.FakeData(&confidencialDomain)
		assert.Nil(t, err)

		err = repo.CreateManyCdaConfidencialRepository(confidencialDomain)
		assert.NotNil(t, err)
	})
}

func TestCreateManyCdaFiimRepository(t *testing.T) {
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

		fiimDomain := []domain.CdaFiimDomain{}
		err := faker.FakeData(&fiimDomain)
		assert.Nil(t, err)

		err = repo.CreateManyCdaFiimRepository(fiimDomain)
		assert.Nil(t, err)

	})

	mtestDb.Run("return_error_from_database", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(databaseName)
		repo := NewFundosRepository(databaseMock)

		fiimDomain := []domain.CdaFiimDomain{}
		err := faker.FakeData(&fiimDomain)
		assert.Nil(t, err)

		err = repo.CreateManyCdaFiimRepository(fiimDomain)
		assert.NotNil(t, err)
	})
}

func TestCreateManyCdaFiimConfidencialRepository(t *testing.T) {
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

		fiimConfidencialDomain := []domain.CdaFiimConfidencialDomain{}
		err := faker.FakeData(&fiimConfidencialDomain)
		assert.Nil(t, err)

		err = repo.CreateManyCdaFiimConfidencialRepository(fiimConfidencialDomain)
		assert.Nil(t, err)

	})

	mtestDb.Run("return_error_from_database", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(databaseName)
		repo := NewFundosRepository(databaseMock)

		fiimConfidencialDomain := []domain.CdaFiimConfidencialDomain{}
		err := faker.FakeData(&fiimConfidencialDomain)
		assert.Nil(t, err)

		err = repo.CreateManyCdaFiimConfidencialRepository(fiimConfidencialDomain)
		assert.NotNil(t, err)
	})
}

func TestCreateManyCdaPatrimonioLiquidoRepository(t *testing.T) {
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

		plDomain := []domain.CdaPatrimonioLiquidoDomain{}
		err := faker.FakeData(&plDomain)
		assert.Nil(t, err)

		err = repo.CreateManyCdaPatrimonioLiquidoRepository(plDomain)
		assert.Nil(t, err)

	})

	mtestDb.Run("return_error_from_database", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(databaseName)
		repo := NewFundosRepository(databaseMock)

		plDomain := []domain.CdaPatrimonioLiquidoDomain{}
		err := faker.FakeData(&plDomain)
		assert.Nil(t, err)

		err = repo.CreateManyCdaPatrimonioLiquidoRepository(plDomain)
		assert.NotNil(t, err)
	})
}
