package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/mohnaofal/golden/helper/databases"
	"github.com/mohnaofal/golden/microservice/cek-harga-service/config"
	"github.com/mohnaofal/golden/microservice/cek-harga-service/models"
)

type hargaRepository struct {
	cfg *config.Config
	DB  databases.ConnPostgres
}

type HargaRepository interface {
	Get(ctx context.Context) (*models.Harga, error)
}

func NewHargaRepository(cfg *config.Config) HargaRepository {
	return &hargaRepository{
		cfg: cfg,
		DB:  cfg.SQLDB(),
	}
}

func (c *hargaRepository) Get(ctx context.Context) (*models.Harga, error) {
	var (
		res                                          = new(models.Harga)
		hargaId, hargaTopup, hargaBuyback, hargaDate int
		hargaAdminId                                 string
	)

	row := c.DB.SqlDB().QueryRow(`SELECT * FROM harga ORDER BY harga_id DESC`)
	if err := row.Scan(&hargaId, &hargaAdminId, &hargaTopup, &hargaBuyback, &hargaDate); err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, nil
		}
		return nil, err
	}

	res = &models.Harga{
		HargaID:      hargaId,
		HargaAdminID: hargaAdminId,
		HargaTopup:   hargaTopup,
		HargaBuyback: hargaBuyback,
		HargaDate:    hargaDate,
	}

	return res, nil
}
