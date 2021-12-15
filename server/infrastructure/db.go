package infrastructure

import (
	"database/sql"
	"errors"
	"log"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

type Pg struct {
	DB *sql.DB
}

var lock = &sync.Mutex{}

var Pgsql *Pg

func newConnection() (*Pg, error) {
	if Pgsql == nil {
		lock.Lock()
		defer lock.Unlock()
		db, err := sql.Open("postgres", "postgres://postgres:password@localhost:5432/gospy?sslmode=disable")
		if err != nil {
			return nil, errors.New("cannot connect to postgres")
		}

		log.Println("Connected to postgres", db)

		db.SetMaxOpenConns(25)
		db.SetMaxIdleConns(25)
		db.SetConnMaxLifetime(5 * time.Minute)
		Pgsql = &Pg{DB: db}
	}

	return Pgsql, nil
}

func HandlePostgre() {
	_, err := newConnection()
	if err != nil {
		panic(err)
	}
}
