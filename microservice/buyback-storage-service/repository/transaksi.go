package repository

import (
	"context"

	"github.com/mohnaofal/golden/helper/databases"
	"github.com/mohnaofal/golden/microservice/buyback-storage-service/config"
	"github.com/mohnaofal/golden/microservice/buyback-storage-service/models"
)

type transaksiRepository struct {
	cfg *config.Config
	DB  databases.ConnPostgres
}

type TransaksiRepository interface {
	InsertTrx(ctx context.Context, data *models.Transaksi) error
}

func NewTransaksiRepository(cfg *config.Config) TransaksiRepository {
	return &transaksiRepository{
		cfg: cfg,
		DB:  cfg.SQLDB(),
	}
}

func (c *transaksiRepository) InsertTrx(ctx context.Context, data *models.Transaksi) error {
	tx, err := c.DB.SqlDB().Begin()
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			return
		}

		if err != nil {
			tx.Rollback()
			return
		}

		if err := tx.Commit(); err != nil {
			return
		}
	}()

	if _, err = tx.Exec(`INSERT INTO buyback (buyback_gram, buyback_harga, buyback_norek, buyback_date) VALUES ($1, $2, $3, $4)`,
		data.TrxGram,
		data.TrxHargaTopup,
		data.TrxNorek,
		data.TrxDate,
	); err != nil {
		return err
	}

	saldo := 0.0
	if err = tx.QueryRow(`UPDATE rekening SET rek_saldo = rek_saldo - $1, rek_updated_at = $2  WHERE rek_norek = $3 RETURNING rek_saldo`,
		data.TrxGram,
		data.TrxDate,
		data.TrxNorek,
	).Scan(&saldo); err != nil {
		return err
	}

	data.TrxSaldo = saldo

	if _, err = tx.Exec(`INSERT INTO transaksi (trx_date, trx_type, trx_gram, trx_harga_topup, trx_harga_buyback, trx_saldo, trx_norek) VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		data.TrxDate,
		data.TrxType,
		data.TrxGram,
		data.TrxHargaTopup,
		data.TrxHargaBuyback,
		data.TrxSaldo,
		data.TrxNorek,
	); err != nil {
		return err
	}

	return nil
}
