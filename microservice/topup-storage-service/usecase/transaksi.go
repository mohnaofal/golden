package usecase

import (
	"context"
	"time"

	"github.com/mohnaofal/golden/microservice/topup-storage-service/models"
	"github.com/mohnaofal/golden/microservice/topup-storage-service/repository"
)

type transaksiUsecase struct {
	transaksiRepo repository.TransaksiRepository
}

type TransaksiUsecase interface {
	TopupStorage(ctx context.Context, form *models.TopupRequest) error
}

func NewTransaksiUsecase(transaksiRepo repository.TransaksiRepository) TransaksiUsecase {
	return &transaksiUsecase{transaksiRepo: transaksiRepo}
}

func (c *transaksiUsecase) TopupStorage(ctx context.Context, form *models.TopupRequest) error {
	data := &models.Transaksi{
		TrxDate:         int(time.Now().Unix()),
		TrxType:         models.TransaksiTypeTopup,
		TrxGram:         form.Gram,
		TrxHargaTopup:   form.Harga,
		TrxHargaBuyback: form.HargaBuyback,
		TrxNorek:        form.Norek,
	}

	if err := c.transaksiRepo.InsertTrx(ctx, data); err != nil {
		return err
	}

	return nil
}
