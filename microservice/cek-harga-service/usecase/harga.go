package usecase

import (
	"context"
	"errors"

	"github.com/mohnaofal/golden/microservice/cek-harga-service/models"
	"github.com/mohnaofal/golden/microservice/cek-harga-service/repository"
)

type hargaUsecase struct {
	hargaRepo repository.HargaRepository
}

type HargaUsecase interface {
	Detail(ctx context.Context) (*models.HargaResponse, error)
}

func NewHargaUsecase(hargaRepo repository.HargaRepository) HargaUsecase {
	return &hargaUsecase{
		hargaRepo: hargaRepo,
	}
}

func (c *hargaUsecase) Detail(ctx context.Context) (*models.HargaResponse, error) {
	query, err := c.hargaRepo.Get(ctx)
	if err != nil {
		return nil, err
	}
	if query == nil {
		return nil, errors.New("data not found")
	}

	data := &models.HargaResponse{
		HargaTopup:   query.HargaTopup,
		HargaBuyback: query.HargaBuyback,
	}

	return data, nil
}
