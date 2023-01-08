package usecase

import (
	"context"
	"time"

	"github.com/mohnaofal/golden/microservice/buyback-storage-service/models"
	"github.com/mohnaofal/golden/microservice/buyback-storage-service/repository"
)

type transaksiUsecase struct {
	transaksiRepo repository.TransaksiRepository
}

type TransaksiUsecase interface {
	BuybackStorage(ctx context.Context, form *models.BuybackRequest) error
}

func NewTransaksiUsecase(transaksiRepo repository.TransaksiRepository) TransaksiUsecase {
	return &transaksiUsecase{transaksiRepo: transaksiRepo}
}

func (c *transaksiUsecase) BuybackStorage(ctx context.Context, form *models.BuybackRequest) error {
	data := &models.Transaksi{
		TrxDate:         int(time.Now().Unix()),
		TrxType:         models.TransaksiTypeBuyback,
		TrxGram:         form.Gram,
		TrxHargaTopup:   form.HargaTopup,
		TrxHargaBuyback: form.Harga,
		TrxNorek:        form.Norek,
	}

	if err := c.transaksiRepo.InsertTrx(ctx, data); err != nil {
		return err
	}

	return nil
}
