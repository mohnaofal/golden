package models

type HargaResponse struct {
	Error   bool      `json:"error"`
	ReffID  string    `json:"reff_id"`
	Message string    `json:"message"`
	Data    HargaBase `json:"data"`
}

type HargaBase struct {
	HargaTopup   int `json:"harga_topup"`
	HargaBuyback int `json:"harga_buyback"`
}
