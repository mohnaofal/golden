package repository

import (
	"context"

	"github.com/mohnaofal/golden/helper/databases"
	"github.com/mohnaofal/golden/microservice/input-harga-storage-service/config"
	"github.com/mohnaofal/golden/microservice/input-harga-storage-service/models"
)

type hargaRepository struct {
	cfg *config.Config
	DB  databases.ConnPostgres
}

type HargaRepository interface {
	Insert(ctx context.Context, data *models.Harga) error
}

func NewHargaRepository(cfg *config.Config) HargaRepository {
	return &hargaRepository{
		cfg: cfg,
		DB:  cfg.SQLDB(),
	}
}

func (c *hargaRepository) Insert(ctx context.Context, data *models.Harga) error {
	if _, err := c.DB.SqlDB().
		Exec(`INSERT INTO harga (harga_admin_id, harga_topup, harga_buyback, harga_date) VALUES($1, $2, $3, $4)`,
			data.HargaAdminID,
			data.HargaTopup,
			data.HargaBuyback,
			data.HargaDate,
		); err != nil {
		return err
	}

	return nil
}
