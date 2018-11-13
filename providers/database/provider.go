package database

import (
	"database/sql"
	"fmt"
	"log"

	"gitlab.com/upaphong/go-assignment/engine"
)

const (
	host     = "db" // localhost
	port     = 5432
	user     = "postgres"
	password = "pwd123"
	dbname   = "goassignment"
)

type Provider struct {
	DB *sql.DB
}

func (provider *Provider) GetKnightRepository() engine.KnightRepository {
	return &knightRepository{db: provider.DB}
}

func (provider *Provider) Close() {
	provider.DB.Close()
}

func NewProvider() *Provider {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	log.Println("db connected!")
	return &Provider{DB: db}
}
