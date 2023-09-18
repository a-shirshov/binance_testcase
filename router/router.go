package router

import (
	"binance_testcase/api"
	"github.com/gin-gonic/gin"
)

func BinanceEndpoints(r *gin.RouterGroup, aD api.Delivery) {
	r.GET("/fetch", aD.FetchData)
}
