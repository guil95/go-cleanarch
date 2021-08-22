package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/guil95/go-cleanarch/api"
	"github.com/guil95/go-cleanarch/cmd/workers"
	"github.com/guil95/go-cleanarch/pkg/mongo"
	"github.com/spf13/cobra"
)

func main() {
	mongoDatabase := mongo.Connect()

	var a = &cobra.Command{
		Use: "api",
		Run: func(cmd *cobra.Command, args []string) {
			api.Run(mongoDatabase)
		},
	}

	var consumer = &cobra.Command{
		Use: "consumer",
		Run: func(cmd *cobra.Command, args []string) {
			workers.Run(mongoDatabase)
		},
	}

	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(a)
	rootCmd.AddCommand(consumer)
	err := rootCmd.Execute()

	if err != nil {
		return
	}
}
