package handlers

import (
	"net/http"

	"github.com/MalukiMuthusi/orders-b/internal/models"
	"github.com/MalukiMuthusi/orders-b/internal/store"
	"github.com/gin-gonic/gin"
)

// GetOrders returns a list of paginated orders
type GetOrders struct {
	Store store.Store
}

// Handle GetOrders returns a list of paginated orders
func (h GetOrders) Handle(c *gin.Context) {

	// parse the request query parameter values

	var queryValues models.GetOrdersRequestQuery

	if err := c.ShouldBind(&queryValues); err != nil {
		queryValues = models.GetOrdersRequestQuery{
			Offset:   0,
			PageSize: 100,
		}
	}

	resp, err := h.Store.GetOrders(c.Request.Context(), &queryValues)
	if err != nil {
		basicError := models.BasicError{
			Code:    "FAILED_GET_ORDERS",
			Message: "failed to get orders",
		}

		c.JSON(http.StatusInternalServerError, basicError)
		return
	}

	c.JSON(http.StatusOK, resp)
}
