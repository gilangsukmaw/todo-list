package db

import (
	"database/sql"
	"flag"
	"github.com/pressly/goose/v3"
	cfg "gitlab.com/todo-list-app1/todo-list-backend/cfg/env"
	"log"
)

var (
	flags = flag.NewFlagSet("goose", flag.ExitOnError)
	dir   = flags.String("dir", "./cfg/db/migrations", "directory with migration files")
)

func DatabaseMigration(cfg *cfg.Config, args []string) {

	if args[0] == "-h" || args[0] == "--help" {
		flags.Usage()
		return
	}

	var driver = cfg.DatabaseDialect

	if len(args) > 1 && args[0] == "create" {
		if err := goose.Run("create", nil, *dir, args[1:]...); err != nil {
			log.Fatalf("goose run: %v", err)
		}
		return
	}

	arguments := []string{}
	if len(args) > 1 {
		arguments = append(arguments, args[0:]...)
	}

	dbstring := ConnString(cfg)

	db, err := sql.Open(driver, dbstring)
	if err != nil {
		log.Fatalf("-dbstring=%q: %v\n", dbstring, err)
	}

	if err := goose.Run(args[0], db, *dir, arguments...); err != nil {
		log.Fatalf("goose run: %v", err)
	}

}

func usage() {
	log.Print(usagePrefix)
	flags.PrintDefaults()
	log.Print(usageCommands)
}

var (
	usagePrefix = `Usage: goose [OPTIONS] DRIVER DBSTRING COMMAND
Drivers:
    postgres
    mysql
    sqlite3
    redshift
Examples:
    goose sqlite3 ./foo.db status
    goose sqlite3 ./foo.db create init sql
    goose sqlite3 ./foo.db create add_some_column sql
    goose sqlite3 ./foo.db create fetch_user_data go
    goose sqlite3 ./foo.db up
    goose postgres "user=postgres dbname=postgres sslmode=disable" status
    goose mysql "user:password@/dbname?parseTime=true" status
    goose redshift "postgres://user:password@qwerty.us-east-1.redshift.amazonaws.com:5439/db"
status
Options:
`

	usageCommands = `
Commands:
    up                   Migrate the DB to the most recent version available
    up-to VERSION        Migrate the DB to a specific VERSION
    down                 Roll back the version by 1
    down-to VERSION      Roll back to a specific VERSION
    redo                 Re-run the latest migration
    status               Dump the migration status for the current DB
    version              Print the current version of the database
    create NAME [sql|go] Creates new migration file with next version
`
)
