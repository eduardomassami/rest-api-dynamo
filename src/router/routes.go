package router

import (
	"github.com/eduardomassami/rest-api-dynamo/adapter/http"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine, handler *http.Handler) {

	basePath := "/users"
	v1 := router.Group(basePath)

	v1.POST("", handler.CreateUser)
	v1.GET("/:id", handler.GetUser)
}
