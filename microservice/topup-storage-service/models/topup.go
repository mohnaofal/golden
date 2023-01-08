package models

const (
	// TopicTopup
	TopicTopup = `topup`

	// TransaksiTypeTopup
	TransaksiTypeTopup = `topup`
)

type Topup struct {
	TopupID    int     `json:"topup_id" db:"topup_id"`
	TopupGram  float64 `json:"topup_gram" db:"topup_gram"`
	TopupHarga int     `json:"topup_harga" db:"topup_harga"`
	TopupNorek string  `json:"topup_norek" db:"topup_norek"`
	TopupDate  int     `json:"topup_date" db:"topup_date"`
}

type TopupRequest struct {
	Gram         float64 `json:"gram"`
	Harga        int     `json:"harga"`
	Norek        string  `json:"norek"`
	HargaBuyback int     `json:"harga_buyback"`
}
