package models

type RekeningResponse struct {
	Error   bool         `json:"error"`
	ReffID  string       `json:"reff_id"`
	Message string       `json:"message"`
	Data    RekeningBase `json:"data"`
}

type RekeningBase struct {
	Norek string  `json:"norek"`
	Saldo float64 `json:"saldo"`
}

type RekeningRequest struct {
	Norek string `json:"norek"`
}
