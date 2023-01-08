package usecase

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/mohnaofal/golden/helper/kafka"
	"github.com/mohnaofal/golden/microservice/topup-service/models"
	"github.com/mohnaofal/golden/microservice/topup-service/repository"
)

type topupUsecase struct {
	hargaRepo     repository.HargaRepository
	kafkaProducer kafka.KafkaProducer
}

type TopupUsecase interface {
	Topup(ctx context.Context, form *models.TopupRequest) error
}

func NewTopupUsecase(hargaRepo repository.HargaRepository, kafkaProducer kafka.KafkaProducer) TopupUsecase {
	return &topupUsecase{hargaRepo: hargaRepo, kafkaProducer: kafkaProducer}
}

func (c *topupUsecase) Topup(ctx context.Context, form *models.TopupRequest) error {
	if err := form.Validate(); err != nil {
		return err
	}

	harga, err := c.hargaRepo.Get(ctx)
	if err != nil {
		return err
	}

	if harga == nil {
		return errors.New("harga not found")
	}

	if harga.Data.HargaTopup != form.Harga {
		return errors.New("harga topup tidak sama")
	}

	form.HargaBuyback = harga.Data.HargaBuyback

	msgBytes, err := json.Marshal(form)
	if err != nil {
		return err
	}

	data := &kafka.KafkaSendRequest{
		Topic:     models.TopicTopup,
		Messages:  string(msgBytes),
		Partition: 0,
	}

	if err = c.kafkaProducer.KafkaSendProducer(data); err != nil {
		return errors.New("kafka not ready")
	}

	return nil
}
