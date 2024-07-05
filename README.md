# Команды для запуска:
```
go mod tidy
go run ./cmd/time-tracker
```
##### Лучше запускать из общей папки, и в ней же создать .env, logs.json файлы
##### Swagger доступен по http://localhost:<порт из конфига>/swagger/index.html

### Флаги в формате <флаг в компандой строке> (аналог в .env или в окружении)

* -config-path value
        Path to .env file
* -db-name (DB_DBNAME) string
        Database provider name (default "TimeTracker")
*-db-password (DB_PASSWORD) string
        Password of database user (default "admin")
* -db-port string
        Url to database server (default "5432")
*-db-provider (DB_PROVIDER) string
        Database provider name (default "postgres")
*  -db-recreate (DB_RECREATE) bool
        Delete database and recreate from migrations
*  -db-url (DB_URL) string
        Url to database server (default "localhost")
*  -db-user (DB_USER) string
        Username of database user (default "admin")
*  -log-level (LOG_LEVEL) value
        log level for logger: 'panic', 'fatal', 'error', 'warn' | 'warning', 'info', 'debug', 'trace'
*  -log-path (LOGS_PATH) value
        Path to .json file for logs
*  -migrations-path (MIGRATIONS_FOLDER) value
        Absolute or relative to project root path to migratinos folder
*  -port (PORT) string
        Port to run server (default "8080")
