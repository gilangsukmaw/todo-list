package cfg

import (
	"fmt"
	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"
)

type Config struct {
	//App App `yaml:"name,omitempty" json:"app"`
	//DB  DB  `yaml:"db,omitempty" json:"db"`

	Name         string `env:"APP_NAME" yaml:"name,omitempty" json:"name"`
	Port         string `env:"APP_PORT" yaml:"port,omitempty" json:"port"`
	ReadTimeOut  int    `env:"APP_READ_TIMEOUT" yaml:"read_time_out" json:"read_time_out"`
	WriteTimeOut int    `env:"APP_WRITE_TIMEOUT" yaml:"write_time_out" json:"write_time_out"`

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
}

func NewDotEnvConfig() (*Config, error) {
	var config Config

	err := godotenv.Load(".env")
	if err != nil {
		panic(fmt.Sprintf(`error loading env --> %v`, err))
	}

	if err := env.Parse(&config); err != nil {
		fmt.Printf("%+v\n", err)
	}

	fmt.Println("config --> ", config)

	return &config, nil
}