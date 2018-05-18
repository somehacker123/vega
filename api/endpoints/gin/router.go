package gin

import (
	"github.com/gin-gonic/gin"
	"vega/api/services"
	"fmt"
)

func NewRouter(orderService services.OrderService) *gin.Engine  {
	gin.SetMode(gin.TestMode)

	fmt.Println(orderService)

	// Set up HTTP router and route handlers
	httpRouter := gin.New()
	httpHandlers := Handlers{
		OrderService: orderService,
	}

	httpRouter.GET(httpHandlers.IndexRoute(), httpHandlers.Index)
	httpRouter.GET(httpHandlers.CreateOrderRoute(), httpHandlers.CreateOrder)

	return httpRouter
}
