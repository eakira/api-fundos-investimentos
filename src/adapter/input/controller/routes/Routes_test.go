package routes

import (
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/test/mocks"
	"net/http"
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
	"go.uber.org/mock/gomock"
)

func TestMain(t *testing.T) {
	t.Run("when_init_http_return_sucess", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		logger.InitTestLogger()

		var mutex = &sync.Mutex{}

		gin.SetMode(gin.TestMode)
		router := gin.Default()

		server := &http.Server{
			Addr:    ":4203",
			Handler: router,
		}
		defer server.Close()

		fundosMock := mocks.NewMockFundosControllerInterface(ctrl)

		InitRoutes(&router.RouterGroup, fundosMock, mutex)
	})
}
