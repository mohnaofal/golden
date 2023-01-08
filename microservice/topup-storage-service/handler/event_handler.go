package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/Shopify/sarama"
	"github.com/mohnaofal/golden/helper/kafka"
	"github.com/mohnaofal/golden/microservice/topup-storage-service/config"
	"github.com/mohnaofal/golden/microservice/topup-storage-service/models"
	"github.com/mohnaofal/golden/microservice/topup-storage-service/repository"
	"github.com/mohnaofal/golden/microservice/topup-storage-service/usecase"
)

type eventHandler struct {
	transaksiUsecase usecase.TransaksiUsecase
	kafkaConsumer    kafka.KafkaConsumer
}

type EventHandler interface {
	RegisterConsumer()
	GlobalHandler(ctx context.Context, msg *sarama.ConsumerMessage)
}

func NewEventHandler(cfg *config.Config) EventHandler {
	transaksiRepository := repository.NewTransaksiRepository(cfg)
	transaksiUsecase := usecase.NewTransaksiUsecase(transaksiRepository)
	return &eventHandler{
		transaksiUsecase: transaksiUsecase,
		kafkaConsumer:    cfg.KafkaConsumer(),
	}
}

func (h *eventHandler) RegisterConsumer() {
	// list topic
	topics := []string{
		models.TopicTopup,
	}

	signals := make(chan os.Signal, 1)

	// init Consume
	h.kafkaConsumer.Consume(topics, h.GlobalHandler, signals)

}

func (h *eventHandler) GlobalHandler(ctx context.Context, msg *sarama.ConsumerMessage) {
	switch msg.Topic {
	case models.TopicTopup:
		form := new(models.TopupRequest)
		err := json.Unmarshal(msg.Value, &form)
		if err != nil {
			fmt.Printf("GlobalHandler - topic : %s, err : %s \n", models.TopicTopup, err.Error())
			return
		}

		if err := h.transaksiUsecase.TopupStorage(ctx, form); err != nil {
			fmt.Printf("GlobalHandler-TopupStorage - topic : %s, err : %s \n", models.TopicTopup, err.Error())
		} else {
			fmt.Printf("GlobalHandler-TopupStorage - topic : %s - Success \n", models.TopicTopup)
		}
	}
}
