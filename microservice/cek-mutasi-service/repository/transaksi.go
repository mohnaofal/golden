package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/mohnaofal/golden/helper/databases"
	"github.com/mohnaofal/golden/microservice/cek-mutasi-service/config"
	"github.com/mohnaofal/golden/microservice/cek-mutasi-service/models"
)

type transaksiRepository struct {
	cfg *config.Config
	DB  databases.ConnPostgres
}

type TransaksiRepository interface {
	Select(ctx context.Context, params *models.TransaksiParams) ([]models.Transaksi, error)
}

func NewTransaksiRepository(cfg *config.Config) TransaksiRepository {
	return &transaksiRepository{
		cfg: cfg,
		DB:  cfg.SQLDB(),
	}
}

func (c *transaksiRepository) Select(ctx context.Context, params *models.TransaksiParams) ([]models.Transaksi, error) {
	var (
		res            = make([]models.Transaksi, 0)
		queryCondition = ``
	)

	if params.Norek != `` {
		queryCondition += func() string {
			if len(queryCondition) > 0 {
				return fmt.Sprintf(` AND trx_norek = '%s'`, params.Norek)
			}
			return fmt.Sprintf(`WHERE trx_norek = '%s'`, params.Norek)
		}()
	}

	if params.StartDate > 0 {
		queryCondition += func() string {
			if len(queryCondition) > 0 {
				return fmt.Sprintf(` AND trx_date >= '%v'`, params.StartDate)
			}
			return fmt.Sprintf(`WHERE trx_date >= '%v'`, params.StartDate)
		}()
	}

	if params.EndDate > 0 {
		queryCondition += func() string {
			if len(queryCondition) > 0 {
				return fmt.Sprintf(` AND trx_date <= '%v'`, params.EndDate)
			}
			return fmt.Sprintf(`WHERE trx_date <= '%v'`, params.EndDate)
		}()
	}

	if params.OrderBy != `` && params.SortBy != `` {
		queryCondition += fmt.Sprintf(` ORDER BY %s %s`, params.OrderBy, params.SortBy)
	}

	sqlPrepare, err := c.DB.SqlDB().Prepare(`SELECT * FROM transaksi ` + queryCondition)
	if err != nil {
		return res, err
	}

	rows, err := sqlPrepare.Query()
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return res, nil
		}
		return res, err
	}

	for rows.Next() {
		var (
			trxID, trxDate, trxHargaTopup, trxHargaBuyback int
			trxType, trxNorek                              string
			trxGram, trxSaldo                              float64
		)

		if err = rows.Scan(&trxID, &trxDate, &trxType, &trxGram, &trxHargaTopup, &trxHargaBuyback, &trxSaldo, &trxNorek); err != nil {
			return res, err
		}

		res = append(res, models.Transaksi{
			TrxID:           trxID,
			TrxDate:         trxDate,
			TrxType:         trxType,
			TrxGram:         trxGram,
			TrxHargaTopup:   trxHargaTopup,
			TrxHargaBuyback: trxHargaBuyback,
			TrxSaldo:        trxSaldo,
			TrxNorek:        trxNorek,
		})
	}

	return res, nil
}
