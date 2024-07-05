package db

import (
	"database/sql"
	"fmt"
	"path"

	"github.com/YEgorLu/time-tracker/internal/logger"
	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var _ DatabaseProvider = &PostgresProvider{}

type PostgresProvider struct {
	username string
	password string
	url      string
	port     string
	dbName   string
	log      logger.Logger
}

func (p PostgresProvider) GetConnection() (*sql.DB, error) {
	p.log.Debug("connecting to postgrs url: ", p.Url())
	return sql.Open("pgx", p.Url())
}

func (p PostgresProvider) Bootstrap(conn *sql.DB, migrationsFolder string) error {
	pathPrefix := "file://"
	p.log.Debug("using folder for migrations: ", migrationsFolder)
	if path.IsAbs(migrationsFolder) {
		pathPrefix = "file:///"
	}

	m, err := migrate.New(pathPrefix+migrationsFolder, p.Url())
	if err != nil {
		return err
	}
	m.Log = p.log
	if err := m.Drop(); err != nil {
		return err
	}
	if err := m.Up(); err != nil {
		return err
	}
	return nil
}

func (p PostgresProvider) Url() string {
	if p.url == "" {
		p.url = "localhost"
	}
	if p.port == "" {
		p.port = "5432"
	}

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", p.username, p.password, p.url, p.port, p.dbName)
}
