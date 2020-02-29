package routes

import (
	"gin-group-buy/server/service/products"
	"github.com/gin-gonic/gin"
)

// Engine - engine
func Engine() *gin.Engine {
	r := gin.Default()

	// line bot callback
	r.POST("/callback", products.PostHandler())

	return r
}
