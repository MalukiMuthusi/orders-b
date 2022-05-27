package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/MalukiMuthusi/orders-b/internal/logger"
	"github.com/MalukiMuthusi/orders-b/internal/models"
	"github.com/MalukiMuthusi/orders-b/internal/store"
	"github.com/MalukiMuthusi/orders-b/internal/utils"
	"github.com/gin-gonic/gin"
)

// SaveOrder handler for saving a record to the database
type SaveOrder struct {
	Store store.Store
}

// Handle save a record to the database
func (h SaveOrder) Handle(c *gin.Context) {

	// parse the request body
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {

		logger.Log.Info(err)

		basicError := models.BasicError{
			Code:    utils.CodeInvalidRequestBody,
			Message: "provide a valid order in request body",
		}

		c.JSON(http.StatusUnprocessableEntity, basicError)
		return
	}

	// unmarshal the request body
	var order models.Order

	err = json.Unmarshal(body, &order)
	if err != nil {

		logger.Log.Info(err)

		basicError := models.BasicError{
			Code:    utils.CodeInvalidRequestBody,
			Message: "provide a valid order in request body",
		}

		c.JSON(http.StatusUnprocessableEntity, basicError)
		return
	}

	// save the order to the storage
	err = h.Store.SaveOrder(c.Request.Context(), &order)
	if err != nil {
		logger.Log.Info(err)

		basicError := models.BasicError{
			Code:    utils.CodeFailedSaveOrder,
			Message: "failed to save the order",
		}

		c.JSON(http.StatusInternalServerError, basicError)
		return

	}

	// return a success response
	resp := models.SaveResponse{
		Status: "SUCCESS",
	}

	c.JSON(http.StatusOK, resp)

}
