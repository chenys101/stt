package route

import (
	"github.com/gin-gonic/gin"
	"trace-stock/internal/controller"
	"trace-stock/internal/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.ContentNegotiation())

	api := r.Group("/api/v1")
	{
		users := api.Group("/users")
		{
			uc := controller.UserController{}
			users.POST("", uc.CreateUser)
			users.GET("/:id", uc.GetUser)
		}
		stocks := api.Group("/std")
		{
			sc := controller.StockController{}
			stocks.GET("", sc.GetStockData)
			stocks.GET("/nc", sc.GetStockDataNotEncrypt)
			stocks.POST("/createStdMonitor", sc.CreateStockMonitor)
			stocks.POST("/updateStdMonitor", sc.UpdateStockMonitor)
			stocks.GET("/getAllStdMonitors", sc.GetAllStockMonitors)
		}
	}

	return r
}
