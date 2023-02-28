/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"go-fiber-v1/cfg/db"
	"go-fiber-v1/cfg/http"
	"go-fiber-v1/cfg/yaml"
	"go-fiber-v1/lib/logger"
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
