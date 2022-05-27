package handlers

import (
	"github.com/MalukiMuthusi/orders-b/internal/store"
	"github.com/gin-gonic/gin"
)

type DebugPrintRouteFunc func(httpMethod, absolutePath, handlerName string, nuHandlers int)

func SetUpRouter(store store.Store, debugPrintRoute DebugPrintRouteFunc) *gin.Engine {

	r := gin.New()

	// log the endpoints
	gin.DebugPrintRouteFunc = debugPrintRoute

	return r

}
