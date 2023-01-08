package usecase

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/mohnaofal/golden/helper/kafka"
	"github.com/mohnaofal/golden/microservice/input-harga-service/models"
)

type hargaUsecase struct {
	kafkaProducer kafka.KafkaProducer
}

type HargaUsecase interface {
	InputHarga(ctx context.Context, form *models.HargaRequest) error
}

func NewHargaUsecase(kafkaProducer kafka.KafkaProducer) HargaUsecase {
	return &hargaUsecase{kafkaProducer: kafkaProducer}
}

func (c *hargaUsecase) InputHarga(ctx context.Context, form *models.HargaRequest) error {
	msgBytes, err := json.Marshal(form)
	if err != nil {
		return err
	}

	data := &kafka.KafkaSendRequest{
		Topic:     models.TopicInputHarga,
		Messages:  string(msgBytes),
		Partition: 0,
	}

	if err = c.kafkaProducer.KafkaSendProducer(data); err != nil {
		return errors.New("kafka not ready")
	}

	return nil
}
