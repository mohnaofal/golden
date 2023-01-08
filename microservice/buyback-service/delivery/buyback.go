package delivery

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mohnaofal/golden/helper/request"
	"github.com/mohnaofal/golden/helper/response"
	"github.com/mohnaofal/golden/microservice/buyback-service/config"
	"github.com/mohnaofal/golden/microservice/buyback-service/models"
	"github.com/mohnaofal/golden/microservice/buyback-service/repository"
	"github.com/mohnaofal/golden/microservice/buyback-service/usecase"
)

type BuybackDelivery struct {
	buybackUsecase usecase.BuybackUsecase
}

func NewBuybackDelivery(cfg *config.Config) BuybackDelivery {
	hargaRepository := repository.NewHargaRepository(cfg)
	rekeningRepository := repository.NewRekeningRepository(cfg)
	buybackUsecase := usecase.NewBuybackUsecase(hargaRepository, rekeningRepository, cfg.KafkaProcedur())
	return BuybackDelivery{buybackUsecase: buybackUsecase}
}

func (h *BuybackDelivery) Apply(c *mux.Router) {
	c.HandleFunc("/buyback", h.Buyback).Methods(http.MethodPost)
}

func (h *BuybackDelivery) Buyback(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	form := new(models.BuybackRequest)
	if err := request.Form(r, form); err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	if err := h.buybackUsecase.Buyback(ctx, form); err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	response.Success(w, http.StatusOK, "", nil)
}
