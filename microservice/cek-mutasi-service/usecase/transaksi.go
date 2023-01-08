package usecase

import (
	"context"

	"github.com/mohnaofal/golden/microservice/cek-mutasi-service/models"
	"github.com/mohnaofal/golden/microservice/cek-mutasi-service/repository"
)

type transaksiUsecase struct {
	transaksiRepo repository.TransaksiRepository
}

type TransaksiUsecase interface {
	CekMutasi(ctx context.Context, form *models.MutasiRequest) ([]models.MutasiResponse, error)
}

func NewTransaksiUsecase(transaksiRepo repository.TransaksiRepository) TransaksiUsecase {
	return &transaksiUsecase{
		transaksiRepo: transaksiRepo,
	}
}

func (c *transaksiUsecase) CekMutasi(ctx context.Context, form *models.MutasiRequest) ([]models.MutasiResponse, error) {
	params := &models.TransaksiParams{
		Norek:     form.Norek,
		StartDate: form.StartDate,
		EndDate:   form.EndDate,
		OrderBy:   `trx_id`,
		SortBy:    `DESC`,
	}

	query, err := c.transaksiRepo.Select(ctx, params)
	if err != nil {
		return nil, err
	}

	data := make([]models.MutasiResponse, 0)
	for _, v := range query {
		data = append(data, models.MutasiResponse{
			Date:         v.TrxDate,
			Type:         v.TrxType,
			Gram:         v.TrxGram,
			HargaTopup:   v.TrxHargaTopup,
			HargaBuyback: v.TrxHargaBuyback,
			Saldo:        v.TrxSaldo,
		})
	}

	return data, nil
}
