package controller

import (
	"api-fundos-investimentos/configuration/resterrors"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestCreateTopicController(t *testing.T) {

	service, fundosController := InitServiceTest(t)

	t.Run("when_sending_a_valid_request_returns_success", func(t *testing.T) {

		service.EXPECT().CreateTopicService(gomock.Any()).Return(
			nil,
			nil,
		).AnyTimes()
		fundosController.CreateTopicController()
	})

	t.Run("when_sending_a_invalid_request_returns_error", func(t *testing.T) {

		service.EXPECT().CreateTopicService(gomock.Any()).Return(
			nil,
			resterrors.NewInternalServerError("Erro pra teste"),
		).AnyTimes()

		fundosController.CreateTopicController()
	})

}
