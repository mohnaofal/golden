package delivery

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mohnaofal/golden/helper/request"
	"github.com/mohnaofal/golden/helper/response"
	"github.com/mohnaofal/golden/microservice/cek-saldo-service/config"
	"github.com/mohnaofal/golden/microservice/cek-saldo-service/models"
	"github.com/mohnaofal/golden/microservice/cek-saldo-service/repository"
	"github.com/mohnaofal/golden/microservice/cek-saldo-service/usecase"
)

type RekeningDelivery struct {
	rekeningUsecase usecase.RekeningUsecase
}

func NewRekeningDelivery(cfg *config.Config) RekeningDelivery {
	rekeningRepository := repository.NewRekeningRepository(cfg)
	rekeningUsecase := usecase.NewRekeningUsecase(rekeningRepository)
	return RekeningDelivery{rekeningUsecase: rekeningUsecase}
}

func (h *RekeningDelivery) Apply(c *mux.Router) {
	c.HandleFunc("/saldo", h.CekSaldo).Methods(http.MethodPost)
}

func (h *RekeningDelivery) CekSaldo(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	form := new(models.RekeningRequest)
	if err := request.Form(r, form); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.rekeningUsecase.CekSaldo(ctx, form)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(w, http.StatusOK, "", result)
}
