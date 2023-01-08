package usecase

import (
	"context"
	"time"

	"github.com/mohnaofal/golden/microservice/input-harga-storage-service/models"
	"github.com/mohnaofal/golden/microservice/input-harga-storage-service/repository"
)

type hargaUsecase struct {
	hargaRepo repository.HargaRepository
}

type HargaUsecase interface {
	InputHargaStorage(ctx context.Context, form *models.HargaRequest) error
}

func NewHargaUsecase(hargaRepo repository.HargaRepository) HargaUsecase {
	return &hargaUsecase{hargaRepo: hargaRepo}
}

func (c *hargaUsecase) InputHargaStorage(ctx context.Context, form *models.HargaRequest) error {
	data := &models.Harga{
		HargaAdminID: form.AdminID,
		HargaTopup:   form.HargaTopup,
		HargaBuyback: form.HargaBuyback,
		HargaDate:    int(time.Now().Unix()),
	}

	if err := c.hargaRepo.Insert(ctx, data); err != nil {
		return err
	}

	return nil
}
