package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mohnaofal/golden/helper/databases"
	"github.com/mohnaofal/golden/helper/network"
	"github.com/mohnaofal/golden/microservice/topup-service/config"
	"github.com/mohnaofal/golden/microservice/topup-service/models"
)

type hargaRepository struct {
	cfg *config.Config
	DB  databases.ConnPostgres
}

type HargaRepository interface {
	Get(ctx context.Context) (*models.HargaResponse, error)
}

func NewHargaRepository(cfg *config.Config) HargaRepository {
	return &hargaRepository{
		cfg: cfg,
		DB:  cfg.SQLDB(),
	}
}

func (c *hargaRepository) Get(ctx context.Context) (*models.HargaResponse, error) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	url := fmt.Sprintf("%s/api/check-harga", c.cfg.HargaSvrHost())
	resp, code, err := network.CallHTTPRequest(ctx, http.MethodPost, url, nil, headers)
	if err != nil || !network.IsStatusSuccess[code] {
		return nil, err
	}

	harga := new(models.HargaResponse)
	if err = json.Unmarshal(resp, harga); err != nil {
		return nil, err
	}

	return harga, nil
}
