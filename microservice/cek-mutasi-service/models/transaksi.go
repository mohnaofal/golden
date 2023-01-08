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

type MutasiRequest struct {
	Norek     string `json:"norek"`
	StartDate int    `json:"start_date"`
	EndDate   int    `json:"end_date"`
}

type MutasiResponse struct {
	Date         int     `json:"date"`
	Type         string  `json:"type"`
	Gram         float64 `json:"gram"`
	HargaTopup   int     `json:"harga_topup"`
	HargaBuyback int     `json:"harga_buyback"`
	Saldo        float64 `json:"saldo"`
}

type TransaksiParams struct {
	Norek     string `json:"norek"`
	StartDate int    `json:"start_date"`
	EndDate   int    `json:"end_date"`
	OrderBy   string `json:"order_by"`
	SortBy    string `json:"sort_by"`
}
