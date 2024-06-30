package main

import (
	"errors"
	"flag"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var AppConfig struct {
	DB      DBConfig
	Port    string
	EnvPath string
}

type DBConfig struct {
	Provider        string
	Url             string
	Port            string
	User            string
	Password        string
	RecreateOnStart bool
}

type paramBinding[T any] struct {
	from string
	to   T
}

var stringsParams = []paramBinding[*string]{
	{"PORT", &AppConfig.Port},
	{"DB_PROVIDER", &AppConfig.DB.Provider},
	{"DB_URL", &AppConfig.DB.Url},
	{"DB_USER", &AppConfig.DB.User},
	{"DB_PASSWORD", &AppConfig.DB.Password},
}

var boolsParams = []paramBinding[*bool]{
	{"DB_RECREATE", &AppConfig.DB.RecreateOnStart},
}

func init() {
	println("initting config")
	parseFlags()
	parseEnvironment()
	parseFileConfig()
}

func parseFlags() {
	flag.Func("config-path", "Path to .env file", func(s string) error {
		file, err := os.Open(s)
		if err != nil {
			return err
		}
		file.Close()
		AppConfig.EnvPath = s
		return nil
	})

	flag.StringVar(&AppConfig.Port, "port", "8080", "Port to run server")
	flag.StringVar(&AppConfig.DB.Provider, "db-provider", "postgres", "Database provider name")
	flag.StringVar(&AppConfig.DB.Url, "db-url", "localhost", "Url to database server")
	flag.StringVar(&AppConfig.DB.User, "db-user", "admin", "Username of database user")
	flag.StringVar(&AppConfig.DB.Password, "db-password", "admin", "Password of database user")
	flag.BoolVar(&AppConfig.DB.RecreateOnStart, "db-recreate", false, "Delete database and recreate from migrations")
	flag.Parse()
}

func parseEnvironment() {
	for _, v := range stringsParams {
		if envValue := os.Getenv(v.from); envValue != "" {
			*v.to = envValue
		}
	}

	for _, v := range boolsParams {
		if envValue, err := strconv.ParseBool(os.Getenv(v.from)); err == nil {
			*v.to = envValue
		}
	}
}

func parseFileConfig() {
	if AppConfig.EnvPath == "" {
		return
	}
	envMap, err := godotenv.Read(AppConfig.EnvPath)
	if err != nil {
		panic(errors.Join(errors.New("invalid .env path provided"), err))
	}

	for _, v := range stringsParams {
		if envValue, ok := envMap[v.from]; ok {
			*v.to = envValue
		}
	}

	for _, v := range boolsParams {
		if envValue, ok := envMap[v.from]; ok {
			if envBoolValue, err := strconv.ParseBool(envValue); err != nil {
				panic(err)
			} else {
				*v.to = envBoolValue
			}
		}
	}
}
