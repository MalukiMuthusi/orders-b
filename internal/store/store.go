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
}
