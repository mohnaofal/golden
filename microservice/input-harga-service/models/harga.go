package models

const (
	// topic input harga
	TopicInputHarga = `input-harga`
)

type Harga struct {
	HargaID      int    `json:"harga_id" db:"harga_id"`
	HargaAdminID string `json:"harga_admin_id" db:"harga_admin_id"`
	HargaTopup   int    `json:"harga_topup" db:"harga_topup"`
	HargaBuyback int    `json:"harga_buyback" db:"harga_buyback"`
	HargaDate    int    `json:"harga_date" db:"harga_date"`
}

type HargaRequest struct {
	AdminID      string `json:"admin_id"`
	HargaTopup   int    `json:"harga_topup"`
	HargaBuyback int    `json:"harga_buyback"`
}

type HargaResponse struct {
	HargaTopup   int `json:"harga_topup"`
	HargaBuyback int `json:"harga_buyback"`
}
