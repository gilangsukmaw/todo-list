package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	cfg "gitlab.com/todo-list-app1/todo-list-backend/cfg/env"
	"log"
	"time"
)

type database struct {
	Db *sql.DB
}

func NewDatabase(cfg *cfg.Config) database {
	var (
		dialect         = cfg.DatabaseDialect
		maxOpen         = cfg.DatabaseMaxOpen
		maxIdle         = cfg.DatabaseMaxIdle
		connMaxLifeTime = cfg.DatabaseLifeTimeMs
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

func ConnString(cfg *cfg.Config) string {
	var (
		dialect      = cfg.DatabaseDialect
		username     = cfg.DatabaseUsername
		password     = cfg.DatabasePassword
		host         = cfg.DatabaseHost
		databaseName = cfg.DatabaseName
	)

	if dialect == "mysql" {
		return fmt.Sprintf(`%s:%s@tcp(%s:3306)/%s`, username, password, host, databaseName)
	}

	return fmt.Sprintf(`postgresql://%s:%s@%s/%s?sslmode=disable`, username, password, host, databaseName)
}
