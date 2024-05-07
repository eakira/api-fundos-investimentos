package controller

import (
	"api-fundos-investimentos/adapter/output/model/request"
	"api-fundos-investimentos/configuration/resterrors"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestProcessarArquivosCVMController(t *testing.T) {

	service, fundosController := InitServiceTest(t)

	t.Run("when_sending_a_valid_request_returns_success", func(t *testing.T) {

		request := request.FundosCvmArquivosQueueRequest{}
		err := faker.FakeData(&request)
		assert.Nil(t, err)

		service.EXPECT().ProcessarArquivosCVMService(gomock.Any()).Return(nil)

		fundosController.ProcessarArquivosCVMController(request)
	})

	t.Run("when_sending_a_invalid_request_returns_error", func(t *testing.T) {

		request := request.FundosCvmArquivosQueueRequest{}
		err := faker.FakeData(&request)
		assert.Nil(t, err)

		service.EXPECT().ProcessarArquivosCVMService(gomock.Any()).Return(
			resterrors.NewBadRequestError("Erro pra teste"),
		)

		fundosController.ProcessarArquivosCVMController(request)
	})

}
