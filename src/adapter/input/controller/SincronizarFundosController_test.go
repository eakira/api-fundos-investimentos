package controller

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestSincronizarFundosController(t *testing.T) {

	service, fundosController := InitServiceTest(t)

	t.Run("when_sending_a_valid_request_returns_success", func(t *testing.T) {
		router := gin.Default()
		service.EXPECT().QueueFundosSincronizarService(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

		router.POST("/sincronizar-fundos", fundosController.SincronizarFundosController)

		req, err := http.NewRequest("POST", "/sincronizar-fundos", strings.NewReader("ignorar_download=true"))
		if err != nil {
			t.Fatalf("Erro ao criar request HTTP: %v", err)
		}

		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code, "Status code esperado não foi retornado")

	})

}
