package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mohnaofal/golden/helper/network"
	"github.com/mohnaofal/golden/microservice/buyback-service/config"
	"github.com/mohnaofal/golden/microservice/buyback-service/models"
)

type rekeningRepository struct {
	cfg *config.Config
}

type RekeningRepository interface {
	Get(ctx context.Context, params *models.RekeningRequest) (*models.RekeningResponse, error)
}

func NewRekeningRepository(cfg *config.Config) RekeningRepository {
	return &rekeningRepository{
		cfg: cfg,
	}
}

func (c *rekeningRepository) Get(ctx context.Context, params *models.RekeningRequest) (*models.RekeningResponse, error) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	bodyRequest, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/api/saldo", c.cfg.SaldoSvrHost())
	resp, code, err := network.CallHTTPRequest(ctx, http.MethodPost, url, bytes.NewBuffer(bodyRequest), headers)
	if err != nil || !network.IsStatusSuccess[code] {
		return nil, err
	}

	rekening := new(models.RekeningResponse)
	if err = json.Unmarshal(resp, rekening); err != nil {
		return nil, err
	}

	return rekening, nil
}
