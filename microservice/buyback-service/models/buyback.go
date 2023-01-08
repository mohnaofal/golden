package models

import (
	"errors"
	"strconv"
	"strings"
)

const (
	TopicBuyback = `buyback`
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
	HargaTopup int     `json:"harga_topup,omitempty"`
}

func (c *BuybackRequest) Validate() error {
	minGram := 0.001
	if c.Gram < minGram {
		return errors.New("gram tidak boleh kurang dari 0.001")
	}

	gramStr := strconv.FormatFloat(c.Gram, 'f', -1, 64)
	decimal := gramStr
	if strings.Contains(gramStr, ".") {
		decimal = strings.Split(gramStr, ".")[1]
	}

	if len(decimal) > 3 {
		return errors.New("gram harus kelipatan dari 0.001")
	}

	if c.Harga < 1 {
		return errors.New("harga tidak boleh kosong")
	}

	if c.Norek == `` {
		return errors.New("norek tidak boleh kosong")
	}

	return nil
}
