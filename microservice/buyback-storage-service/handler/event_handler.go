package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/Shopify/sarama"
	"github.com/mohnaofal/golden/helper/kafka"
	"github.com/mohnaofal/golden/microservice/buyback-storage-service/config"
	"github.com/mohnaofal/golden/microservice/buyback-storage-service/models"
	"github.com/mohnaofal/golden/microservice/buyback-storage-service/repository"
	"github.com/mohnaofal/golden/microservice/buyback-storage-service/usecase"
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
		models.TopicBuyback,
	}

	signals := make(chan os.Signal, 1)

	// init Consume
	h.kafkaConsumer.Consume(topics, h.GlobalHandler, signals)

}

func (h *eventHandler) GlobalHandler(ctx context.Context, msg *sarama.ConsumerMessage) {
	switch msg.Topic {
	case models.TopicBuyback:
		form := new(models.BuybackRequest)
		err := json.Unmarshal(msg.Value, &form)
		if err != nil {
			fmt.Printf("GlobalHandler - topic : %s, err : %s \n", models.TopicBuyback, err.Error())
			return
		}

		if err := h.transaksiUsecase.BuybackStorage(ctx, form); err != nil {
			fmt.Printf("GlobalHandler-BuybackStorage - topic : %s, err : %s \n", models.TopicBuyback, err.Error())
		} else {
			fmt.Printf("GlobalHandler-BuybackStorage - topic : %s - Success \n", models.TopicBuyback)
		}
	}
}
