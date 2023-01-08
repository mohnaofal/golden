package config

import (
	"log"
	"os"
	"strconv"

	"github.com/mohnaofal/golden/helper/databases"
	"github.com/mohnaofal/golden/helper/kafka"
)

type Config struct {
	port          string
	sqlDB         databases.ConnPostgres
	kafkaConsumer kafka.KafkaConsumer
}

func LoadConfig() *Config {
	cfg := new(Config)

	cfg.SetPORT()
	cfg.InitSQLDB()
	cfg.InitKafkaConsumer()

	return cfg
}

func (c *Config) SetPORT() {
	c.port = os.Getenv("PORT")
}

func (c *Config) GetPORT() string {
	return c.port
}

func (c *Config) InitSQLDB() {
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	pass := os.Getenv("POSTGRES_PASS")
	dbname := os.Getenv("POSTGRES_DBNAME")
	port, _ := strconv.Atoi(os.Getenv("POSTGRES_PORT"))

	c.sqlDB = databases.Initialize(host, user, pass, dbname, port)
}

func (c *Config) SQLDB() databases.ConnPostgres {
	return c.sqlDB
}

func (c *Config) InitKafkaConsumer() {
	host := os.Getenv("KAFKA_HOST")
	consumer, err := kafka.NewConsumer(host)
	if err != nil {
		log.Fatal(err)
	}

	c.kafkaConsumer = consumer
}

func (c *Config) KafkaConsumer() kafka.KafkaConsumer {
	return c.kafkaConsumer
}
