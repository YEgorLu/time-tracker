package config

var App struct {
	Port     string
	EnvPath  string
	LogsPath string
}

var DB struct {
	Provider        string
	Url             string
	Port            string
	User            string
	Password        string
	DbName          string
	RecreateOnStart bool
}
