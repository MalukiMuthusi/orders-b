package postgres

import (
	"context"

	"github.com/MalukiMuthusi/orders-b/internal/logger"
	"github.com/MalukiMuthusi/orders-b/internal/models"
	"github.com/MalukiMuthusi/orders-b/internal/utils"
)

// BatchSaveOrders saves orders in batches
func (p Postgres) BatchSaveOrders(ctx context.Context, orders []*models.Order) error {

	tx := p.Db.Create(&orders)

	if tx.Error != nil {

		logger.Log.Info(tx.Error)

		return tx.Error
	}

	if tx.RowsAffected < int64(len(orders)) {
		return utils.ErrPartialInsert
	}

	return nil
}
