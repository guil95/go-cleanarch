package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/guil95/go-cleanarch/api"
	"github.com/guil95/go-cleanarch/cmd/workers"
	db "github.com/guil95/go-cleanarch/pkg/mysql"
	"github.com/spf13/cobra"
)

func main() {
	mysqlDatabase := db.Connect()

	var api = &cobra.Command{
		Use: "api",
		Run: func(cmd *cobra.Command, args []string) {
			api.Run(mysqlDatabase)
		},
	}

	var consumer = &cobra.Command{
		Use: "consumer",
		Run: func(cmd *cobra.Command, args []string) {
			workers.Run(mysqlDatabase)
		},
	}

	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(api)
	rootCmd.AddCommand(consumer)
	err := rootCmd.Execute()

	if err != nil {
		return
	}
}
