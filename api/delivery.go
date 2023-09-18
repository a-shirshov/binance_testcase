package api

import (
	"github.com/gin-gonic/gin"
)

type Delivery interface {
	FetchData(c *gin.Context)
}
