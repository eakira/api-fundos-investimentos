package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/application/port/input"
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/configuration/resterrors"
	"api-fundos-investimentos/test/mocks"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func InitServiceTest(t *testing.T) (
	repository *mocks.MockFundosPort,
	service input.FundosDomainService,
) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger.InitTestLogger()

	repository = mocks.NewMockFundosPort(ctrl)
	queue := mocks.NewMockFundosQueuePort(ctrl)
	externo := mocks.NewMockFundosExternoPort(ctrl)
	service = NewFundosDomainService(repository, queue, externo)

	return
}

func TestCreateBalanceteService(t *testing.T) {

	repository, service := InitServiceTest(t)

	t.Run("when_sending_a_valid_domain_returns_success", func(t *testing.T) {

		balanceteDomain := []domain.BalanceteDomain{}
		err := faker.FakeData(&balanceteDomain)
		assert.Nil(t, err)

		repository.EXPECT().CreateManyBalecenteRepository(balanceteDomain).Return(
			nil,
		)
		err = service.CreateBalanceteService(balanceteDomain)
		assert.Nil(t, err)
	})

	t.Run("return_error_from_repository", func(t *testing.T) {

		balanceteDomain := []domain.BalanceteDomain{}

		repository.EXPECT().CreateManyBalecenteRepository(balanceteDomain).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		err := service.CreateBalanceteService(balanceteDomain)
		assert.NotNil(t, err)

	})
}
