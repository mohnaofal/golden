package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/mohnaofal/golden/helper/databases"
	"github.com/mohnaofal/golden/microservice/cek-saldo-service/config"
	"github.com/mohnaofal/golden/microservice/cek-saldo-service/models"
)

type rekeningRepository struct {
	cfg *config.Config
	DB  databases.ConnPostgres
}

type RekeningRepository interface {
	GetByNorek(ctx context.Context, norek string) (*models.Rekening, error)
}

func NewRekeningRepository(cfg *config.Config) RekeningRepository {
	return &rekeningRepository{
		cfg: cfg,
		DB:  cfg.SQLDB(),
	}
}

func (c *rekeningRepository) GetByNorek(ctx context.Context, norek string) (*models.Rekening, error) {
	var (
		res                 = new(models.Rekening)
		rekId, rekUpdatedAt int
		rekNorek            string
		rekSaldo            float64
	)

	row := c.DB.SqlDB().QueryRow(`SELECT * FROM rekening WHERE rek_norek = $1`, norek)
	if err := row.Scan(&rekId, &rekNorek, &rekSaldo, &rekUpdatedAt); err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, nil
		}
		return nil, err
	}

	res = &models.Rekening{
		RekID:        rekId,
		RekNorek:     norek,
		RekSaldo:     rekSaldo,
		RekUpdatedAt: rekUpdatedAt,
	}

	return res, nil
}
