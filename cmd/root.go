/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/todo-list-app1/todo-list-backend/cfg/db"
	"gitlab.com/todo-list-app1/todo-list-backend/cfg/http"
	"gitlab.com/todo-list-app1/todo-list-backend/cfg/yaml"
	"gitlab.com/todo-list-app1/todo-list-backend/lib/logger"
	"log"
)

func Start() {
	rootCmd := &cobra.Command{}
	//set logger here
	logger.LoggerJson()

	//setting config
	cfg, err := yaml.NewConfig()
	if err != nil {
		panic(err)
	}

	migrate := &cobra.Command{
		Use:   "db:migrate",
		Short: "Migrate",
		Run: func(cmd *cobra.Command, args []string) {
			db.DatabaseMigration(cfg, args)
		},
	}

	cmd := []*cobra.Command{
		{
			Use:   "http",
			Short: "Run HTTP Server",
			Run: func(cmd *cobra.Command, args []string) {
				http.Run(cfg)
			},
		},
	}

	rootCmd.AddCommand(cmd...)
	rootCmd.AddCommand(migrate)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
