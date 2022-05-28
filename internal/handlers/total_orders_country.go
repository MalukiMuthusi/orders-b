package handlers

import (
	"fmt"
	"net/http"

	"github.com/MalukiMuthusi/orders-b/internal/models"
	"github.com/MalukiMuthusi/orders-b/internal/store"
	"github.com/MalukiMuthusi/orders-b/internal/utils"
	"github.com/gin-gonic/gin"
)

// TotalOrdersPerCountry get the total number of orders for a given country
type TotalOrdersPerCountry struct {
	Store store.Store
}

// TotalOrdersPerCountry get the total number of orders for a given country
func (h TotalOrdersPerCountry) Handle(c *gin.Context) {

	// parse the request query parameters
	var queryParameter models.TotalOrdersCountryQuery

	if err := c.ShouldBind(&queryParameter); err != nil {
		basicError := models.BasicError{
			Code:    utils.CodeProvideCountry,
			Message: "provide a country in query parameter",
		}

		c.JSON(http.StatusUnprocessableEntity, basicError)
		return
	}

	// call the storage to return the total orders for the country
	total, err := h.Store.GetTotalOrdersPerCountry(c.Request.Context(), &queryParameter)

	if err != nil {
		basicError := models.BasicError{
			Code:    utils.CodeFailedGetTotalOrders,
			Message: fmt.Sprintf("failed to get the total orders for %s", queryParameter.Country),
		}

		c.JSON(http.StatusInternalServerError, basicError)
		return
	}

	resp := struct {
		Total *int64 `json:"total"`
	}{
		Total: total,
	}

	c.JSON(http.StatusOK, resp)
}
