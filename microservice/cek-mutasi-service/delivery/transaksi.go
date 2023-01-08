package delivery

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mohnaofal/golden/helper/request"
	"github.com/mohnaofal/golden/helper/response"
	"github.com/mohnaofal/golden/microservice/cek-mutasi-service/config"
	"github.com/mohnaofal/golden/microservice/cek-mutasi-service/models"
	"github.com/mohnaofal/golden/microservice/cek-mutasi-service/repository"
	"github.com/mohnaofal/golden/microservice/cek-mutasi-service/usecase"
)

type TransaksiDelivery struct {
	transaksiUsecase usecase.TransaksiUsecase
}

func NewTransaksiDelivery(cfg *config.Config) TransaksiDelivery {
	transaksiRepository := repository.NewTransaksiRepository(cfg)
	transaksiUsecase := usecase.NewTransaksiUsecase(transaksiRepository)
	return TransaksiDelivery{transaksiUsecase: transaksiUsecase}
}

func (h *TransaksiDelivery) Apply(c *mux.Router) {
	c.HandleFunc("/mutasi", h.CekMutasi).Methods(http.MethodPost)
}

func (h *TransaksiDelivery) CekMutasi(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	form := new(models.MutasiRequest)
	if err := request.Form(r, form); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.transaksiUsecase.CekMutasi(ctx, form)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(w, http.StatusOK, "", result)
}
