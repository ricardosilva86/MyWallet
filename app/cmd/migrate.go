package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate will sync your schema with the database",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Running migrations")
	},
}
