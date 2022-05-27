package handlers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/MalukiMuthusi/orders-b/internal/logger"
	"github.com/MalukiMuthusi/orders-b/internal/models"
	"github.com/MalukiMuthusi/orders-b/internal/store"
	"github.com/MalukiMuthusi/orders-b/internal/utils"
	"github.com/gin-gonic/gin"
)

// BatchSaveOrders saves orders in batches
type BatchSaveOrders struct {
	Store store.Store
}

// Handle save orders in batches
func (h BatchSaveOrders) Handle(c *gin.Context) {

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
	var orders []*models.Order

	err = json.Unmarshal(body, &orders)
	if err != nil {

		logger.Log.Info(err)

		basicError := models.BasicError{
			Code:    utils.CodeInvalidRequestBody,
			Message: "provide a valid order in request body",
		}

		c.JSON(http.StatusUnprocessableEntity, basicError)
		return
	}

	// save the orders
	err = h.Store.BatchSaveOrders(c.Request.Context(), orders)

	if err != nil {

		if errors.Is(err, utils.ErrPartialInsert) {
			c.JSON(http.StatusOK, models.SaveResponse{Status: "PARTIALLY_SAVED"})
			return
		}

		c.JSON(http.StatusInternalServerError, models.BasicError{Code: utils.CodeFailedSaveOrder, Message: "failed to save the orders"})

		return
	}

	// return a success response
	resp := models.SaveResponse{
		Status: "SUCCESS",
	}

	c.JSON(http.StatusOK, resp)
}
