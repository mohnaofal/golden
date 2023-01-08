package models

const (
	TopicBuyback = `buyback`

	TransaksiTypeBuyback = `buyback`
)

type Buyback struct {
	BuybackID    int     `json:"buyback_id" db:"buyback_id"`
	BuybackGram  float64 `json:"buyback_gram" db:"buyback_gram"`
	BuybackHarga int     `json:"buyback_harga" db:"buyback_harga"`
	BuybackNorek string  `json:"buyback_norek" db:"buyback_norek"`
	BuybackDate  int     `json:"buyback_date" db:"buyback_date"`
}

type BuybackRequest struct {
	Gram       float64 `json:"gram"`
	Harga      int     `json:"harga"`
	Norek      string  `json:"norek"`
	HargaTopup int     `json:"harga_topup"`
}
