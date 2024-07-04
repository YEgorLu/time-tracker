package config

var App = struct {
	Port     string
	EnvPath  string
	LogsPath string
}{
	Port:     "8080",
	EnvPath:  "../.env",
	LogsPath: "../logs.json",
}

var DB = struct {
	Provider             string
	Url                  string
	Port                 string
	User                 string
	Password             string
	DbName               string
	RecreateOnStart      bool
	MigrationsFolderPath string
}{
	Provider:             "postgres",
	Url:                  "localhost",
	Port:                 "5432",
	User:                 "admin",
	Password:             "16ed2ad3b",
	DbName:               "TimeTracker",
	RecreateOnStart:      false,
	MigrationsFolderPath: "./migrations",
}
