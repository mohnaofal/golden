package databases

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type conn struct {
	sqlDB *sql.DB
}

type ConnPostgres interface {
	SqlDB() *sql.DB
}

func (c *conn) SqlDB() *sql.DB {
	return c.sqlDB
}

func Initialize(host, user, password, dbname string, port int) ConnPostgres {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxLifetime(time.Minute)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return &conn{sqlDB: db}
}
