package handlers

import (
	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {

	r := gin.Default()
	r.GET("/", startPage)
	return r
}
