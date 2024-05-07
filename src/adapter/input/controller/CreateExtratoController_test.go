package controller

import (
	"api-fundos-investimentos/adapter/output/model/request"
	"api-fundos-investimentos/configuration/resterrors"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateExtratoController(t *testing.T) {

	service, fundosController := InitServiceTest(t)

	t.Run("when_sending_a_valid_request_returns_success", func(t *testing.T) {

		extratoRequest := []request.ExtratoRequest{}
		err := faker.FakeData(&extratoRequest)
		assert.Nil(t, err)

		service.EXPECT().CreateExtratoService(gomock.Any()).Return(nil)

		fundosController.CreateExtratoController(extratoRequest)
	})

	t.Run("when_sending_a_invalid_request_returns_error", func(t *testing.T) {

		extratoRequest := []request.ExtratoRequest{}
		err := faker.FakeData(&extratoRequest)
		assert.Nil(t, err)

		service.EXPECT().CreateExtratoService(gomock.Any()).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)
		fundosController.CreateExtratoController(extratoRequest)
	})

}