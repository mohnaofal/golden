package models

type Transaksi struct {
	TrxID           int     `json:"trx_id" db:"trx_id"`
	TrxDate         int     `json:"trx_date" db:"trx_date"`
	TrxType         string  `json:"trx_type" db:"trx_type"`
	TrxGram         float64 `json:"trx_gram" db:"trx_gram"`
	TrxHargaTopup   int     `json:"trx_harga_topup" db:"trx_harga_topup"`
	TrxHargaBuyback int     `json:"trx_harga_buyback" db:"trx_harga_buyback"`
	TrxSaldo        float64 `json:"trx_saldo" db:"trx_saldo"`
	TrxNorek        string  `json:"trx_norek" db:"trx_norek"`
}
