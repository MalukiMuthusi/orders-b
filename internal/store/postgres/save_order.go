package postgres

import (
	"context"

	"github.com/MalukiMuthusi/orders-b/internal/logger"
	"github.com/MalukiMuthusi/orders-b/internal/models"
	"github.com/MalukiMuthusi/orders-b/internal/utils"
)

// SaveOrder saves an order record to the database
func (p Postgres) SaveOrder(ctx context.Context, order *models.Order) error {

	tx := p.Db.Create(order)

	if tx.Error != nil {

		logger.Log.Info(tx.Error)

		return tx.Error
	}

	if tx.RowsAffected != 1 {
		return utils.ErrInsertFailed
	}

	return nil
}
