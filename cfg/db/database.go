package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"go-fiber-v1/cfg/yaml"
	"log"
	"time"
)

type database struct {
	Db *sql.DB
}

func NewDatabase(cfg *yaml.Config) database {
	var (
		dialect         = cfg.DB.Dialect
		maxOpen         = cfg.DB.MaxOpen
		maxIdle         = cfg.DB.MaxIdle
		connMaxLifeTime = cfg.DB.LifeTimeMs
	)

	connStr := ConnString(cfg)

	// Connect to database
	db, err := sql.Open(dialect, connStr)
	if err != nil {
		log.Fatal(err)
	}

	//add setting
	db.SetConnMaxLifetime(time.Minute * time.Duration(connMaxLifeTime))
	db.SetMaxOpenConns(maxOpen)
	db.SetMaxIdleConns(maxIdle)

	return database{Db: db}
}

func ConnString(cfg *yaml.Config) string {
	var (
		dialect      = cfg.DB.Dialect
		username     = cfg.DB.Username
		password     = cfg.DB.Password
		host         = cfg.DB.Host
		databaseName = cfg.DB.DbName
	)

	if dialect == "mysql" {
		return fmt.Sprintf(`%s:%s@tcp(%s:3306)/%s`, username, password, host, databaseName)
	}

	return fmt.Sprintf(`postgresql://%s:%s@%s/%s?sslmode=disable`, username, password, host, databaseName)
}
