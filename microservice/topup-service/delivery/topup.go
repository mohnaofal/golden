package delivery

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mohnaofal/golden/helper/request"
	"github.com/mohnaofal/golden/helper/response"
	"github.com/mohnaofal/golden/microservice/topup-service/config"
	"github.com/mohnaofal/golden/microservice/topup-service/models"
	"github.com/mohnaofal/golden/microservice/topup-service/repository"
	"github.com/mohnaofal/golden/microservice/topup-service/usecase"
)

type TopupDelivery struct {
	topupUsecase usecase.TopupUsecase
}

func NewTopupDelivery(cfg *config.Config) TopupDelivery {
	hargaRepository := repository.NewHargaRepository(cfg)
	topupUsecase := usecase.NewTopupUsecase(hargaRepository, cfg.KafkaProcedur())
	return TopupDelivery{topupUsecase: topupUsecase}
}

func (h *TopupDelivery) Apply(c *mux.Router) {
	c.HandleFunc("/topup", h.Topup).Methods(http.MethodPost)
}

func (h *TopupDelivery) Topup(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	form := new(models.TopupRequest)
	if err := request.Form(r, form); err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	if err := h.topupUsecase.Topup(ctx, form); err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	response.Success(w, http.StatusOK, "", nil)
}
