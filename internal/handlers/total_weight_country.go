package handlers

import (
	"fmt"
	"net/http"

	"github.com/MalukiMuthusi/orders-b/internal/models"
	"github.com/MalukiMuthusi/orders-b/internal/store"
	"github.com/MalukiMuthusi/orders-b/internal/utils"
	"github.com/gin-gonic/gin"
)

// TotalWeightCountry gets the sum of parcel weight for the provided country
type TotalWeightCountry struct {
	Store store.Store
}

// Handle gets the sum of parcel weight for the provided country
func (h TotalWeightCountry) Handle(c *gin.Context) {

	// parse the query parameters

	var queryParameters models.TotalWeightCountry

	if err := c.ShouldBind(&queryParameters); err != nil {

		basicError := models.BasicError{
			Code:    utils.CodeProvideCountry,
			Message: "provide a country in query parameter",
		}

		c.JSON(http.StatusUnprocessableEntity, basicError)
		return
	}

	// call the storage to return the sum of parcel weights for the provided country
	total, err := h.Store.TotalWeightCountry(c.Request.Context(), &queryParameters)

	if err != nil {

		basicError := models.BasicError{
			Code:    utils.CodeFailedGetTotalWeight,
			Message: fmt.Sprintf("failed to get the sum of parcel weights for %s", queryParameters.Country),
		}

		c.JSON(http.StatusInternalServerError, basicError)
		return

	}

	// return the sum of parcel weights
	resp := struct {
		Total *float32 `json:"total"`
	}{
		Total: total,
	}

	c.JSON(http.StatusOK, resp)

}
