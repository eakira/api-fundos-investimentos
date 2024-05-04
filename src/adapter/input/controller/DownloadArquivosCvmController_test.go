package controller

import (
	"api-fundos-investimentos/adapter/output/model/request"
	"api-fundos-investimentos/configuration/resterrors"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestDownloadArquivosCVMController(t *testing.T) {

	service, fundosController := InitServiceTest(t)

	t.Run("when_sending_a_valid_request_returns_success", func(t *testing.T) {

		DownloadArquivosCVMRequest := request.FundosCvmArquivosQueueRequest{}
		err := faker.FakeData(&DownloadArquivosCVMRequest)
		assert.Nil(t, err)

		service.EXPECT().DownloadArquivosCVMService(gomock.Any()).Return(nil)

		fundosController.DownloadArquivosCVMController(DownloadArquivosCVMRequest)
	})

	t.Run("when_sending_a_invalid_request_returns_error", func(t *testing.T) {

		DownloadArquivosCVMRequest := request.FundosCvmArquivosQueueRequest{}
		err := faker.FakeData(&DownloadArquivosCVMRequest)
		assert.Nil(t, err)

		service.EXPECT().DownloadArquivosCVMService(gomock.Any()).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		)

		fundosController.DownloadArquivosCVMController(DownloadArquivosCVMRequest)
	})

}
