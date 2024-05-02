package main

import (
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/test/mocks"
	"net/http"
	"os"
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestMain(t *testing.T) {

	t.Run("when_not_dot_env_return_error", func(t *testing.T) {
		err := os.Rename(".env", "Renomeando.env")
		assert.Nil(t, err)

		defer func() {
			err := os.Rename("Renomeando.env", ".env")
			if err != nil {
				t.Fatalf("Erro ao restaurar o nome do arquivo original: %s", err)
			}
		}()

		err = LoadDotEnv()
		assert.NotNil(t, err)
	})

	t.Run("when_dot_env_return_sucess", func(t *testing.T) {

		err := LoadDotEnv()
		assert.Nil(t, err)
	})

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

		go initHttp(router, server, fundosMock, mutex)

		resp, err := http.Get("http://localhost:4203/dev-tools/test")
		assert.Nil(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode, "O código de status não é o esperado")

	})
}
