package delivery

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mohnaofal/golden/helper/request"
	"github.com/mohnaofal/golden/helper/response"
	"github.com/mohnaofal/golden/microservice/input-harga-service/config"
	"github.com/mohnaofal/golden/microservice/input-harga-service/models"
	"github.com/mohnaofal/golden/microservice/input-harga-service/usecase"
)

type HargaDelivery struct {
	hargaUsecase usecase.HargaUsecase
}

func NewHargaDelivery(cfg *config.Config) HargaDelivery {
	hargaUsecase := usecase.NewHargaUsecase(cfg.KafkaProcedur())
	return HargaDelivery{hargaUsecase: hargaUsecase}
}

func (h *HargaDelivery) Apply(c *mux.Router) {
	c.HandleFunc("/input-harga", h.InputHarga).Methods(http.MethodPost)
}

func (h *HargaDelivery) InputHarga(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	form := new(models.HargaRequest)
	if err := request.Form(r, form); err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	if err := h.hargaUsecase.InputHarga(ctx, form); err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	response.Success(w, http.StatusOK, "", nil)
}
