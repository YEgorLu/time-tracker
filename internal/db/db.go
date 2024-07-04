package db

import (
	"database/sql"
	"errors"
	"sync"

	"github.com/YEgorLu/time-tracker/internal/config"
	"github.com/YEgorLu/time-tracker/internal/logger"
)

type DatabaseProvider interface {
	GetConnection() (*sql.DB, error)
	Bootstrap(*sql.DB, string) error
}

var openedConnections []*sql.DB
var bootrstapOnce sync.Once

func GetConnection(log logger.Logger) *sql.DB {
	providerName := config.DB.Provider
	var provider DatabaseProvider
	switch providerName {
	default:
		provider = &PostgresProvider{
			username: config.DB.User,
			password: config.DB.Password,
			url:      config.DB.Url,
			port:     config.DB.Port,
			dbName:   config.DB.DbName,
			log:      log,
		}
	}

	conn, err := provider.GetConnection()
	if err != nil {
		log.Error(err)
		panic(err)
	}
	if err := conn.Ping(); err != nil {
		// pgx сам закрывает подключение в случае ошибки
		log.Error(err)
		panic(err)
	}
	openedConnections = append(openedConnections, conn)
	if config.DB.RecreateOnStart {
		bootrstapOnce.Do(func() {
			if err := provider.Bootstrap(conn, config.DB.MigrationsFolderPath); err != nil {
				log.Error(err)
				panic(err)
			}
		})
	}

	return conn
}

func CloseAll() error {
	closingErrors := make([]error, 0, len(openedConnections))
	for _, conn := range openedConnections {
		closingErrors = append(closingErrors, conn.Close())
	}
	return errors.Join(closingErrors...)
}
