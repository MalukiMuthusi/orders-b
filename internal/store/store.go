package store

import (
	"context"

	"github.com/MalukiMuthusi/orders-b/internal/models"
)

// Store repository abstraction
type Store interface {

	// SaveOrder to the storage
	SaveOrder(ctx context.Context, order *models.Order) error

	// BatchSaveOrders saves a batch of orders
	BatchSaveOrders(ctx context.Context, orders []*models.Order) error

	// GetOrders returns a list of paginated orders
	GetOrders(ctx context.Context, filters *models.GetOrdersRequestQuery) (*models.GetOrdersResponse, error)

	// GetTotalOrdersPerCountry returns the total orders for the given country
	GetTotalOrdersPerCountry(ctx context.Context, r *models.TotalOrdersCountryQuery) (*int64, error)
}
