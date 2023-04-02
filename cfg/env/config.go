package cfg

import (
	"fmt"
	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Config struct {
	//App App `yaml:"name,omitempty" json:"app"`
	//DB  DB  `yaml:"db,omitempty" json:"db"`

	Name         string `env:"APP_NAME" yaml:"name,omitempty" json:"name"`
	Port         string `env:"APP_PORT" yaml:"port,omitempty" json:"port"`
	ReadTimeOut  int    `env:"APP_READ_TIMEOUT" yaml:"read_time_out" json:"read_time_out"`
	WriteTimeOut int    `env:"APP_WRITE_TIMEOUT" yaml:"write_time_out" json:"write_time_out"`
	Environment  string `env:"APP_ENVIRONMENT" yaml:"environment" json:"environment"`

	DatabaseDialect       string `env:"DB_DIALECT" yaml:"dialect" json:"dialect"`
	DatabaseHost          string `env:"DB_HOST" yaml:"host" json:"host"`
	DatabaseDbPort        string `env:"DB_PORT" yaml:"db_port" json:"db_portport"`
	DatabaseName          string `env:"DB_DBNAME" yaml:"db_name" json:"db_name"`
	DatabaseUsername      string `env:"DB_USER" yaml:"username" json:"username"`
	DatabasePassword      string `env:"DB_PASSWORD" yaml:"password" json:"password"`
	DatabaseMaxOpen       int    `env:"DB_MAX_OPEN" yaml:"max_open" json:"max_open"`
	DatabaseMaxIdle       int    `env:"DB_MAX_IDLE" yaml:"max_idle" json:"max_idle"`
	DatabaseTimeOutSecond int    `env:"DB_TIMEOUT_SECOND" yaml:"time_out_second" json:"time_out_second"`
	DatabaseLifeTimeMs    int    `env:"DB_LIFE_TIME_MS" yaml:"life_time_ms" json:"life_time_ms"`
	DatabaseCharset       string `env:"DB_CHARSET" yaml:"charset" json:"charset"`

	JwtToken   string `env:"JWT_TOKEN" yaml:"jwt_token" json:"jwt_token"`
	JwtExpired int    `env:"JWT_EXPIRED" yaml:"jwt_expired" json:"jwt_expired"`
}

func NewDotEnvConfig() (*Config, error) {
	var config Config

	config.Environment = os.Getenv("APP_ENVIRONMENT")

	if config.Environment == "local" || config.Environment == "" {
		err := godotenv.Load(".env")
		if err != nil {
			panic(fmt.Sprintf(`error loading env --> %v`, err))
		}

		if err := env.Parse(&config); err != nil {
			fmt.Printf("%+v\n", err)
		}
	} else {
		//TODO find better way to handle env in dockerfile
		readTimeOut, err := strconv.Atoi(os.Getenv("APP_READ_TIMEOUT"))
		if err != nil {
			panic(err)
		}

		writeTimeOut, err := strconv.Atoi(os.Getenv("APP_READ_TIMEOUT"))
		if err != nil {
			panic(err)
		}

		dbMaxOpen, err := strconv.Atoi(os.Getenv("DB_MAX_OPEN"))
		if err != nil {
			panic(err)
		}

		dbMaxIdle, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE"))
		if err != nil {
			panic(err)
		}

		dbTimeOutSecond, err := strconv.Atoi(os.Getenv("DB_TIMEOUT_SECOND"))
		if err != nil {
			panic(err)
		}

		dbLifeTimeMs, err := strconv.Atoi(os.Getenv("DB_LIFE_TIME_MS"))
		if err != nil {
			panic(err)
		}

		jwtExpired, err := strconv.Atoi(os.Getenv("JWT_EXPIRED"))
		if err != nil {
			panic(err)
		}

		config.Name = os.Getenv("APP_NAME")
		config.Port = os.Getenv("APP_PORT")
		config.ReadTimeOut = readTimeOut
		config.WriteTimeOut = writeTimeOut
		config.DatabaseDialect = os.Getenv("DB_DIALECT")
		config.DatabaseHost = os.Getenv("DB_HOST")
		config.DatabaseDbPort = os.Getenv("DB_PORT")
		config.DatabaseName = os.Getenv("DB_DBNAME")
		config.DatabaseUsername = os.Getenv("DB_USER")
		config.DatabasePassword = os.Getenv("DB_PASSWORD")
		config.DatabaseMaxOpen = dbMaxOpen
		config.DatabaseMaxIdle = dbMaxIdle
		config.DatabaseTimeOutSecond = dbTimeOutSecond
		config.DatabaseLifeTimeMs = dbLifeTimeMs
		config.DatabaseCharset = os.Getenv("DB_CHARSET")
		config.JwtToken = os.Getenv("JWT_TOKEN")
		config.JwtExpired = jwtExpired
	}

	return &config, nil
}
