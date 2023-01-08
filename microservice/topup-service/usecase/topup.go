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
	rekeningRepo  repository.RekeningRepository
	kafkaProducer kafka.KafkaProducer
}

type TopupUsecase interface {
	Topup(ctx context.Context, form *models.TopupRequest) error
}

func NewTopupUsecase(hargaRepo repository.HargaRepository, rekeningRepo repository.RekeningRepository, kafkaProducer kafka.KafkaProducer) TopupUsecase {
	return &topupUsecase{hargaRepo: hargaRepo, rekeningRepo: rekeningRepo, kafkaProducer: kafkaProducer}
}

func (c *topupUsecase) Topup(ctx context.Context, form *models.TopupRequest) error {
	if err := form.Validate(); err != nil {
		return err
	}

	// cek harga
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

	// cek no rekening
	rekening, err := c.rekeningRepo.Get(ctx, &models.RekeningRequest{Norek: form.Norek})
	if err != nil {
		return err
	}

	if rekening == nil {
		return errors.New("rekening not found")
	}

	// assign harga buyback
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

	// send procedur
	if err = c.kafkaProducer.KafkaSendProducer(data); err != nil {
		return errors.New("kafka not ready")
	}

	return nil
}
