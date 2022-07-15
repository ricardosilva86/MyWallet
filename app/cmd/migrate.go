package cmd

import (
	"MyWallet/api/migrations"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrations",
	Short: "Migrations will sync your schema with the database",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Initializing DB")
		migrations.InitDB()
		log.Println("Performing Database Migration...")
		if ok := migrations.Migrate(); ok != nil {
			log.Printf("error while performing migrations: %s\n", ok)
		} else {
			log.Println("Migrations finished with success")
		}
	},
}
