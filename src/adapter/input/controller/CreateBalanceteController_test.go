package controller

import (
	"api-fundos-investimentos/adapter/output/model/request"
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/configuration/resterrors"
	"api-fundos-investimentos/test/mocks"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func InitServiceTest(t *testing.T) (
	service *mocks.MockFundosDomainService,
	fundosController FundosControllerInterface,
) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger.InitTestLogger()

	service = mocks.NewMockFundosDomainService(ctrl)
	fundosController = NewFundosControllerInterface(service)
	return
}

func TestCreateBalanceteService(t *testing.T) {

	service, fundosController := InitServiceTest(t)

	t.Run("when_sending_a_valid_request_returns_success", func(t *testing.T) {

		balanceteRequest := []request.BalanceteRequest{}
		err := faker.FakeData(&balanceteRequest)
		assert.Nil(t, err)

		service.EXPECT().CreateBalanceteService(gomock.Any()).Return(nil)

		fundosController.CreateBalanceteController(balanceteRequest)
	})

	t.Run("when_sending_a_invalid_request_returns_error", func(t *testing.T) {

		balanceteRequest := []request.BalanceteRequest{}
		err := faker.FakeData(&balanceteRequest)
		assert.Nil(t, err)

		service.EXPECT().CreateBalanceteService(gomock.Any()).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)
		fundosController.CreateBalanceteController(balanceteRequest)
	})

}
