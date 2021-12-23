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
		db, err := sql.Open("postgres", "postgres://postgres:password@localhost:5432/gospy?sslmode=disable&parseTime=true")
		if err != nil {
			return nil, errors.New("cannot connect to postgres")
		}

		log.Println("Connected to postgres")

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

func (db *Pg) Query(statment string, args ...interface{}) *sql.Rows {
	result, err := db.DB.Query(statment, args...)
	if err != nil {
		panic(err)
	}

	return result
}

func (db *Pg) Insert(statment string, args ...interface{}) int64 {
	var id int64
	err := db.DB.QueryRow(statment+" RETURNING id", args...).Scan(&id)
	if err != nil {
		panic(err)
	}
	return id
}
