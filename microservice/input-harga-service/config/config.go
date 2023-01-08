package config

import (
	"log"
	"os"

	"github.com/mohnaofal/golden/helper/kafka"
)

type Config struct {
	port          string
	kafkaProducer kafka.KafkaProducer
}

func LoadConfig() *Config {
	cfg := new(Config)

	cfg.SetPORT()
	cfg.InitKafkaProcedur()

	return cfg
}

func (c *Config) SetPORT() {
	c.port = os.Getenv("PORT")
}

func (c *Config) GetPORT() string {
	return c.port
}

func (c *Config) InitKafkaProcedur() {
	host := os.Getenv("KAFKA_HOST")
	procedur, err := kafka.NewProducer(host)
	if err != nil {
		log.Fatal(err)
	}

	c.kafkaProducer = procedur
}

func (c *Config) KafkaProcedur() kafka.KafkaProducer {
	return c.kafkaProducer
}
