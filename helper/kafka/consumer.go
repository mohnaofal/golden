package kafka

import (
	"context"
	"log"
	"os"

	"github.com/Shopify/sarama"
)

type KafkaConsumer struct {
	Consumer sarama.Consumer
}

func NewConsumer(host string) (KafkaConsumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// Create new consumer
	conn, err := sarama.NewConsumer([]string{host}, config)
	if err != nil {
		return KafkaConsumer{}, err
	}

	return KafkaConsumer{Consumer: conn}, nil
}

// Consume function to consume message from apache kafka
func (c *KafkaConsumer) Consume(topics []string, handler func(ctx context.Context, msg *sarama.ConsumerMessage), signals chan os.Signal) {
	chanMessage := make(chan *sarama.ConsumerMessage, 256)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for _, topic := range topics {
		partitionList, err := c.Consumer.Partitions(topic)
		if err != nil {
			log.Printf("Unable to get partition got error %v", err)
			continue
		}
		for _, partition := range partitionList {
			go consumeMessage(c.Consumer, topic, partition, chanMessage)
		}
	}

	log.Printf("Kafka is consuming....")

	for {
		select {
		case msg := <-chanMessage:
			handler(ctx, msg)
		case sig := <-signals:
			if sig == os.Interrupt {
				c.Consumer.Close()
				return
			}
		}
	}
}

func consumeMessage(consumer sarama.Consumer, topic string, partition int32, c chan *sarama.ConsumerMessage) {
	msg, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
	if err != nil {
		log.Printf("Unable to consume partition %v got error %v", partition, err)
		return
	}

	defer func() {
		if err := msg.Close(); err != nil {
			log.Printf("Unable to close partition %v: %v", partition, err)
		}
	}()

	for {
		msg := <-msg.Messages()
		c <- msg
	}

}
