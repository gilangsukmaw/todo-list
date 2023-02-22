package db

import (
	"database/sql"
	"fmt"
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
		username        = cfg.DB.Username
		password        = cfg.DB.Password
		host            = cfg.DB.Host
		databaseName    = cfg.DB.DbName
		maxOpen         = cfg.DB.MaxOpen
		maxIdle         = cfg.DB.MaxIdle
		connMaxLifeTime = cfg.DB.LifeTimeMs
	)

	connStr := fmt.Sprintf(`postgresql://%s:%s@%s/%s?sslmode=disable`, username, password, host, databaseName)
	// Connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	//add setting
	db.SetConnMaxLifetime(time.Minute * time.Duration(connMaxLifeTime))
	db.SetMaxOpenConns(maxOpen)
	db.SetMaxIdleConns(maxIdle)

	return database{Db: db}
}

func (db *database) QueryRow() {

}

//
//func DbConnect() {
//
//	db, err := sql.Open("mysql", "user7:s$cret@tcp(127.0.0.1:3306)/testdb")
//	defer db.Close()
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	res, err := db.Query("SELECT * FROM cities")
//
//	defer res.Close()
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	for res.Next() {
//
//		var city City
//		err := res.Scan(&city.Id, &city.Name, &city.Population)
//
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		fmt.Printf("%v\n", city)
//	}
//}
