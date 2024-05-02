package routes

import (
	"api-fundos-investimentos/adapter/input/controller"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

func InitRoutes(
	r *gin.RouterGroup,
	fundosController controller.FundosControllerInterface,
	mutex *sync.Mutex,
) {
	r.Use(func(c *gin.Context) {
		mutex.Lock()
		defer mutex.Unlock()
		c.Next()
	})

	v1 := r.Group("/api/v1/fundos")
	{
		v1.POST("sincronizar", fundosController.SincronizarFundosController)
	}

	acessoDev := r.Group("/dev-tools")
	{
		acessoDev.GET("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status": "up",
			})
		})

	}

}
