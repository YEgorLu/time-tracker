package db

import (
	"database/sql"
	"errors"
	"sync"

	"github.com/YEgorLu/time-tracker/internal/config"
)

type DatabaseProvider interface {
	GetConnection() (*sql.DB, error)
	Bootstrap(*sql.DB) error
}

var openedConnections []*sql.DB
var bootrstapOnce sync.Once

func GetConnection() (*sql.DB, error) {
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
		}
	}

	conn, err := provider.GetConnection()
	if err != nil {
		return nil, err
	}
	if err := conn.Ping(); err != nil {
		// pgx сам закрывает подключение в случае ошибки
		return nil, err
	}
	openedConnections = append(openedConnections, conn)
	if config.DB.RecreateOnStart {
		bootrstapOnce.Do(func() {
			provider.Bootstrap(conn)
		})
	}

	return conn, nil
}

func CloseAll() error {
	closingErrors := make([]error, 0, len(openedConnections))
	for _, conn := range openedConnections {
		closingErrors = append(closingErrors, conn.Close())
	}
	return errors.Join(closingErrors...)
}
