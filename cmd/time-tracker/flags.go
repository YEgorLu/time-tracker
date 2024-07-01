package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/YEgorLu/time-tracker/internal/config"
	"github.com/YEgorLu/time-tracker/internal/util"
	"github.com/joho/godotenv"
)

type paramBinding[T any] struct {
	from string
	to   T
}

type fileParamBinding struct {
	paramBinding[*string]
	ext string
}

var f fileParamBinding

var filePathParams = []fileParamBinding{
	fileParamBinding{paramBinding: paramBinding[*string]{"LOGS_PATH", &config.App.LogsPath}, ext: ".json"},
}

var stringsParams = []paramBinding[*string]{
	{"PORT", &config.App.Port},
	{"DB_PROVIDER", &config.DB.Provider},
	{"DB_URL", &config.DB.Url},
	{"DB_USER", &config.DB.User},
	{"DB_PASSWORD", &config.DB.Password},
	{"DB_DBNAME", &config.DB.DbName},
}

var boolsParams = []paramBinding[*bool]{
	{"DB_RECREATE", &config.DB.RecreateOnStart},
}

func init() {
	println("initting config")
	parseFlags()
	parseEnvironment()
	parseFileConfig()
}

func parseFlags() {
	flag.Func("config-path", "Path to .env file", checkFile(&config.App.EnvPath, ".env"))
	flag.Func("log-path", "Path to .json file for logs", checkFile(&config.App.LogsPath, ".json"))

	flag.StringVar(&config.App.Port, "port", "8080", "Port to run server")
	flag.StringVar(&config.DB.Provider, "db-provider", "postgres", "Database provider name")
	flag.StringVar(&config.DB.DbName, "db-name", "TimeTracker", "Database provider name")
	flag.StringVar(&config.DB.Url, "db-url", "localhost", "Url to database server")
	flag.StringVar(&config.DB.Port, "db-port", "5432", "Url to database server")
	flag.StringVar(&config.DB.User, "db-user", "admin", "Username of database user")
	flag.StringVar(&config.DB.Password, "db-password", "admin", "Password of database user")
	flag.BoolVar(&config.DB.RecreateOnStart, "db-recreate", false, "Delete database and recreate from migrations")
	flag.Parse()
}

func parseEnvironment() {
	for _, v := range filePathParams {
		if envValue := os.Getenv(v.from); envValue != "" {
			checkFile(v.to, v.ext)(envValue)
		}
	}

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
	if config.App.EnvPath == "" {
		return
	}
	envMap, err := godotenv.Read(config.App.EnvPath)
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

func checkFile(toPastePath *string, ext string) func(string) error {
	return func(s string) error {
		exists, err := util.FileExists(s)
		if err != nil {
			panic(err)
		}
		if !exists {
			return os.ErrNotExist
		}
		if !util.FileHasExt(s, ext) {
			return newExtensionError(s, ext)
		}
		writable, err := util.FileWritable(s)
		if err != nil {
			panic(err)
		}
		if !writable {
			return newNotWritableError(s)
		}
		*toPastePath = s
		return nil
	}
}

func newExtensionError(path string, wantExt string) error {
	return errors.New(fmt.Sprintf("invalid file extension, want: %s, path: %s", wantExt, path))
}

func newNotWritableError(path string) error {
	return errors.New(fmt.Sprintf("file not writable, path: %s", path))
}