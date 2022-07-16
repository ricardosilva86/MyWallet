package cmd

import (
	"MyWallet/api/server"
	"MyWallet/util"
	"fmt"
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
		config, err := util.LoadServerConfig(".")
		if err != nil {
			log.Fatalf("error loading server config: %s", err)
		}
		log.Fatal(server.RunAPI(fmt.Sprintf("%s:%s", config.Host, config.Port)))
	},
}
