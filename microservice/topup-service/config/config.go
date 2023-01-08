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
	kafkaProducer kafka.KafkaProducer

	hargaSvrHost string
}

func LoadConfig() *Config {
	cfg := new(Config)

	cfg.SetPORT()
	cfg.InitSQLDB()
	cfg.InitKafkaProcedur()
	cfg.SetHargaSvrHost()

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

func (c *Config) SetHargaSvrHost() {
	c.hargaSvrHost = os.Getenv("HARGA_SVR_HOST")
}

func (c *Config) HargaSvrHost() string {
	return c.hargaSvrHost
}
