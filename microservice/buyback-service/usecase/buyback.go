package usecase

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/mohnaofal/golden/helper/kafka"
	"github.com/mohnaofal/golden/microservice/buyback-service/models"
	"github.com/mohnaofal/golden/microservice/buyback-service/repository"
)

type buybackUsecase struct {
	hargaRepo     repository.HargaRepository
	rekeningRepo  repository.RekeningRepository
	kafkaProducer kafka.KafkaProducer
}

type BuybackUsecase interface {
	Buyback(ctx context.Context, form *models.BuybackRequest) error
}

func NewBuybackUsecase(hargaRepo repository.HargaRepository, rekeningRepo repository.RekeningRepository, kafkaProducer kafka.KafkaProducer) BuybackUsecase {
	return &buybackUsecase{hargaRepo: hargaRepo, rekeningRepo: rekeningRepo, kafkaProducer: kafkaProducer}
}

func (c *buybackUsecase) Buyback(ctx context.Context, form *models.BuybackRequest) error {
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

	if harga.Data.HargaBuyback != form.Harga {
		return errors.New("harga buyback tidak sama")
	}

	// cek no rekening
	rekening, err := c.rekeningRepo.Get(ctx, &models.RekeningRequest{Norek: form.Norek})
	if err != nil {
		return err
	}

	if rekening == nil {
		return errors.New("rekening not found")
	}

	if (rekening.Data.Saldo - form.Gram) < 0 {
		return errors.New("saldo tidak mencukupi")
	}

	// assign harga topup
	form.HargaTopup = harga.Data.HargaTopup

	msgBytes, err := json.Marshal(form)
	if err != nil {
		return err
	}

	data := &kafka.KafkaSendRequest{
		Topic:     models.TopicBuyback,
		Messages:  string(msgBytes),
		Partition: 0,
	}

	// send procedur
	if err = c.kafkaProducer.KafkaSendProducer(data); err != nil {
		return errors.New("kafka not ready")
	}

	return nil
}
