package postgres

import (
	"context"

	"github.com/MalukiMuthusi/orders-b/internal/logger"
	"github.com/MalukiMuthusi/orders-b/internal/models"
)

// TotalWeightCountry returns the sum of parcel_weights for the provided country
func (p Postgres) TotalWeightCountry(ctx context.Context, r *models.TotalWeightCountry) (*float32, error) {

	var total float32

	tx := p.Db.Where("country = ?", r.Country).Select("sum(parcel_weight) as total").Scan(&total)

	if tx.Error != nil {

		logger.Log.Info(tx.Error)

		return nil, tx.Error
	}

	return &total, nil
}
