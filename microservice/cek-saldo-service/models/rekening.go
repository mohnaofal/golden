package models

type Rekening struct {
	RekID        int     `json:"rek_id" db:"rek_id"`
	RekNorek     string  `json:"rek_norek" db:"rek_norek"`
	RekSaldo     float64 `json:"rek_saldo" db:"rek_saldo"`
	RekUpdatedAt int     `json:"rek_updated_at" db:"rek_updated_at"`
}

type RekeningRequest struct {
	Norek string `json:"norek"`
}

type RekeningResponse struct {
	Norek string  `json:"norek"`
	Saldo float64 `json:"saldo"`
}
