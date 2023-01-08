package config

import (
	"os"
	"strconv"

	"github.com/mohnaofal/golden/helper/databases"
)

type Config struct {
	port  string
	sqlDB databases.ConnPostgres
}

func LoadConfig() *Config {
	cfg := new(Config)

	cfg.SetPORT()
	cfg.InitSQLDB() // postgres sql

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
