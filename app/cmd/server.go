package cmd

import (
	"MyWallet/api/server"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the API server",
	Run: func(cmd *cobra.Command, args []string) {
		log.Fatal(server.RunAPI("127.0.0.1:8080"))
	},
}
