package cmd

import (
	"financials/internal/app/postgres"
	"financials/internal/app/postgres/seed"
	"fmt"
	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Seed database with data",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
		fmt.Println("migrate called")

		db := postgres.Connection()

		seed.Execute(db)

	},
}

func init() {
	rootCmd.AddCommand(seedCmd)
}
