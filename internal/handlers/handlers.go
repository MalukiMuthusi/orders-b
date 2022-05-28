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

	// Save order

	saveOrder := SaveOrder{
		Store: store,
	}

	r.POST("saveorder", saveOrder.Handle)

	// Batch save orders

	batchSaveOrders := BatchSaveOrders{
		Store: store,
	}
	r.POST("batchsaveorders", batchSaveOrders.Handle)

	// Get orders, paginated and filtered

	getOrders := GetOrders{
		Store: store,
	}

	r.GET("orders", getOrders.Handle)

	// Get total orders for a given country

	getTotalOrdersCountry := TotalOrdersPerCountry{
		Store: store,
	}

	r.GET("totalorders", getTotalOrdersCountry.Handle)

	// Get sum of parcel weight for a country

	totalWeight := TotalWeightCountry{
		Store: store,
	}

	r.GET("totalweight", totalWeight.Handle)

	return r

}
