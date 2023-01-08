package models

import (
	"errors"
	"strconv"
	"strings"
)

const (
	TopicTopup = `topup`
)

type TopupRequest struct {
	Gram         float64 `json:"gram"`
	Harga        int     `json:"harga"`
	Norek        string  `json:"norek"`
	HargaBuyback int     `json:"harga_buyback,omitempty"`
}

func (c *TopupRequest) Validate() error {
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
