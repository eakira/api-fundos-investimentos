package routes

import (
	"api-fundos-investimentos/adapter/input/controller"

	"github.com/gin-gonic/gin"
)

func InitRoutes(
	r *gin.RouterGroup,
	fundosController controller.FundosControllerInterface) {

	v1 := r.Group("/api/v1/fundos")
	{
		v1.GET("sincronizar", fundosController.SincronizarFundosController)

	}

}
