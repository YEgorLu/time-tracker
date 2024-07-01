package db

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5"
)

var _ DatabaseProvider = &PostgresProvider{}

type PostgresProvider struct {
	username string
	password string
	url      string
	port     string
	dbName   string
}

func (p PostgresProvider) GetConnection() (*sql.DB, error) {
	if p.url == "" {
		p.url = "localhost"
	}
	if p.port == "" {
		p.port = "5432"
	}

	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", p.username, p.password, p.url, p.port, p.dbName)
	return sql.Open("pgx", connectionString)
}

func (p PostgresProvider) Bootstrap(conn *sql.DB) error {
	//TODO: применить миграции
	return nil
}
