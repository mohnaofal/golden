package delivery

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mohnaofal/golden/helper/response"
	"github.com/mohnaofal/golden/microservice/cek-harga-service/config"
	"github.com/mohnaofal/golden/microservice/cek-harga-service/repository"
	"github.com/mohnaofal/golden/microservice/cek-harga-service/usecase"
)

type HargaDelivery struct {
	hargaUsecase usecase.HargaUsecase
}

func NewHargaDelivery(cfg *config.Config) HargaDelivery {
	hargaRepository := repository.NewHargaRepository(cfg)
	hargaUsecase := usecase.NewHargaUsecase(hargaRepository)
	return HargaDelivery{hargaUsecase: hargaUsecase}
}

func (h *HargaDelivery) Apply(c *mux.Router) {
	c.HandleFunc("/check-harga", h.CekHarga).Methods(http.MethodPost)
}

func (h *HargaDelivery) CekHarga(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	result, err := h.hargaUsecase.Detail(ctx)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(w, http.StatusOK, "", result)
}
