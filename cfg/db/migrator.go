package db

import (
	"fmt"
	"github.com/pressly/goose/v3"
	"go-fiber-v1/cfg/yaml"
)

func MigratorUp(cfg *yaml.Config) {
	var (
		username     = cfg.DB.Username
		password     = cfg.DB.Password
		host         = cfg.DB.Host
		databaseName = cfg.DB.DbName
	)

	connStr := fmt.Sprintf(`postgresql://%s:%s@%s/%s?sslmode=disable`, username, password, host, databaseName)
	conn, err := goose.OpenDBWithDriver("postgres", connStr)

	fmt.Println("asu koe")
}
