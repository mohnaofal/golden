package kafka

import (
	"log"

	"github.com/Shopify/sarama"
)

type KafkaProducer struct {
	Producer sarama.SyncProducer
}

type KafkaSendRequest struct {
	Topic     string
	Messages  string
	Partition int
}

// NewProducer
func NewProducer(host string) (KafkaProducer, error) {
	configKafka := sarama.NewConfig()
	configKafka.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{host}, configKafka)
	if err != nil {
		return KafkaProducer{}, err
	}

	return KafkaProducer{Producer: producer}, err
}

// KafkaSendProducer
func (p *KafkaProducer) KafkaSendProducer(data *KafkaSendRequest) error {
	kafkaMsg := &sarama.ProducerMessage{
		Topic:     data.Topic,
		Value:     sarama.StringEncoder(data.Messages),
		Partition: int32(data.Partition),
	}

	partition, offset, err := p.Producer.SendMessage(kafkaMsg)
	if err != nil {
		log.Printf("Send message error: %v", err)
		return err
	}

	log.Printf("Send message success, Topic %v, Partition %v, Offset %d", data.Topic, partition, offset)
	return nil
}
