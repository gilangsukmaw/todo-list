package yaml

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	App App `yaml:"app,omitempty" json:"app"`
	DB  DB  `yaml:"db" json:"db"`
}

type App struct {
	Name         string `yaml:"name,omitempty" json:"name"`
	Port         string `yaml:"port,omitempty" json:"port"`
	ReadTimeOut  int    `yaml:"read_time_out" json:"read_time_out"`
	WriteTimeOut int    `yaml:"write_time_out" json:"write_time_out"`
}

type DB struct {
	Dialect       string `yaml:"dialect" json:"dialect"`
	Host          string `yaml:"host" json:"host"`
	Port          string `yaml:"port" json:"port"`
	DbName        string `yaml:"db_name" json:"db_name"`
	Username      string `yaml:"username" json:"username"`
	Password      string `yaml:"password" json:"password"`
	MaxOpen       int    `yaml:"max_open" json:"max_open"`
	MaxIdle       int    `yaml:"max_idle" json:"max_idle"`
	TimeOutSecond int    `yaml:"time_out_second" json:"time_out_second"`
	LifeTimeMs    int    `yaml:"life_time_ms" json:"life_time_ms"`
	Charset       string `yaml:"charset" json:"charset"`
}

func NewConfig() (*Config, error) {
	var config *Config

	yfile, err := ioutil.ReadFile("./cfg/yaml/app.yaml")

	if err != nil {
		fmt.Println("error 1", err)
	}

	err = yaml.Unmarshal(yfile, &config)
	if err != nil {
		fmt.Println("error 2", err)
	}

	return config, nil
}
