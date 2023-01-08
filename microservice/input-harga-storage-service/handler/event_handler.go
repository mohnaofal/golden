package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/Shopify/sarama"
	"github.com/mohnaofal/golden/helper/kafka"
	"github.com/mohnaofal/golden/microservice/input-harga-storage-service/config"
	"github.com/mohnaofal/golden/microservice/input-harga-storage-service/models"
	"github.com/mohnaofal/golden/microservice/input-harga-storage-service/repository"
	"github.com/mohnaofal/golden/microservice/input-harga-storage-service/usecase"
)

type eventHandler struct {
	hargaUsecase  usecase.HargaUsecase
	kafkaConsumer kafka.KafkaConsumer
}

type EventHandler interface {
	RegisterConsumer()
	GlobalHandler(ctx context.Context, msg *sarama.ConsumerMessage)
}

func NewEventHandler(cfg *config.Config) EventHandler {
	hargaRepository := repository.NewHargaRepository(cfg)
	hargaUsecase := usecase.NewHargaUsecase(hargaRepository)
	return &eventHandler{
		hargaUsecase:  hargaUsecase,
		kafkaConsumer: cfg.KafkaConsumer(),
	}
}

func (h *eventHandler) RegisterConsumer() {
	// list topic
	topics := []string{
		models.TopicInputHarga,
	}

	signals := make(chan os.Signal, 1)

	// init Consume
	h.kafkaConsumer.Consume(topics, h.GlobalHandler, signals)

}

func (h *eventHandler) GlobalHandler(ctx context.Context, msg *sarama.ConsumerMessage) {
	switch msg.Topic {
	case models.TopicInputHarga:
		form := new(models.HargaRequest)
		err := json.Unmarshal(msg.Value, &form)
		if err != nil {
			fmt.Printf("GlobalHandler - topic : %s, err : %s \n", models.TopicInputHarga, err.Error())
			return
		}

		if err := h.hargaUsecase.InputHargaStorage(ctx, form); err != nil {
			fmt.Printf("GlobalHandler-InputHargaStorage - topic : %s, err : %s \n", models.TopicInputHarga, err.Error())
		} else {
			fmt.Printf("GlobalHandler-InputHargaStorage - topic : %s - Success \n", models.TopicInputHarga)
		}
	}
}
