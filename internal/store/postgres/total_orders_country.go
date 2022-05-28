package postgres

import (
	"context"

	"github.com/MalukiMuthusi/orders-b/internal/logger"
	"github.com/MalukiMuthusi/orders-b/internal/models"
)

// GetTotalOrdersPerCountry returns the total orders for the given country
func (p Postgres) GetTotalOrdersPerCountry(ctx context.Context, r *models.TotalOrdersCountryQuery) (*int64, error) {

	var total int64

	tx := p.Db.Where("country = ?", r.Country).Count(&total)

	if tx.Error != nil {
		logger.Log.Info(tx.Error)
		return nil, tx.Error
	}

	return &total, nil
}
