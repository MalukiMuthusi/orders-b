package postgres

import (
	"context"

	"github.com/MalukiMuthusi/orders-b/internal/logger"
	"github.com/MalukiMuthusi/orders-b/internal/models"
)

// GetOrders returns orders paginated
func (p Postgres) GetOrders(ctx context.Context, filters *models.GetOrdersRequestQuery) (*models.GetOrdersResponse, error) {

	var orders []*models.Order

	db := p.Db.Limit(filters.PageSize).Offset(filters.Offset)

	if filters.Country != "" {
		db = db.Where("country = ?", filters.Country)
	}

	if filters.WeightLimit != 0 {
		db = db.Where("weight_limit < ?", filters.WeightLimit)

	}

	tx := db.Find(&orders)

	if tx.Error != nil {

		logger.Log.Info(tx.Error)

		return nil, tx.Error
	}

	// set up the new offset

	newOffset := filters.PageSize + filters.Offset

	// check if there are more order items
	if len(orders) < filters.PageSize {
		newOffset = -1
	}

	return &models.GetOrdersResponse{
		Orders: orders,
		Offset: newOffset,
	}, nil
}
