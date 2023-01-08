package usecase

import (
	"context"
	"errors"

	"github.com/mohnaofal/golden/microservice/cek-saldo-service/models"
	"github.com/mohnaofal/golden/microservice/cek-saldo-service/repository"
)

type rekeningUsecase struct {
	rekeningRepo repository.RekeningRepository
}

type RekeningUsecase interface {
	CekSaldo(ctx context.Context, form *models.RekeningRequest) (*models.RekeningResponse, error)
}

func NewRekeningUsecase(rekeningRepo repository.RekeningRepository) RekeningUsecase {
	return &rekeningUsecase{
		rekeningRepo: rekeningRepo,
	}
}

func (c *rekeningUsecase) CekSaldo(ctx context.Context, form *models.RekeningRequest) (*models.RekeningResponse, error) {
	query, err := c.rekeningRepo.GetByNorek(ctx, form.Norek)
	if err != nil {
		return nil, err
	}
	if query == nil {
		return nil, errors.New("data not found")
	}

	data := &models.RekeningResponse{
		Norek: query.RekNorek,
		Saldo: query.RekSaldo,
	}

	return data, nil
}
